package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	ethcom "github.com/ethereum/go-ethereum/common"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/common"
)

type Config struct {
	KeyManagerConfig KeyManagerConfig `json:"key_manager_config"`
	DBConfig         DBConfig         `json:"db_config"`
	ChainConfigs     []ChainConfig    `json:"chain_configs"`
	LogConfig        LogConfig        `json:"log_config"`
	AlertConfig      AlertConfig      `json:"alert_config"`
	AdminConfig      AdminConfig      `json:"admin_config"`
}

func (cfg *Config) Validate() {
	cfg.DBConfig.Validate()
	cfg.LogConfig.Validate()
	cfg.AlertConfig.Validate()

	ids := make(map[string]struct{})
	for _, c := range cfg.ChainConfigs {
		c.Validate()

		if _, ok := ids[c.ID]; ok {
			panic("chain id is duplicated")
		}

		ids[c.ID] = struct{}{}
	}
}

type AlertConfig struct {
	TelegramBotId  string `json:"telegram_bot_id"`
	TelegramChatId string `json:"telegram_chat_id"`

	BlockUpdateTimeout int64 `json:"block_update_timeout"`
}

func (cfg AlertConfig) Validate() {
	if cfg.BlockUpdateTimeout <= 0 {
		panic("block_update_timeout should be larger than 0")
	}
}

type KeyManagerConfig struct {
	KeyType       string `json:"key_type"`
	AWSRegion     string `json:"aws_region"`
	AWSSecretName string `json:"aws_secret_name"`
	HMACKey       string `json:"hmac_key"`
}

type KeyConfig struct {
	HMACKey            string `json:"hmac_key"`
	PrivateKey         string `json:"private_key"`
	ETHChainPrivateKey string `json:"eth_private_key"`
	AdminApiKey        string `json:"admin_api_key"`
	AdminSecretKey     string `json:"admin_secret_key"`
}

func (cfg KeyManagerConfig) Validate() {
	if cfg.KeyType == common.LocalPrivateKey && len(cfg.HMACKey) == 0 {
		panic("missing local hmac key")
	}
	// if cfg.KeyType == common.LocalPrivateKey && len(cfg.LocalPrivateKey) == 0 {
	// 	panic("missing local source chain private key")
	// }
	// if cfg.KeyType == common.LocalPrivateKey && len(cfg.LocalETHChainPrivateKey) == 0 {
	// 	panic("missing local destination chain private key")
	// }

	// if cfg.KeyType == common.LocalPrivateKey && len(cfg.LocalAdminApiKey) == 0 {
	// 	panic("missing local admin api key")
	// }

	// if cfg.KeyType == common.LocalPrivateKey && len(cfg.LocalAdminSecretKey) == 0 {
	// 	panic("missing local admin secret key")
	// }

	if cfg.KeyType == common.AWSPrivateKey && (cfg.AWSRegion == "" || cfg.AWSSecretName == "") {
		panic("Missing aws key region or name")
	}
}

type TokenSecretKey struct {
	Symbol             string `json:"symbol"`
	PrivateKey         string `json:"private_key"`
	ETHChainPrivateKey string `json:"destination_private_key"`
}

type DBConfig struct {
	LogLevel string `json:"log_level"`
	DSN      string `json:"dsn"`
}

func (cfg DBConfig) Validate() {
	if cfg.DSN == "" {
		panic("db path should not be empty")
	}
}

type ChainConfig struct {
	BalanceAlertThreshold  string `json:"balance_alert_threshold"`
	BalanceMonitorInterval int64  `json:"balance_monitor_interval"`
	ID                     string `json:"id"`
	Name                   string `json:"name"`
	ObserverFetchInterval  int64  `json:"observer_fetch_interval"`
	StartHeight            int64  `json:"start_height"`
	PrivateKey             string `json:"private_key"`
	Provider               string `json:"provider"`
	ConfirmNum             int64  `json:"confirm_num"`
	ERC721SwapAgentAddr    string `json:"erc_721_swap_agent_addr"`
	ERC1155SwapAgentAddr   string `json:"erc_1155_swap_agent_addr"`
	ExplorerUrl            string `json:"explorer_url"`
	MaxTrackRetry          int64  `json:"max_track_retry"`
	WaitMilliSecBetweenTx  int64  `json:"wait_milli_sec_between_tx"`
}

func (cfg ChainConfig) Validate() {
	if cfg.Name == "" {
		panic("name should not be empty")
	}
	if cfg.StartHeight < 0 {
		panic("start_height should not be less than 0")
	}
	if cfg.Provider == "" {
		panic("provider should not be empty")
	}
	if cfg.ConfirmNum <= 0 {
		panic("confirm_num should be larger than 0")
	}
	if !ethcom.IsHexAddress(cfg.ERC721SwapAgentAddr) {
		panic(fmt.Sprintf("invalid swap_contract_addr: %s", cfg.ERC721SwapAgentAddr))
	}
	if cfg.MaxTrackRetry <= 0 {
		panic("max_track_retry should be larger than 0")
	}
}

type LogConfig struct {
	Level                        string `json:"level"`
	Filename                     string `json:"filename"`
	MaxFileSizeInMB              int    `json:"max_file_size_in_mb"`
	MaxBackupsOfLogFiles         int    `json:"max_backups_of_log_files"`
	MaxAgeToRetainLogFilesInDays int    `json:"max_age_to_retain_log_files_in_days"`
	UseConsoleLogger             bool   `json:"use_console_logger"`
	UseFileLogger                bool   `json:"use_file_logger"`
	Compress                     bool   `json:"compress"`
}

func (cfg LogConfig) Validate() {
	if cfg.UseFileLogger {
		if cfg.Filename == "" {
			panic("filename should not be empty if use file logger")
		}
		if cfg.MaxFileSizeInMB <= 0 {
			panic("max_file_size_in_mb should be larger than 0 if use file logger")
		}
		if cfg.MaxBackupsOfLogFiles <= 0 {
			panic("max_backups_off_log_files should be larger than 0 if use file logger")
		}
	}
}

type AdminConfig struct {
	ListenAddr string `json:"listen_addr"`
}

func ParseConfigFromFile(filePath string) *Config {
	bz, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(bz, &config); err != nil {
		panic(err)
	}

	return &config
}

func ParseConfigFromJson(content string) *Config {
	var config Config
	if err := json.Unmarshal([]byte(content), &config); err != nil {
		panic(err)
	}

	return &config
}
