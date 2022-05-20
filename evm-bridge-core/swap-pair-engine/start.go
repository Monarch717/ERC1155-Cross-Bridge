package engine

import (
	"reflect"
	"runtime"
	"time"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

const (
	watchEventDelay = time.Duration(5) * time.Second
)

func (e *Engine) Start() {
	// ERC721
	go e.run(e.manageERC721OngoingRegistration, watchRegisterEventDelay)
	go e.run(e.manageERC721ConfirmedRegitration, watchRegisterEventDelay)
	go e.run(e.manageERC721TxCreatedRegistration, watchRegisterEventDelay)
	go e.run(e.manageERC721TxSentRegistration, watchRegisterEventDelay)

	// ERC1155
	go e.run(e.manageERC1155OngoingRegistration, watchRegisterEventDelay)
	go e.run(e.manageERC1155ConfirmedRegitration, watchRegisterEventDelay)
	go e.run(e.manageERC1155TxCreatedRegistration, watchRegisterEventDelay)
	go e.run(e.manageERC1155TxSentRegistration, watchRegisterEventDelay)
}

func (e *Engine) run(fn func(), delay time.Duration) {
	fnName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	if delay.Seconds() == 0 {
		delay = watchEventDelay
	}

	for {
		time.Sleep(watchEventDelay)

		if e.deps.Recorder[e.chainID()].LatestBlockCached() == nil {
			util.Logger.Infof("[Engine.run][%s]: no latest block cache found for chain id %s", fnName, e.chainID())

			continue
		}

		fn()
	}
}
