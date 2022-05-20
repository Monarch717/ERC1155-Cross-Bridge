package engine

import (
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/erc1155"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
	"gorm.io/gorm"
)

func (e *Engine) generateERC1155TxHash(s *erc1155.Swap) (string, error) {
	request, err := e.sendERC1155FillSwapRequest(s, true)
	if err != nil {
		return "", errors.Wrap(err, "[Engine.generateERC1155TxHash]: failed to dry run sending a request")
	}

	return request.Hash().String(), nil
}

// sendERC1155FillSwapRequest sends transaction to fill a swap on destination chain
func (e *Engine) sendERC1155FillSwapRequest(s *erc1155.Swap, dryRun bool) (*types.Transaction, error) {
	dstChainID := s.DstChainID
	dstChainIDInt := util.StrToBigInt(dstChainID)
	if _, ok := e.deps.Client[dstChainID]; !ok {
		return nil, errors.Errorf("[Engine.sendERC1155FillSwapRequest]: client for chain id %s is not supported", dstChainID)
	}
	if _, ok := e.deps.ERC1155SwapAgent[dstChainID]; !ok {
		return nil, errors.Errorf("[Engine.sendERC1155FillSwapRequest]: swap agent for chain id %s is not supported", dstChainID)
	}

	ctx := context.Background()
	txOpts, err := util.TxOpts(ctx, e.deps.Client[dstChainID], e.conf.PrivateKey, dstChainIDInt)
	if err != nil {
		return nil, errors.Wrap(err, "[Engine.sendERC1155FillSwapRequest]: failed to create tx opts")
	}

	tokenAddr := s.SrcTokenAddr
	tokenChainID := s.SrcChainID
	if s.SwapDirection == erc1155.SwapDirectionBackward {
		tokenAddr = s.DstTokenAddr
		tokenChainID = s.SrcChainID
	}

	var ids []string
	if err := json.Unmarshal(s.IDs, &ids); err != nil {
		return nil, errors.Wrap(err, "[Engine.sendERC1155FillSwapRequest]: failed to unmarshal ids")
	}
	var amounts []string
	if err := json.Unmarshal(s.Amounts, &amounts); err != nil {
		return nil, errors.Wrap(err, "[Engine.sendERC1155FillSwapRequest]: failed to unmarshal ids")
	}

	txOpts.NoSend = dryRun
	tx, err := e.deps.ERC1155SwapAgent[dstChainID].Fill(
		txOpts,
		common.HexToHash(s.RequestTxHash),
		common.HexToAddress(tokenAddr),
		common.HexToAddress(s.Recipient),
		util.StrToBigInt(tokenChainID),
		util.StrSliceToBigIntSlice(ids),
		util.StrSliceToBigIntSlice(amounts),
	)
	if err != nil {
		return nil, errors.Wrap(err, "[Engine.sendERC1155FillSwapRequest]: failed to send swap pair creation tx")
	}

	return tx, nil
}

// queryERC1155Swap queries Swap this engine is responsible
func (e *Engine) queryERC1155Swap(fromChainID string, states []erc1155.SwapState) ([]*erc1155.Swap, error) {
	// TODO: check the index
	var ss []*erc1155.Swap
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
		return nil, errors.Wrap(err, "[Engine.queryERC1155Swap]: failed to query Swap")
	}

	return ss, nil
}

// fillERC1155RequiredInfo fills swap destination tokens
func (e *Engine) fillERC1155RequiredInfo(ss []*erc1155.Swap) error {
	for _, s := range ss {
		if s.IsRequiredInfoValid() {
			continue
		}

		s.RequestTrackRetry += 1

		var skip bool
		var err error
		if s.SwapDirection == erc1155.SwapDirectionForward {
			skip, err = e.fillERC1155Forward(s)
		} else {
			skip, err = e.fillERC1155Backward(s)
		}

		if skip {
			continue
		}
		if err != nil {
			return errors.Wrap(err, "[Engine.fillERC1155RequiredInfo]: failed to fill required information")
		}
	}

	return nil
}

func (e *Engine) fillERC1155Forward(s *erc1155.Swap) (skip bool, err error) {
	// TODO: check db index
	var sp erc1155.SwapPair
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
		return false, errors.Wrap(err, "[Engine.fillERC1155Forward]: failed to query available Swaps")
	}

	s.SetRequiredInfo(sp.DstTokenAddr)

	return false, nil
}

func (e *Engine) fillERC1155Backward(s *erc1155.Swap) (skip bool, err error) {
	// TODO: check db index
	var sp erc1155.SwapPair
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
		return false, errors.Wrap(err, "[Engine.fillERC1155Backward]: failed to query available Swaps")
	}

	s.SetRequiredInfo(sp.SrcTokenAddr)

	return false, nil
}

func (e *Engine) separateERC1155SwapEvents(ss []*erc1155.Swap) (pass []*erc1155.Swap, pending []*erc1155.Swap, rejected []*erc1155.Swap) {
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

// filterERC1155ConfirmedSwapEvents checks block confirmation of the chain this engine is responsible
func (e *Engine) filterERC1155ConfirmedSwapEvents(ss []*erc1155.Swap) (events []*erc1155.Swap, err error) {
	for _, s := range ss {
		confirmed, err := e.hasBlockConfirmed(s.RequestTxHash, e.chainID())
		if err != nil {
			util.Logger.Warning(errors.Wrap(err, "[Engine.filterERC1155ConfirmedSwapEvents]: failed to check block confirmation"))
			continue
		}
		if confirmed {
			events = append(events, s)
		}
	}

	return events, nil
}
