package recorder

import (
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"

	"github.com/pkg/errors"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/block"
)

func (r *Recorder) Record(tx *gorm.DB, b *block.Log) error {
	var g errgroup.Group
	g.Go(func() error {
		if err := r.recordERC721RegisterTx(tx, b); err != nil {
			return errors.Wrap(err, "[Recorder.Record]: failed to record ERC721 register tx")
		}
		if err := r.recordERC721SwapTx(tx, b); err != nil {
			return errors.Wrap(err, "[Recorder.Record]: failed to record ERC721 swap tx")
		}
		if err := r.recordERC721BackwardSwapTx(tx, b); err != nil {
			return errors.Wrap(err, "[Recorder.Record]: failed to record ERC721 backward swap tx")
		}

		return nil
	})

	g.Go(func() error {
		if err := r.recordERC1155RegisterTx(tx, b); err != nil {
			return errors.Wrap(err, "[Recorder.Record]: failed to record ERC1155 register tx")
		}
		if err := r.recordERC1155SwapTx(tx, b); err != nil {
			return errors.Wrap(err, "[Recorder.Record]: failed to record ERC1155 swap tx")
		}
		if err := r.recordERC1155BackwardSwapTx(tx, b); err != nil {
			return errors.Wrap(err, "[Recorder.Record]: failed to record ERC1155 backward swap tx")
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return errors.Wrap(err, "[Recorder.Record]: failed to monitor events")
	}

	return nil
}
