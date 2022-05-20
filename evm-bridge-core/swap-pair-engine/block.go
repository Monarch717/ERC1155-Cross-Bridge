package engine

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

func (e *Engine) hasBlockConfirmed(txHash, chainID string) (bool, error) {
	if _, ok := e.deps.Recorder[chainID]; !ok {
		return false, errors.Errorf("[Engine.hasBlockConfirmed]: chain id %s is not supported", chainID)
	}

	block := e.deps.Recorder[chainID].LatestBlockCached()
	if block == nil {
		util.Logger.Infof("[Engine.hasBlockConfirmed]:: no latest block cache found for chain id %s", chainID)

		return false, nil
	}

	ctx := context.Background()
	txRecipient, err := e.deps.Client[chainID].TransactionReceipt(ctx, common.HexToHash(txHash))
	if err != nil {
		return false, errors.Wrap(err, "[Engine.hasBlockConfirmed]: failed to get tx receipt")
	}
	if block.Height < txRecipient.BlockNumber.Int64()+e.conf.ConfirmNum {
		return false, nil
	}

	return true, nil
}
