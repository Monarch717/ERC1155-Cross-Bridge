package block

import (
	"time"

	"gorm.io/gorm"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

type Log struct {
	ID         string `gorm:"size:26;primary_key"`
	ChainID    string `gorm:"not null;index:chain_id"`
	BlockHash  string `gorm:"not null;index:block_hash"`
	ParentHash string `gorm:"not null;index:parent_hash"`
	Height     int64  `gorm:"not null;index:height"`
	BlockTime  int64
	CreateTime time.Time
}

func (Log) TableName() string {
	return "block_logs"
}

func (b *Log) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = util.ULID()
	b.CreateTime = time.Now()
	return nil
}
