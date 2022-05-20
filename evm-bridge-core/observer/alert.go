package observer

import (
	"fmt"
	"time"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/common"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

// Alert sends alerts to tg group if there is no new block fetched in a specific time
func (o *Observer) Alert() {
	for {
		curOtherChainBlockLog, err := o.GetCurrentBlockLog()
		if err != nil {
			util.Logger.Errorf("[Observer.Alert]: get current block log error, err=%s", err.Error())
			time.Sleep(common.ObserverAlertInterval)

			continue
		}

		if curOtherChainBlockLog.Height > 0 {
			if time.Now().Unix()-curOtherChainBlockLog.CreateTime.Unix() > int64(o.conf.BlockUpdateTimeout.Seconds()) {
				msg := fmt.Sprintf("[Observer.Alert]: last block fetched at %s, chain=%s, height=%d",
					time.Unix(curOtherChainBlockLog.CreateTime.Unix(), 0).String(), o.deps.Recorder.ChainID(), curOtherChainBlockLog.Height)
				util.SendTelegramMessage(msg)
			}
		}

		time.Sleep(common.ObserverAlertInterval)
	}
}
