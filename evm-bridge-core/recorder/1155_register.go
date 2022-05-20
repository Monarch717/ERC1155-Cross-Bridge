package recorder

import (
	"context"
	"math"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/block"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/erc1155"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

func (r *Recorder) recordERC1155RegisterTx(tx *gorm.DB, b *block.Log) error {
	ctx, cancel := context.WithTimeout(context.Background(), registerFilterLogsTimeout)
	defer cancel()

	height := uint64(b.Height)
	opts := bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: ctx,
	}
	iter, err := r.deps.ERC1155SwapAgent[r.ChainID()].FilterSwapPairRegister(&opts, nil, nil)
	if err != nil {
		return errors.Wrap(err, "[Recorder.recordERC1155RegisterTx]: failed to filter logs")
	}
	defer func() {
		if err := iter.Close(); err != nil {
			util.Logger.Errorf("[Recorder.recordERC1155RegisterTx]: failed to close iterator, %s", err.Error())
		}
	}()

	var ss []erc1155.SwapPair
	for iter.Next() {
		s := erc1155.SwapPair{
			SrcChainID:   r.ChainID(),
			DstChainID:   iter.Event.ToChainId.String(),
			SrcTokenAddr: iter.Event.TokenAddress.String(),
			DstTokenAddr: "",
			Sponsor:      iter.Event.Sponsor.String(),
			Available:    false,
			Signature:    "",
			URI:          "",

			State: erc1155.SwapPairStateRegistrationOngoing,

			RegisterTxHash:     iter.Event.Raw.TxHash.String(),
			RegisterHeight:     int64(iter.Event.Raw.BlockNumber),
			RegisterBlockHash:  iter.Event.Raw.BlockHash.String(),
			RegisterBlockLog:   nil,
			RegisterBlockLogID: &b.ID,

			CreateTxHash:     "",
			CreateHeight:     math.MaxInt64,
			CreateBlockHash:  "",
			CreateBlockLog:   nil,
			CreateBlockLogID: nil,
		}

		uri, err := r.retrieveERC1155URI(s.SrcTokenAddr)
		if err != nil {
			return errors.Wrap(err, "[Recorder.recordERC1155RegisterTx]: failed to get uri")
		}

		s.URI = uri
		ss = append(ss, s)
	}

	if err := iter.Error(); err != nil {
		return errors.Wrap(err, "[Recorder.recordERC1155RegisterTx]: failed to iterate events")
	}

	err = tx.Clauses(
		clause.OnConflict{DoNothing: true},
	).Omit(
		clause.Associations,
	).CreateInBatches(
		&ss, 100,
	).Error
	if err != nil {
		return errors.Wrap(err, "[Recorder.recordERC1155RegisterTx]: failed to bulk create")
	}

	return nil
}

func (r *Recorder) retrieveERC1155URI(tokenAddr string) (string, error) {
	token, ok := r.deps.ERC1155Token[r.ChainID()]
	if !ok {
		return "", errors.Errorf("[Recorder.retrieveERC1155URI]: unsupported chain id %s", r.ChainID())
	}

	opts := &bind.CallOpts{
		Pending: true,
	}
	uri, err := token.URI(opts, tokenAddr)
	if err != nil {
		return "", errors.Wrap(err, "[Recorder.retrieveERC1155URI]: failed to retrieve base URI")
	}

	return uri, nil
}
