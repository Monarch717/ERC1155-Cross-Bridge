package observer

import (
	"time"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/common"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/block"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

// Prune prunes the outdated blocks
func (ob *Observer) Prune() {
	for {
		curBlockLog, err := ob.GetCurrentBlockLog()
		if err != nil {
			util.Logger.Errorf("[Observer.Prune]: get current block log error, err=%s", err.Error())
			time.Sleep(common.ObserverPruneInterval)
			continue
		}

		height := curBlockLog.Height - common.ObserverMaxBlockNumber
		if height > 0 {
			err = ob.deps.DB.Where(
				"chain_id = ? and height < ?",
				ob.deps.Recorder.ChainID(),
				height,
			).Delete(
				block.Log{},
			).Error
			if err != nil {
				util.Logger.Errorf("[Observer.Prune]: prune block logs error, err=%s", err.Error())
			}
		}

		time.Sleep(common.ObserverPruneInterval)
	}
}
