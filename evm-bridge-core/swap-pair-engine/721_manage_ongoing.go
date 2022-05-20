package engine

import (
	"github.com/pkg/errors"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/erc721"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

func (e *Engine) manageERC721OngoingRegistration() {
	fromChainID := e.chainID()
	ss, err := e.queryERC721SwapPair(fromChainID, []erc721.SwapPairState{
		erc721.SwapPairStateRegistrationOngoing,
	})
	if err != nil {
		util.Logger.Error(errors.Wrap(err, "[Engine.manageERC721OngoingRegistration]: failed to query onging SwapPairs"))
		return
	}

	ss, err = e.filterERC721ConfirmedRegisterEvents(ss)
	if err != nil {
		util.Logger.Error(errors.Wrap(err, "[Engine.manageERC721OngoingRegistration]: failed to filter confirmed SwapPairs"))
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
		"state": erc721.SwapPairStateRegistrationConfirmed,
	}).Error; err != nil {
		util.Logger.Error(
			errors.Wrapf(err, "[Engine.manageERC721OngoingRegistration]: failed to update state '%s'", erc721.SwapPairStateRegistrationConfirmed),
		)
	}

	for _, s := range ss {
		util.Logger.Infof("[Engine.manageERC721OngoingRegistration]: updated SwapPair %s state to '%s'", s.ID, s.State)
	}
}
