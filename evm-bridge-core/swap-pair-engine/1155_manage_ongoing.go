package engine

import (
	"github.com/pkg/errors"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/erc1155"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

func (e *Engine) manageERC1155OngoingRegistration() {
	fromChainID := e.chainID()
	ss, err := e.queryERC1155SwapPair(fromChainID, []erc1155.SwapPairState{
		erc1155.SwapPairStateRegistrationOngoing,
	})
	if err != nil {
		util.Logger.Error(errors.Wrap(err, "[Engine.manageERC1155OngoingRegistration]: failed to query onging SwapPairs"))
		return
	}

	ss, err = e.filterERC1155ConfirmedRegisterEvents(ss)
	if err != nil {
		util.Logger.Error(errors.Wrap(err, "[Engine.manageERC1155OngoingRegistration]: failed to filter confirmed SwapPairs"))
		return
	}

	if len(ss) == 0 {
		return
	}

	ids := make([]string, len(ss))
	for idx, s := range ss {
		ids[idx] = s.ID
	}

	if err := e.deps.DB.Model(&ss).Where("id in ?", ids).Updates(map[string]interface{}{
		"state": erc1155.SwapPairStateRegistrationConfirmed,
	}).Error; err != nil {
		util.Logger.Error(
			errors.Wrapf(err, "[Engine.manageERC1155OngoingRegistration]: failed to update state '%s'", erc1155.SwapPairStateRegistrationConfirmed),
		)
	}

	for _, s := range ss {
		util.Logger.Infof("[Engine.manageERC1155OngoingRegistration]: updated SwapPair %s state to '%s'", s.ID, s.State)
	}
}
