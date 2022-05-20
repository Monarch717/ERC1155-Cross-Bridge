package main

import (
	"flag"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	contractabi "github.com/synycboom/bsc-evm-compatible-bridge-core/abi"
	erc1155agent "github.com/synycboom/bsc-evm-compatible-bridge-core/agent/erc1155"
	erc721agent "github.com/synycboom/bsc-evm-compatible-bridge-core/agent/erc721"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/client"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/model"
	observer "github.com/synycboom/bsc-evm-compatible-bridge-core/observer"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/recorder"
	sengine "github.com/synycboom/bsc-evm-compatible-bridge-core/swap-engine"
	spengine "github.com/synycboom/bsc-evm-compatible-bridge-core/swap-pair-engine"
	erc1155token "github.com/synycboom/bsc-evm-compatible-bridge-core/token/erc1155"
	erc721token "github.com/synycboom/bsc-evm-compatible-bridge-core/token/erc721"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

const (
	flagConfigType         = "config-type"
	flagConfigAwsRegion    = "aws-region"
	flagConfigAwsSecretKey = "aws-secret-key"
	flagConfigPath         = "config-path"
)

const (
	ConfigTypeLocal = "local"
	ConfigTypeAws   = "aws"
)

func initFlags() {
	flag.String(flagConfigPath, "", "config path")
	flag.String(flagConfigType, "", "config type, local or aws")
	flag.String(flagConfigAwsRegion, "", "aws s3 region")
	flag.String(flagConfigAwsSecretKey, "", "aws s3 secret key")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(fmt.Sprintf("bind flags error, err=%s", err))
	}
}

func printUsage() {
	fmt.Print("usage: ./swap --config-type [local or aws] --config-path config_file_path\n")
}

