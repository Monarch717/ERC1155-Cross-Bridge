package recorder

import (
	"context"
	"math"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/block"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/erc721"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

var (
	swapFilterLogsTimeout = time.Duration(20) * time.Second
)

func (r *Recorder) recordERC721SwapTx(tx *gorm.DB, b *block.Log) error {
	ctx, cancel := context.WithTimeout(context.Background(), swapFilterLogsTimeout)
	defer cancel()

	height := uint64(b.Height)
	opts := bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: ctx,
	}
	iter, err := r.deps.ERC721SwapAgent[r.ChainID()].FilterSwapStarted(&opts, nil, nil, nil)
	if err != nil {
		return errors.Wrap(err, "[Recorder.recordERC721SwapTx]: failed to filter logs")
	}
	defer func() {
		if err := iter.Close(); err != nil {
			util.Logger.Errorf("[Recorder.recordERC721SwapTx]: failed to close iterator, %s", err.Error())
		}
	}()

	var ss []erc721.Swap
	for iter.Next() {
		s := erc721.Swap{
			SrcChainID:            r.ChainID(),
			DstChainID:            iter.Event.DstChainId.String(),
			SrcTokenAddr:          iter.Event.TokenAddr.String(),
			DstTokenAddr:          "",
			SrcTokenName:          "",
			DstTokenName:          "",
			Sender:                iter.Event.Sender.String(),
			Recipient:             iter.Event.Recipient.String(),
			TokenID:               iter.Event.TokenId.String(),
			TokenURI:              "",
			Signature:             "",
			State:                 erc721.SwapStateRequestOngoing,
			SwapDirection:         erc721.SwapDirectionForward,
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
		return errors.Wrap(err, "[Recorder.recordERC721SwapTx]: failed to iterate events")
	}

	err = tx.Clauses(
		clause.OnConflict{DoNothing: true},
	).Omit(
		clause.Associations,
	).CreateInBatches(
		&ss, 100,
	).Error
	if err != nil {
		return errors.Wrap(err, "[Recorder.recordERC721SwapTx]: failed to bulk create")
	}

	return nil
}

func (r *Recorder) recordERC721BackwardSwapTx(tx *gorm.DB, b *block.Log) error {
	ctx, cancel := context.WithTimeout(context.Background(), swapFilterLogsTimeout)
	defer cancel()

	height := uint64(b.Height)
	opts := bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: ctx,
	}
	iter, err := r.deps.ERC721SwapAgent[r.ChainID()].FilterBackwardSwapStarted(&opts, nil, nil, nil)
	if err != nil {
		return errors.Wrap(err, "[Recorder.recordERC721BackwardSwapTx]: failed to filter logs")
	}
	defer func() {
		if err := iter.Close(); err != nil {
			util.Logger.Errorf("[Recorder.recordERC721BackwardSwapTx]: failed to close iterator, %s", err.Error())
		}
	}()

	var ss []erc721.Swap
	for iter.Next() {
		s := erc721.Swap{
			SrcChainID:            r.ChainID(),
			DstChainID:            iter.Event.DstChainId.String(),
			SrcTokenAddr:          iter.Event.MirroredTokenAddr.String(),
			DstTokenAddr:          "",
			SrcTokenName:          "",
			DstTokenName:          "",
			Sender:                iter.Event.Sender.String(),
			Recipient:             iter.Event.Recipient.String(),
			TokenID:               iter.Event.TokenId.String(),
			TokenURI:              "",
			Signature:             "",
			State:                 erc721.SwapStateRequestOngoing,
			SwapDirection:         erc721.SwapDirectionBackward,
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
		return errors.Wrap(err, "[Recorder.recordERC721BackwardSwapTx]: failed to iterate events")
	}

	err = tx.Clauses(
		clause.OnConflict{DoNothing: true},
	).Omit(
		clause.Associations,
	).CreateInBatches(
		&ss, 100,
	).Error
	if err != nil {
		return errors.Wrap(err, "[Recorder.recordERC721BackwardSwapTx]: failed to bulk create")
	}

	return nil
}
