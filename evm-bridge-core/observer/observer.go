package observer

import (
	"time"

	"gorm.io/gorm"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/recorder"
)

type Config struct {
	StartHeight        int64
	ConfirmNum         int64
	FetchInterval      time.Duration
	BlockUpdateTimeout time.Duration
}

type Dependencies struct {
	DB       *gorm.DB
	Recorder recorder.IRecorder
}

type Observer struct {
	conf *Config
	deps *Dependencies
}

// NewObserver returns the observer instance
func NewObserver(c *Config, d *Dependencies) *Observer {
	return &Observer{
		conf: c,
		deps: d,
	}
}

// Start starts the routines of observer
func (o *Observer) Start() {
	go o.Update()
	go o.Prune()
	go o.Alert()
}
