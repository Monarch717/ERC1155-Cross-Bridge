package engine

import (
	"math/big"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/block"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/erc721"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

func (e *Engine) manageERC721TxCreatedSwap() {
	fromChainID := e.chainID()
	ss, err := e.queryERC721Swap(fromChainID, []erc721.SwapState{
		erc721.SwapStateFillTxCreated,
	})
	if err != nil {
		util.Logger.Error(errors.Wrap(err, "[Engine.manageERC721TxCreatedSwap]: failed to query tx_created Swaps"))
		return
	}

	for _, s := range ss {
		ethTx, isPending, err := e.retrieveTx(s.FillTxHash, s.DstChainID)
		if err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC721TxCreatedSwap]: failed to get Swap creation tx %s", s.FillTxHash),
			)

			continue
		}
		if isPending {
			util.Logger.Infof("[Engine.manageERC721TxCreatedSwap]: the tx %s is pending in mempools, skip", s.FillTxHash)
			continue
		}

		receipt, err := e.retrieveTxReceipt(s.FillTxHash, s.DstChainID)
		if err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC721TxCreatedSwap]: failed to get Swap creation receipt for tx %s", s.FillTxHash),
			)

			continue
		}

		if ethTx == nil {
			util.Logger.Infof("[Engine.manageERC721TxCreatedSwap]: the tx is not found while cheking tx %s", s.FillTxHash)
		}

		if receipt == nil {
			util.Logger.Infof("[Engine.manageERC721TxCreatedSwap]: the receipt is not found while cheking tx %s", s.FillTxHash)
		}

		if ethTx == nil || receipt == nil {
			s.FillTrackRetry += 1
			if err := e.deps.DB.Save(s).Error; err != nil {
				util.Logger.Error(
					errors.Wrapf(err, "[Engine.manageERC721TxCreatedSwap]: failed to increase create track retry counter %s", s.ID),
				)

				continue
			}

			if s.FillTrackRetry > e.conf.MaxTrackRetry {
				s.State = erc721.SwapStateFillTxMissing
				s.MessageLog = "[Engine.manageERC721TxCreatedSwap]: tx is missing"
				if err := e.deps.DB.Save(s).Error; err != nil {
					util.Logger.Error(
						errors.Wrapf(err, "[Engine.manageERC721TxCreatedSwap]: failed to update Swap %s to '%s' state", s.ID, s.State),
					)

					continue
				}
			}

			continue
		}

		var b block.Log
		err = e.deps.DB.Where(
			"chain_id = ? and block_hash = ?",
			s.DstChainID,
			receipt.BlockHash.String(),
		).Select(
			"id",
		).First(
			&b,
		).Error
		if err == gorm.ErrRecordNotFound {
			util.Logger.Infof("[Engine.manageERC721TxCreatedSwap]: wait for the system to catch up the block %s in chain id %s", receipt.BlockHash.String(), e.chainID())

			continue
		}
		if err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC721TxCreatedSwap]: failed to update Swap %s to '%s' state", s.ID, s.State),
			)

			continue
		}

		var isValid bool
		fillBlockHeight := receipt.BlockNumber.Int64()
		if s.SwapDirection == erc721.SwapDirectionForward {
			isValid, err = e.verifyERC721ForwardSwapFillEvent(uint64(fillBlockHeight), s.RequestTxHash, s.DstChainID)
			if err != nil {
				util.Logger.Error(
					errors.Wrapf(err, "[Engine.manageERC721TxCreatedSwap]: failed to verify swap fill event for Swap %s", s.ID),
				)

				continue
			}
		} else {
			isValid, err = e.verifyERC721BackwardSwapFillEvent(uint64(fillBlockHeight), s.RequestTxHash, s.DstChainID)
			if err != nil {
				util.Logger.Error(
					errors.Wrapf(err, "[Engine.manageERC721TxCreatedSwap]: failed to verify backward swap fill event for Swap %s", s.ID),
				)

				continue
			}
		}
		if !isValid {
			s.State = erc721.SwapStateFillTxFailed
			s.MessageLog = "[Engine.manageERC721TxCreatedSwap]: swap fill event was not found!"
			if err := e.deps.DB.Save(s).Error; err != nil {
				util.Logger.Error(
					errors.Wrapf(err, "[Engine.manageERC721TxCreatedSwap]: failed to update Swap %s to '%s' state", s.ID, s.State),
				)
			}

			continue
		}

		gasPrice := big.NewInt(0)
		gasPrice.SetString(ethTx.GasPrice().String(), 10)
		s.FillGasPrice = gasPrice.String()
		s.FillConsumedFeeAmount = big.NewInt(1).Mul(gasPrice, big.NewInt(s.FillGasUsed)).String()
		s.FillGasUsed = int64(receipt.GasUsed)
		s.FillHeight = fillBlockHeight
		s.FillBlockHash = receipt.BlockHash.String()
		s.FillBlockLogID = &b.ID
		s.State = erc721.SwapStateFillTxSent
		if err := e.deps.DB.Save(s).Error; err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC721TxCreatedSwap]: failed to update Swap %s basic info", s.ID),
			)

			continue
		}

		util.Logger.Infof("[Engine.manageERC721TxCreatedSwap]: updated Swap %s after sending out with tx hash %s", s.ID, s.FillTxHash)
	}
}
