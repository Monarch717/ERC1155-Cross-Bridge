package common

import (
	"errors"
	"time"
)

const (
	ObserverMaxBlockNumber = 10000
	ObserverPruneInterval  = 10 * time.Second
	ObserverAlertInterval  = 5 * time.Second

	DBDialectMysql   = "mysql"
	DBDialectSqlite3 = "sqlite3"

	LocalPrivateKey = "local_private_key"
	AWSPrivateKey   = "aws_private_key"
)

var (
	ErrBlockNotFound    = errors.New("block is not found")
	ErrFunctionNotFound = errors.New("attempting to unmarshall an empty string while arguments are expected")
)

type Block struct {
	Height          int64
	Chain           string
	BlockHash       string
	ParentBlockHash string
	BlockTime       int64
}
