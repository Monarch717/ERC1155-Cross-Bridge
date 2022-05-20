package recorder

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/common"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/block"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/erc721"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

var (
	registerFilterLogsTimeout = time.Duration(20) * time.Second
)

func (r *Recorder) recordERC721RegisterTx(tx *gorm.DB, b *block.Log) error {
	ctx, cancel := context.WithTimeout(context.Background(), registerFilterLogsTimeout)
	defer cancel()

	height := uint64(b.Height)
	opts := bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: ctx,
	}
	iter, err := r.deps.ERC721SwapAgent[r.ChainID()].FilterSwapPairRegister(&opts, nil, nil)
	if err != nil {
		return errors.Wrap(err, "[Recorder.recordERC721RegisterTx]: failed to filter logs")
	}
	defer func() {
		if err := iter.Close(); err != nil {
			util.Logger.Errorf("[Recorder.recordERC721RegisterTx]: failed to close iterator, %s", err.Error())
		}
	}()

	var ss []erc721.SwapPair
	for iter.Next() {
		s := erc721.SwapPair{
			SrcChainID:   r.ChainID(),
			DstChainID:   iter.Event.ToChainId.String(),
			SrcTokenAddr: iter.Event.TokenAddress.String(),
			DstTokenAddr: "",
			SrcTokenName: iter.Event.TokenName,
			DstTokenName: fmt.Sprintf("%s mirrored from %s", iter.Event.TokenName, r.ChainName()),
			Sponsor:      iter.Event.Sponsor.String(),
			Available:    false,
			Signature:    "",
			Symbol:       iter.Event.TokenSymbol,
			BaseURI:      "",

			State: erc721.SwapPairStateRegistrationOngoing,

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

		baseURI, err := r.retrieveERC21BaseURI(s.SrcTokenAddr)
		if err != nil {
			return errors.Wrap(err, "[Recorder.recordERC721RegisterTx]: failed to get baseURI")
		}
		if baseURI == "" {
			util.Logger.Infof("[Recorder.recordERC721RegisterTx]: chain id %s, token %s has no baseURI", s.SrcChainID, s.SrcTokenAddr)
		}

		s.BaseURI = baseURI
		ss = append(ss, s)
	}

	if err := iter.Error(); err != nil {
		return errors.Wrap(err, "[Recorder.recordERC721RegisterTx]: failed to iterate events")
	}

	err = tx.Clauses(
		clause.OnConflict{DoNothing: true},
	).Omit(
		clause.Associations,
	).CreateInBatches(
		&ss, 100,
	).Error
	if err != nil {
		return errors.Wrap(err, "[Recorder.recordERC721RegisterTx]: failed to bulk create")
	}

	return nil
}

func (r *Recorder) retrieveERC21BaseURI(tokenAddr string) (string, error) {
	token, ok := r.deps.ERC721Token[r.ChainID()]
	if !ok {
		return "", errors.Errorf("[Recorder.retrieveERC21BaseURI]: unsupported chain id %s", r.ChainID())
	}

	opts := &bind.CallOpts{
		Pending: true,
	}
	uri, err := token.BaseURI(opts, tokenAddr)
	if err != nil {
		if strings.Contains(err.Error(), common.ErrFunctionNotFound.Error()) {
			return "", nil
		}

		return "", errors.Wrap(err, "[Recorder.retrieveERC21BaseURI]: failed to retrieve base URI")
	}

	return uri, nil
}
