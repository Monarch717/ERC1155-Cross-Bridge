package engine

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/erc721"
)

var (
	querySwapPairLimit      = 50
	watchRegisterEventDelay = time.Duration(3) * time.Second
)

// queryERC721SwapPair queries SwapPair this engine is responsible
func (e *Engine) queryERC721SwapPair(fromChainID string, states []erc721.SwapPairState) ([]*erc721.SwapPair, error) {
	// TODO: check the index
	var ss []*erc721.SwapPair
	err := e.deps.DB.Where(
		"state in ? and src_chain_id = ?",
		states,
		fromChainID,
	).Order(
		"register_height asc",
	).Limit(
		querySwapPairLimit,
	).Find(&ss).Error
	if err != nil {
		return nil, errors.Wrap(err, "[Engine.queryERC721SwapPair]: failed to query SwapPair")
	}

	return ss, nil
}

// filterERC721ConfirmedRegisterEvents checks block confirmation of the chain this engine is responsible
func (e *Engine) filterERC721ConfirmedRegisterEvents(ss []*erc721.SwapPair) (events []*erc721.SwapPair, err error) {
	for _, s := range ss {
		confirmed, err := e.hasBlockConfirmed(s.RegisterTxHash, e.chainID())
		if err != nil {
			util.Logger.Warning(errors.Wrap(err, "[Engine.filterERC721ConfirmedRegisterEvents]: failed to check block confirmation"))
			continue
		}
		if confirmed {
			events = append(events, s)
		}
	}

	return events, nil
}

func (e *Engine) generateERC721TxHash(s *erc721.SwapPair) (string, error) {
	request, err := e.sendERC721CreatePairRequest(s, true)
	if err != nil {
		return "", errors.Wrap(err, "[Engine.generateERC721TxHash]: failed to dry run sending a request")
	}

	return request.Hash().String(), nil
}

// sendERC721CreatePairRequest sends transaction to create a swap pair on destination chain
func (e *Engine) sendERC721CreatePairRequest(s *erc721.SwapPair, dryRun bool) (*types.Transaction, error) {
	dstChainID := s.DstChainID
	dstChainIDInt := util.StrToBigInt(dstChainID)
	if _, ok := e.deps.Client[dstChainID]; !ok {
		return nil, errors.Errorf("[Engine.sendERC721CreatePairRequest]: client for chain id %s is not supported", dstChainID)
	}
	if _, ok := e.deps.ERC721SwapAgent[dstChainID]; !ok {
		return nil, errors.Errorf("[Engine.sendERC721CreatePairRequest]: swap agent for chain id %s is not supported", dstChainID)
	}

	ctx := context.Background()
	txOpts, err := util.TxOpts(ctx, e.deps.Client[dstChainID], e.conf.PrivateKey, dstChainIDInt)
	if err != nil {
		return nil, errors.Wrap(err, "[Engine.sendERC721CreatePairRequest]: failed to create tx opts")
	}

	txOpts.NoSend = dryRun

	tx, err := e.deps.ERC721SwapAgent[dstChainID].CreateSwapPair(
		txOpts,
		common.HexToHash(s.RegisterTxHash),
		common.HexToAddress(s.SrcTokenAddr),
		util.StrToBigInt(s.SrcChainID),
		s.BaseURI,
		s.SrcTokenName,
		s.Symbol,
	)
	if err != nil {
		return nil, errors.Wrap(err, "[Engine.sendERC721CreatePairRequest]: failed to send swap pair creation tx")
	}

	return tx, nil
}
