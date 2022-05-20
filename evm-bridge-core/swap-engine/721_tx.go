package engine

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	corecommon "github.com/synycboom/bsc-evm-compatible-bridge-core/common"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

func (e *Engine) verifyERC721ForwardSwapFillEvent(height uint64, requestTxHash, chainID string) (bool, error) {
	if _, ok := e.deps.Client[chainID]; !ok {
		return false, errors.Errorf("[Engine.verifyERC721ForwardSwapFillEvent]: client for chain id %s is not supported", chainID)
	}

	ctx := context.Background()
	opts := bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: ctx,
	}
	txHash := [32]byte(common.HexToHash(requestTxHash))
	iter, err := e.deps.ERC721SwapAgent[chainID].FilterSwapFilled(&opts, [][32]byte{txHash}, nil, nil)
	if err != nil {
		return false, errors.Wrap(err, "[Engine.verifyERC721ForwardSwapFillEvent]: failed to filter logs")
	}
	defer func() {
		if err := iter.Close(); err != nil {
			util.Logger.Errorf("[Engine.verifyERC721ForwardSwapFillEvent]: failed to close iterator, %s", err.Error())
		}
	}()

	return iter.Next(), nil
}

func (e *Engine) verifyERC721BackwardSwapFillEvent(height uint64, requestTxHash, chainID string) (bool, error) {
	if _, ok := e.deps.Client[chainID]; !ok {
		return false, errors.Errorf("[Engine.verifyERC721BackwardSwapFillEvent]: client for chain id %s is not supported", chainID)
	}

	ctx := context.Background()
	opts := bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: ctx,
	}
	txHash := [32]byte(common.HexToHash(requestTxHash))
	iter, err := e.deps.ERC721SwapAgent[chainID].FilterBackwardSwapFilled(&opts, [][32]byte{txHash}, nil, nil)
	if err != nil {
		return false, errors.Wrap(err, "[Engine.verifyERC721BackwardSwapFillEvent]: failed to filter logs")
	}
	defer func() {
		if err := iter.Close(); err != nil {
			util.Logger.Errorf("[Engine.verifyERC721BackwardSwapFillEvent]: failed to close iterator, %s", err.Error())
		}
	}()

	return iter.Next(), nil
}

func (e *Engine) retrieveERC721TokenURI(tokenAddr, tokenID, chainID string) (string, error) {
	token, ok := e.deps.ERC721Token[chainID]
	if !ok {
		return "", errors.Errorf("[Engine.retrieveERC721TokenURI]: unsupported chain id %s", chainID)
	}

	opts := &bind.CallOpts{
		Pending: true,
	}
	tID := util.StrToBigInt(tokenID)
	uri, err := token.TokenURI(opts, tokenAddr, tID)
	if err != nil {
		if strings.Contains(err.Error(), corecommon.ErrFunctionNotFound.Error()) {
			return "", nil
		}

		return "", errors.Wrap(err, "[Recorder.retrieveBaseURI]: failed to retrieve token URI")
	}

	return uri, nil
}
