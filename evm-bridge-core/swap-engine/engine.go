package engine

import (
	"math/big"

	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"
	erc1155agent "github.com/synycboom/bsc-evm-compatible-bridge-core/agent/erc1155"
	erc721agent "github.com/synycboom/bsc-evm-compatible-bridge-core/agent/erc721"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/client"
	recorder "github.com/synycboom/bsc-evm-compatible-bridge-core/recorder"
	erc1155token "github.com/synycboom/bsc-evm-compatible-bridge-core/token/erc1155"
	erc721token "github.com/synycboom/bsc-evm-compatible-bridge-core/token/erc721"
)

type Recorder interface {
}

type Config struct {
	ExplorerURL               string
	PrivateKey                string
	ChainID                   *big.Int
	ConfirmNum                int64
	MaxTrackRetry             int64
	ERC721SwapAgentAddresses  map[string]common.Address
	ERC1155SwapAgentAddresses map[string]common.Address
}

type Dependencies struct {
	Client           map[string]client.ETHClient
	DB               *gorm.DB
	Recorder         map[string]recorder.IRecorder
	ERC721SwapAgent  map[string]erc721agent.SwapAgent
	ERC721Token      map[string]erc721token.IToken
	ERC1155SwapAgent map[string]erc1155agent.SwapAgent
	ERC1155Token     map[string]erc1155token.IToken
}

type Engine struct {
	conf *Config
	deps *Dependencies
}

func NewEngine(c *Config, d *Dependencies) *Engine {
	return &Engine{
		conf: c,
		deps: d,
	}
}
