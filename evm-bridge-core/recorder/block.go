package recorder

import (
	"context"
	"math/big"
	"time"

	"github.com/pkg/errors"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/common"
)

func (r *Recorder) Block(height int64) (*common.Block, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	header, err := r.deps.Client[r.ChainID()].HeaderByNumber(ctx, big.NewInt(height))
	if err != nil {
		return nil, errors.Wrap(err, "[Recorder.Block]: failed to get block")
	}

	r.latestBlockCached = &common.Block{
		Height:          height,
		Chain:           r.conf.ChainID.String(),
		BlockHash:       header.Hash().String(),
		ParentBlockHash: header.ParentHash.String(),
		BlockTime:       int64(header.Time),
	}

	return r.latestBlockCached, nil
}

func (r *Recorder) LatestBlockCached() *common.Block {
	return r.latestBlockCached
}
