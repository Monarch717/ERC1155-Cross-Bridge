package engine

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/erc1155"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

// queryERC1155SwapPair queries SwapPair this engine is responsible
func (e *Engine) queryERC1155SwapPair(fromChainID string, states []erc1155.SwapPairState) ([]*erc1155.SwapPair, error) {
	// TODO: check the index
	var ss []*erc1155.SwapPair
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
		return nil, errors.Wrap(err, "[Engine.queryERC1155SwapPair]: failed to query SwapPair")
	}

	return ss, nil
}

// filterERC1155ConfirmedRegisterEvents checks block confirmation of the chain this engine is responsible
func (e *Engine) filterERC1155ConfirmedRegisterEvents(ss []*erc1155.SwapPair) (events []*erc1155.SwapPair, err error) {
	for _, s := range ss {
		confirmed, err := e.hasBlockConfirmed(s.RegisterTxHash, e.chainID())
		if err != nil {
			util.Logger.Warning(errors.Wrap(err, "[Engine.filterERC1155ConfirmedRegisterEvents]: failed to check block confirmation"))
			continue
		}
		if confirmed {
			events = append(events, s)
		}
	}

	return events, nil
}

func (e *Engine) generateERC1155TxHash(s *erc1155.SwapPair) (string, error) {
	request, err := e.sendERC1155CreatePairRequest(s, true)
	if err != nil {
		return "", errors.Wrap(err, "[Engine.generateERC1155TxHash]: failed to dry run sending a request")
	}

	return request.Hash().String(), nil
}

// sendERC1155CreatePairRequest sends transaction to create a swap pair on destination chain
func (e *Engine) sendERC1155CreatePairRequest(s *erc1155.SwapPair, dryRun bool) (*types.Transaction, error) {
	dstChainID := s.DstChainID
	dstChainIDInt := util.StrToBigInt(dstChainID)
	if _, ok := e.deps.Client[dstChainID]; !ok {
		return nil, errors.Errorf("[Engine.sendERC1155CreatePairRequest]: client for chain id %s is not supported", dstChainID)
	}
	if _, ok := e.deps.ERC1155SwapAgent[dstChainID]; !ok {
		return nil, errors.Errorf("[Engine.sendERC1155CreatePairRequest]: swap agent for chain id %s is not supported", dstChainID)
	}

	ctx := context.Background()
	txOpts, err := util.TxOpts(ctx, e.deps.Client[dstChainID], e.conf.PrivateKey, dstChainIDInt)
	if err != nil {
		return nil, errors.Wrap(err, "[Engine.sendERC1155CreatePairRequest]: failed to create tx opts")
	}

	txOpts.NoSend = dryRun

	tx, err := e.deps.ERC1155SwapAgent[dstChainID].CreateSwapPair(
		txOpts,
		common.HexToHash(s.RegisterTxHash),
		common.HexToAddress(s.SrcTokenAddr),
		util.StrToBigInt(s.SrcChainID),
		s.URI,
	)
	if err != nil {
		return nil, errors.Wrap(err, "[Engine.sendERC1155CreatePairRequest]: failed to send swap pair creation tx")
	}

	return tx, nil
}
