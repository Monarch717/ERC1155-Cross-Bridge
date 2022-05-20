package engine

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/erc721"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
	"gorm.io/gorm"
)

var (
	watchSwapEventDelay = time.Duration(3) * time.Second
	querySwapLimit      = 50
)

func (e *Engine) generateERC721TxHash(s *erc721.Swap) (string, error) {
	request, err := e.sendERC721FillSwapRequest(s, true)
	if err != nil {
		return "", errors.Wrap(err, "[Engine.generateERC721TxHash]: failed to dry run sending a request")
	}

	return request.Hash().String(), nil
}

// sendERC721FillSwapRequest sends transaction to fill a swap on destination chain
func (e *Engine) sendERC721FillSwapRequest(s *erc721.Swap, dryRun bool) (*types.Transaction, error) {
	dstChainID := s.DstChainID
	dstChainIDInt := util.StrToBigInt(dstChainID)
	if _, ok := e.deps.Client[dstChainID]; !ok {
		return nil, errors.Errorf("[Engine.sendERC721FillSwapRequest]: client for chain id %s is not supported", dstChainID)
	}
	if _, ok := e.deps.ERC721SwapAgent[dstChainID]; !ok {
		return nil, errors.Errorf("[Engine.sendERC721FillSwapRequest]: swap agent for chain id %s is not supported", dstChainID)
	}

	ctx := context.Background()
	txOpts, err := util.TxOpts(ctx, e.deps.Client[dstChainID], e.conf.PrivateKey, dstChainIDInt)
	if err != nil {
		return nil, errors.Wrap(err, "[Engine.sendERC721FillSwapRequest]: failed to create tx opts")
	}

	tokenAddr := s.SrcTokenAddr
	tokenChainID := s.SrcChainID
	if s.SwapDirection == erc721.SwapDirectionBackward {
		tokenAddr = s.DstTokenAddr
		tokenChainID = s.SrcChainID
	}

	txOpts.NoSend = dryRun
	tx, err := e.deps.ERC721SwapAgent[dstChainID].Fill(
		txOpts,
		common.HexToHash(s.RequestTxHash),
		common.HexToAddress(tokenAddr),
		common.HexToAddress(s.Recipient),
		util.StrToBigInt(tokenChainID),
		util.StrToBigInt(s.TokenID),
		s.TokenURI,
	)
	if err != nil {
		return nil, errors.Wrap(err, "[Engine.sendERC721FillSwapRequest]: failed to send swap pair creation tx")
	}

	return tx, nil
}

// queryERC721Swap queries Swap this engine is responsible
func (e *Engine) queryERC721Swap(fromChainID string, states []erc721.SwapState) ([]*erc721.Swap, error) {
	// TODO: check the index
	var ss []*erc721.Swap
	err := e.deps.DB.Where(
		"state in ? and src_chain_id = ?",
		states,
		fromChainID,
	).Order(
		"request_height asc",
	).Limit(
		querySwapLimit,
	).Find(&ss).Error
	if err != nil {
		return nil, errors.Wrap(err, "[Engine.queryERC721Swap]: failed to query Swap")
	}

	return ss, nil
}

// fillERC721RequiredInfo fills swap destination tokens
func (e *Engine) fillERC721RequiredInfo(ss []*erc721.Swap) error {
	for _, s := range ss {
		if s.IsRequiredInfoValid() {
			continue
		}

		s.RequestTrackRetry += 1

		var skip bool
		var err error
		if s.SwapDirection == erc721.SwapDirectionForward {
			skip, err = e.fillERC721Forward(s)
		} else {
			skip, err = e.fillERC721Backward(s)
		}

		if skip {
			continue
		}
		if err != nil {
			return errors.Wrap(err, "[Engine.fillERC721RequiredInfo]: failed to fill required information")
		}
	}

	return nil
}

func (e *Engine) fillERC721Forward(s *erc721.Swap) (skip bool, err error) {
	// TODO: check db index
	var sp erc721.SwapPair
	err = e.deps.DB.Where(
		"src_token_addr = ? and src_chain_id = ? and dst_chain_id = ? and available = ?",
		s.SrcTokenAddr,
		s.SrcChainID,
		s.DstChainID,
		true,
	).First(&sp).Error

	if err == gorm.ErrRecordNotFound {
		return true, nil
	}
	if err != nil {
		return false, errors.Wrap(err, "[Engine.fillERC721Forward]: failed to query available Swaps")
	}

	var tokenURI string
	if sp.BaseURI == "" {
		tokenURI, err = e.retrieveERC721TokenURI(s.SrcTokenAddr, s.TokenID, s.SrcChainID)
		if err != nil {
			return false, errors.Wrapf(err, "[Engine.fillERC721Forward]: failed to retrieve token uri of token %s, chain id %", s.SrcTokenAddr, s.SrcChainID)
		}
		if tokenURI == "" {
			util.Logger.Infof("[Engine.fillERC721Forward]: token %s, chain id % has no token uri", s.SrcTokenAddr, s.SrcChainID)
		}
	} else {
		tokenURI = s.TokenID
	}

	s.SetRequiredInfo(
		sp.SrcTokenName,
		sp.DstTokenName,
		sp.DstTokenAddr,
		sp.BaseURI,
		tokenURI,
	)

	return false, nil
}

func (e *Engine) fillERC721Backward(s *erc721.Swap) (skip bool, err error) {
	// TODO: check db index
	var sp erc721.SwapPair
	err = e.deps.DB.Where(
		"dst_token_addr = ? and dst_chain_id = ? and src_chain_id = ? and available = ?",
		s.SrcTokenAddr,
		s.SrcChainID,
		s.DstChainID,
		true,
	).First(&sp).Error

	if err == gorm.ErrRecordNotFound {
		return true, nil
	}
	if err != nil {
		return false, errors.Wrap(err, "[Engine.fillERC721Backward]: failed to query available Swaps")
	}

	s.SetRequiredInfo(
		sp.DstTokenName,
		sp.SrcTokenName,
		sp.SrcTokenAddr,
		"",
		"",
	)

	return false, nil
}

func (e *Engine) separateERC721SwapEvents(ss []*erc721.Swap) (pass []*erc721.Swap, pending []*erc721.Swap, rejected []*erc721.Swap) {
	for _, s := range ss {
		if !s.IsRequiredInfoValid() {
			if s.RequestTrackRetry > e.conf.MaxTrackRetry {
				rejected = append(rejected, s)
			} else {
				pending = append(pending, s)
			}

			continue
		}

		pass = append(pass, s)
	}

	return
}

// filterERC721ConfirmedSwapEvents checks block confirmation of the chain this engine is responsible
func (e *Engine) filterERC721ConfirmedSwapEvents(ss []*erc721.Swap) (events []*erc721.Swap, err error) {
	for _, s := range ss {
		confirmed, err := e.hasBlockConfirmed(s.RequestTxHash, e.chainID())
		if err != nil {
			util.Logger.Warning(errors.Wrap(err, "[Engine.filterERC721ConfirmedSwapEvents]: failed to check block confirmation"))
			continue
		}
		if confirmed {
			events = append(events, s)
		}
	}

	return events, nil
}
