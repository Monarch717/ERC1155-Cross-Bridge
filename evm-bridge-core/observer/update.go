package observer

import (
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/common"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/block"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

// Update starts the main routine for updating blocks
func (ob *Observer) Update() {
	chainID := ob.deps.Recorder.ChainID()
	startHeight := ob.conf.StartHeight
	for {
		curBlockLog, err := ob.GetCurrentBlockLog()
		if err != nil {
			util.Logger.Errorf("[Observer.Update]: get current block log from db error: %s", err.Error())
			time.Sleep(ob.conf.FetchInterval)
			continue
		}

		nextHeight := curBlockLog.Height + 1
		if curBlockLog.Height == 0 && startHeight != 0 {
			nextHeight = startHeight
		}

		util.Logger.Debugf("[Observer.Update]: fetch from chain id %s, height=%d", chainID, nextHeight)
		err = ob.updateBlock(curBlockLog.Height, nextHeight, curBlockLog.BlockHash)
		if err != nil {
			if errors.Cause(err) != common.ErrBlockNotFound {
				util.Logger.Errorf("[Observer.Update]: fetch from chain id %s error, err=%s", chainID, err.Error())
			}

			util.Logger.Debugf("[Observer.Update]: failed to fetch from chain id %s error, err=%s", chainID, err.Error())

			time.Sleep(ob.conf.FetchInterval)
		}
	}
}

// updateBlock fetches the next block of BSC and saves it to database. if the next block hash
// does not match to the parent hash, the current block will be deleted for there is a fork.
func (ob *Observer) updateBlock(curHeight, nextHeight int64, curBlockHash string) error {
	block, err := ob.deps.Recorder.Block(nextHeight)
	if err != nil {
		return errors.Wrapf(err, "[Observer.updateBlock]: failed to get block info, height=%d", nextHeight)
	}

	if curHeight != 0 && block.ParentBlockHash != curBlockHash {
		if err := ob.DeleteBlock(curHeight); err != nil {
			return errors.Wrap(err, "[Observer.updateBlock]: failed to delete a forked block")
		}

		return nil
	}

	if err := ob.RecordBlockAndTxs(block); err != nil {
		return errors.Wrap(err, "[Observer.updateBlock]: failed to save and process block")
	}

	return nil
}

func (ob *Observer) RecordBlockAndTxs(b *common.Block) error {
	if err := ob.deps.DB.Transaction(func(tx *gorm.DB) error {
		nextBlockLog := block.Log{
			BlockHash:  b.BlockHash,
			ParentHash: b.ParentBlockHash,
			Height:     b.Height,
			BlockTime:  b.BlockTime,
			ChainID:    b.Chain,
		}
		if err := tx.Create(&nextBlockLog).Error; err != nil {
			return errors.Wrap(err, "[Observer.RecordBlockAndTxs]: failed to create block log")
		}

		if err := ob.deps.Recorder.Record(tx, &nextBlockLog); err != nil {
			return errors.Wrap(err, "[Observer.RecordBlockAndTxs]: failed to record a new block from recorder")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "[Observer.RecordBlockAndTxs]: failed to update the block")
	}

	return nil
}

func (ob *Observer) DeleteBlock(height int64) error {
	if err := ob.deps.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("height = ?", height).Delete(block.Log{}).Error; err != nil {
			return errors.Wrap(err, "[Observer.DeleteBlock]: failed to delete block log")
		}

		if err := ob.deps.Recorder.Delete(tx, height); err != nil {
			return errors.Wrap(err, "[Observer.DeleteBlock]: failed to delete block from executor")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "[Observer.DeleteBlock]: failed to delete the block")
	}

	return nil
}
