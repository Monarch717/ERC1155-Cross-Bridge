package recorder

import (
	"context"
	"encoding/json"
	"math"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/block"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/erc1155"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

func (r *Recorder) recordERC1155SwapTx(tx *gorm.DB, b *block.Log) error {
	ctx, cancel := context.WithTimeout(context.Background(), swapFilterLogsTimeout)
	defer cancel()

	height := uint64(b.Height)
	opts := bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: ctx,
	}
	iter, err := r.deps.ERC1155SwapAgent[r.ChainID()].FilterSwapStarted(&opts, nil, nil, nil)
	if err != nil {
		return errors.Wrap(err, "[Recorder.recordERC1155SwapTx]: failed to filter logs")
	}
	defer func() {
		if err := iter.Close(); err != nil {
			util.Logger.Errorf("[Recorder.recordERC1155SwapTx]: failed to close iterator, %s", err.Error())
		}
	}()

	var ss []erc1155.Swap
	for iter.Next() {
		idList := util.BigIntSliceToStrSlice(iter.Event.Ids)
		amountList := util.BigIntSliceToStrSlice(iter.Event.Amounts)
		if len(idList) != len(amountList) {
			util.Logger.Warningf(
				"[Recorder.recordERC1155SwapTx]: chain id %s, token %s, length of ids and amounts are not equal",
				r.ChainID(),
				iter.Event.TokenAddr.String(),
			)

			continue
		}

		ids, err := json.Marshal(idList)
		if err != nil {
			return errors.Wrap(err, "[Recorder.recordERC1155SwapTx]: failed to marshal ids")
		}
		amounts, err := json.Marshal(amountList)
		if err != nil {
			return errors.Wrap(err, "[Recorder.recordERC1155SwapTx]: failed to marshal amounts")
		}

		s := erc1155.Swap{
			SrcChainID:            r.ChainID(),
			DstChainID:            iter.Event.DstChainId.String(),
			SrcTokenAddr:          iter.Event.TokenAddr.String(),
			DstTokenAddr:          "",
			Sender:                iter.Event.Sender.String(),
			Recipient:             iter.Event.Recipient.String(),
			IDs:                   datatypes.JSON(ids),
			Amounts:               datatypes.JSON(amounts),
			Signature:             "",
			State:                 erc1155.SwapStateRequestOngoing,
			SwapDirection:         erc1155.SwapDirectionForward,
			RequestTxHash:         iter.Event.Raw.TxHash.String(),
			RequestHeight:         int64(iter.Event.Raw.BlockNumber),
			RequestBlockHash:      iter.Event.Raw.BlockHash.String(),
			RequestBlockLogID:     &b.ID,
			RequestBlockLog:       nil,
			RequestTrackRetry:     0,
			FillConsumedFeeAmount: "",
			FillGasPrice:          "",
			FillGasUsed:           0,
			FillHeight:            math.MaxInt64,
			FillTxHash:            "",
			FillTrackRetry:        0,
			FillBlockHash:         "",
			FillBlockLogID:        nil,
			FillBlockLog:          nil,
			MessageLog:            "",
		}

		ss = append(ss, s)
	}

	if err := iter.Error(); err != nil {
		return errors.Wrap(err, "[Recorder.recordERC1155SwapTx]: failed to iterate events")
	}

	err = tx.Clauses(
		clause.OnConflict{DoNothing: true},
	).Omit(
		clause.Associations,
	).CreateInBatches(
		&ss, 100,
	).Error
	if err != nil {
		return errors.Wrap(err, "[Recorder.recordERC1155SwapTx]: failed to bulk create")
	}

	return nil
}

func (r *Recorder) recordERC1155BackwardSwapTx(tx *gorm.DB, b *block.Log) error {
	ctx, cancel := context.WithTimeout(context.Background(), swapFilterLogsTimeout)
	defer cancel()

	height := uint64(b.Height)
	opts := bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: ctx,
	}
	iter, err := r.deps.ERC1155SwapAgent[r.ChainID()].FilterBackwardSwapStarted(&opts, nil, nil, nil)
	if err != nil {
		return errors.Wrap(err, "[Recorder.recordERC1155BackwardSwapTx]: failed to filter logs")
	}
	defer func() {
		if err := iter.Close(); err != nil {
			util.Logger.Errorf("[Recorder.recordERC1155BackwardSwapTx]: failed to close iterator, %s", err.Error())
		}
	}()

	var ss []erc1155.Swap
	for iter.Next() {
		idList := util.BigIntSliceToStrSlice(iter.Event.Ids)
		amountList := util.BigIntSliceToStrSlice(iter.Event.Amounts)
		if len(idList) != len(amountList) {
			util.Logger.Warningf(
				"[Recorder.recordERC1155BackwardSwapTx]: chain id %s, token %s, length of ids and amounts are not equal",
				r.ChainID(),
				iter.Event.MirroredTokenAddr.String(),
			)

			continue
		}

		ids, err := json.Marshal(idList)
		if err != nil {
			return errors.Wrap(err, "[Recorder.recordERC1155BackwardSwapTx]: failed to marshal ids")
		}
		amounts, err := json.Marshal(amountList)
		if err != nil {
			return errors.Wrap(err, "[Recorder.recordERC1155BackwardSwapTx]: failed to marshal amounts")
		}

		s := erc1155.Swap{
			SrcChainID:            r.ChainID(),
			DstChainID:            iter.Event.DstChainId.String(),
			SrcTokenAddr:          iter.Event.MirroredTokenAddr.String(),
			DstTokenAddr:          "",
			Sender:                iter.Event.Sender.String(),
			Recipient:             iter.Event.Recipient.String(),
			IDs:                   datatypes.JSON(ids),
			Amounts:               datatypes.JSON(amounts),
			Signature:             "",
			State:                 erc1155.SwapStateRequestOngoing,
			SwapDirection:         erc1155.SwapDirectionBackward,
			RequestTxHash:         iter.Event.Raw.TxHash.String(),
			RequestHeight:         int64(iter.Event.Raw.BlockNumber),
			RequestBlockHash:      iter.Event.Raw.BlockHash.String(),
			RequestBlockLogID:     &b.ID,
			RequestBlockLog:       nil,
			RequestTrackRetry:     0,
			FillConsumedFeeAmount: "",
			FillGasPrice:          "",
			FillGasUsed:           0,
			FillHeight:            math.MaxInt64,
			FillTxHash:            "",
			FillTrackRetry:        0,
			FillBlockHash:         "",
			FillBlockLogID:        nil,
			FillBlockLog:          nil,
			MessageLog:            "",
		}

		ss = append(ss, s)
	}

	if err := iter.Error(); err != nil {
		return errors.Wrap(err, "[Recorder.recordERC1155BackwardSwapTx]: failed to iterate events")
	}

	err = tx.Clauses(
		clause.OnConflict{DoNothing: true},
	).Omit(
		clause.Associations,
	).CreateInBatches(
		&ss, 100,
	).Error
	if err != nil {
		return errors.Wrap(err, "[Recorder.recordERC1155BackwardSwapTx]: failed to bulk create")
	}

	return nil
}
