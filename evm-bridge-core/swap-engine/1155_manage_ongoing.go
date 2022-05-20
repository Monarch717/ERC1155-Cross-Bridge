package engine

import (
	"github.com/pkg/errors"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/erc1155"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

func (e *Engine) manageERC1155OngoingRequest() {
	fromChainID := e.chainID()
	ss, err := e.queryERC1155Swap(fromChainID, []erc1155.SwapState{
		erc1155.SwapStateRequestOngoing,
	})
	if err != nil {
		util.Logger.Error(errors.Wrap(err, "[Engine.manageERC1155OngoingRequest]: failed to query onging Swaps"))
		return
	}

	// Fill required information without updating to DB
	if err := e.fillERC1155RequiredInfo(ss); err != nil {
		util.Logger.Error(errors.Wrap(err, "[Engine.manageERC1155OngoingRequest]: failed to fill destination"))
		return
	}

	// Separate ready Swaps, pending Swaps, and rejected Swaps
	ss, pp, rr := e.separateERC1155SwapEvents(ss)
	for _, r := range rr {
		r.State = erc1155.SwapStateRequestRejected
		if err := e.deps.DB.Save(&r).Error; err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC1155OngoingRequest]: failed to update Swap %s to state '%s'", r.ID, r.State),
			)
		}
	}
	for _, p := range pp {
		if err := e.deps.DB.Save(&p).Error; err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC1155OngoingRequest]: failed to update Swap %s", p.ID),
			)
		}
	}

	ss, err = e.filterERC1155ConfirmedSwapEvents(ss)
	if err != nil {
		util.Logger.Error(errors.Wrap(err, "[Engine.manageERC1155OngoingRequest]: failed to filter confirmed Swaps"))
		return
	}

	for _, s := range ss {
		s.State = erc1155.SwapStateRequestConfirmed
		if err := e.deps.DB.Save(&s).Error; err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC1155OngoingRequest]: failed to update Swap %s to state '%s'", s.ID, s.State),
			)
		}
	}
}
