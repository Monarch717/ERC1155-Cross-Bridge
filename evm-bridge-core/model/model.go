package model

import (
	"gorm.io/gorm"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/block"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/erc1155"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/erc721"
)

func InitTables(db *gorm.DB) {
	db.AutoMigrate(&block.Log{})
	db.AutoMigrate(&erc721.SwapPair{})
	db.AutoMigrate(&erc721.Swap{})
	db.AutoMigrate(&erc1155.SwapPair{})
	db.AutoMigrate(&erc1155.Swap{})
}
