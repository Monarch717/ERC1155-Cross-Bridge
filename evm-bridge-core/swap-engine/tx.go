package engine

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

func (e *Engine) retrieveTx(txHash, chainID string) (*types.Transaction, bool, error) {
	if _, ok := e.deps.Client[chainID]; !ok {
		return nil, false, errors.Errorf("[Engine.retrieveTx]: client for chain id %s is not supported", chainID)
	}

	ctx := context.Background()
	tx, isPending, err := e.deps.Client[chainID].TransactionByHash(ctx, common.HexToHash(txHash))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, isPending, nil
		}

		return nil, isPending, errors.Wrap(err, "[Engine.retrieveTx]: failed to get a transaction")
	}

	return tx, isPending, nil
}

func (e *Engine) retrieveTxReceipt(txHash, chainID string) (*types.Receipt, error) {
	if _, ok := e.deps.Client[chainID]; !ok {
		return nil, errors.Errorf("[Engine.retrieveTxReceipt]: client for chain id %s is not supported", chainID)
	}

	ctx := context.Background()
	txRecipient, err := e.deps.Client[chainID].TransactionReceipt(ctx, common.HexToHash(txHash))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, nil
		}

		return nil, errors.Wrap(err, "[Engine.retrieveTxReceipt]: failed to get a transaction receipt")
	}

	return txRecipient, nil
}