func main() {
	initFlags()

	configType := viper.GetString(flagConfigType)
	if configType == "" {
		printUsage()
		return
	}

	if configType != ConfigTypeAws && configType != ConfigTypeLocal {
		printUsage()
		return
	}

	var config *util.Config
	if configType == ConfigTypeAws {
		awsSecretKey := viper.GetString(flagConfigAwsSecretKey)
		if awsSecretKey == "" {
			printUsage()
			return
		}

		awsRegion := viper.GetString(flagConfigAwsRegion)
		if awsRegion == "" {
			printUsage()
			return
		}

		configContent, err := util.GetSecret(awsSecretKey, awsRegion)
		if err != nil {
			fmt.Printf("get aws config error, err=%s", err.Error())
			return
		}
		config = util.ParseConfigFromJson(configContent)
	} else {
		configFilePath := viper.GetString(flagConfigPath)
		if configFilePath == "" {
			printUsage()
			return
		}
		config = util.ParseConfigFromFile(configFilePath)
	}
	config.Validate()

	util.InitLogger(config.LogConfig)
	util.InitTgAlerter(config.AlertConfig)

	mysqlConn := mysql.New(mysql.Config{
		DSN: config.DBConfig.DSN,
	})
	db, err := gorm.Open(mysqlConn, &gorm.Config{
		Logger: logger.Default.LogMode(dbLogLevel(config.DBConfig.LogLevel)),
	})
	if err != nil {
		panic(errors.Wrap(err, "[main]: open db error"))
	}

	model.InitTables(db)

	erc721SwapAgents := make(map[string]erc721agent.SwapAgent)
	erc721SwapAgentAddresses := make(map[string]common.Address)
	erc721Tokens := make(map[string]erc721token.IToken)
	erc1155SwapAgents := make(map[string]erc1155agent.SwapAgent)
	erc1155SwapAgentAddresses := make(map[string]common.Address)
	erc1155Tokens := make(map[string]erc1155token.IToken)
	clients := make(map[string]client.ETHClient)
	for _, c := range config.ChainConfigs {
		ec, err := ethclient.Dial(c.Provider)
		if err != nil {
			panic(errors.Wrap(err, "[main]: new eth client error"))
		}

		erc721SwapAgentAddr := common.HexToAddress(c.ERC721SwapAgentAddr)
		erc721SwapAgent, err := contractabi.NewERC721SwapAgent(erc721SwapAgentAddr, ec)
		if err != nil {
			panic(errors.Wrap(err, "[main]: failed to create ERC721 swap agent"))
		}

		erc1155SwapAgentAddr := common.HexToAddress(c.ERC1155SwapAgentAddr)
		erc1155SwapAgent, err := contractabi.NewERC1155SwapAgent(erc1155SwapAgentAddr, ec)
		if err != nil {
			panic(errors.Wrap(err, "[main]: failed to create ERC1155 swap agent"))
		}

		clients[c.ID] = client.NewClient(ec)
		erc721Tokens[c.ID] = erc721token.NewToken(ec)
		erc721SwapAgents[c.ID] = erc721SwapAgent
		erc721SwapAgentAddresses[c.ID] = erc721SwapAgentAddr

		erc1155Tokens[c.ID] = erc1155token.NewToken(ec)
		erc1155SwapAgents[c.ID] = erc1155SwapAgent
		erc1155SwapAgentAddresses[c.ID] = erc1155SwapAgentAddr
	}

	recorders := make(map[string]recorder.IRecorder)
	for _, c := range config.ChainConfigs {
		chainID := util.StrToBigInt(c.ID)
		if chainID.Cmp(big.NewInt(0)) == 0 {
			panic(errors.New("[main]: chain id is 0"))
		}

		recorders[c.ID] = recorder.NewRecorder(&recorder.Config{
			ChainID:   chainID,
			ChainName: c.Name,
			HMACKey:   config.KeyManagerConfig.HMACKey,
		}, &recorder.Dependencies{
			Client:           clients,
			DB:               db.Session(&gorm.Session{}),
			ERC721SwapAgent:  erc721SwapAgents,
			ERC721Token:      erc721Tokens,
			ERC1155SwapAgent: erc1155SwapAgents,
			ERC1155Token:     erc1155Tokens,
		})
	}

	for _, c := range config.ChainConfigs {
		chainID := util.StrToBigInt(c.ID)

		// TODO: implement SwapAgent instance and implement mutex lock to prevent multiple calls
		// TODO: send tg when logging has error
		ob := observer.NewObserver(&observer.Config{
			StartHeight:        c.StartHeight,
			ConfirmNum:         c.ConfirmNum,
			FetchInterval:      time.Duration(c.ObserverFetchInterval) * time.Second,
			BlockUpdateTimeout: time.Duration(config.AlertConfig.BlockUpdateTimeout) * time.Second,
		}, &observer.Dependencies{
			DB:       db.Session(&gorm.Session{}),
			Recorder: recorders[c.ID],
		})
		ob.Start()

		e := spengine.NewEngine(&spengine.Config{
			ChainID:                   chainID,
			ConfirmNum:                c.ConfirmNum,
			ExplorerURL:               c.ExplorerUrl,
			PrivateKey:                c.PrivateKey,
			MaxTrackRetry:             c.MaxTrackRetry,
			ERC721SwapAgentAddresses:  erc721SwapAgentAddresses,
			ERC1155SwapAgentAddresses: erc1155SwapAgentAddresses,
		}, &spengine.Dependencies{
			Client:           clients,
			DB:               db.Session(&gorm.Session{}),
			Recorder:         recorders,
			ERC721SwapAgent:  erc721SwapAgents,
			ERC1155SwapAgent: erc1155SwapAgents,
		})
		e.Start()

		se := sengine.NewEngine(&sengine.Config{
			ChainID:                   chainID,
			ConfirmNum:                c.ConfirmNum,
			ExplorerURL:               c.ExplorerUrl,
			PrivateKey:                c.PrivateKey,
			MaxTrackRetry:             c.MaxTrackRetry,
			ERC721SwapAgentAddresses:  erc721SwapAgentAddresses,
			ERC1155SwapAgentAddresses: erc1155SwapAgentAddresses,
		}, &sengine.Dependencies{
			Client:           clients,
			DB:               db.Session(&gorm.Session{}),
			Recorder:         recorders,
			ERC721SwapAgent:  erc721SwapAgents,
			ERC721Token:      erc721Tokens,
			ERC1155SwapAgent: erc1155SwapAgents,
			ERC1155Token:     erc1155Tokens,
		})
		se.Start()
	}

	select {}
}

func dbLogLevel(level string) logger.LogLevel {
	switch level {
	case "SILENT":
		return logger.Silent
	case "ERROR":
		return logger.Error
	case "WARN":
		return logger.Warn
	case "INFO":
		return logger.Info
	}

	return logger.Warn
}
