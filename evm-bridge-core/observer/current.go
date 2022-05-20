package observer

import (
	"github.com/pkg/errors"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/block"
	"gorm.io/gorm"
)

// GetCurrentBlockLog returns the highest block log
func (ob *Observer) GetCurrentBlockLog() (*block.Log, error) {
	blockLog := block.Log{}
	err := ob.deps.DB.Where(
		"chain_id = ?",
		ob.deps.Recorder.ChainID(),
	).Order(
		"height desc",
	).First(
		&blockLog,
	).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "[Observer.GetCurrentBlockLog]: failed to get data from block log")
	}

	return &blockLog, nil
}
