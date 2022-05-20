package engine

import (
	"math"

	"github.com/ethereum/go-ethereum/core"
	"github.com/pkg/errors"

	"github.com/synycboom/bsc-evm-compatible-bridge-core/model/erc721"
	"github.com/synycboom/bsc-evm-compatible-bridge-core/util"
)

func (e *Engine) manageERC721ConfirmedRegitration() {
	fromChainID := e.chainID()
	ss, err := e.queryERC721SwapPair(fromChainID, []erc721.SwapPairState{
		erc721.SwapPairStateRegistrationConfirmed,
	})
	if err != nil {
		util.Logger.Error(errors.Wrap(err, "[Engine.manageERC721ConfirmedRegitration]: failed to query confirmed SwapPairs"))
		return
	}

	for _, s := range ss {
		txHash, err := e.generateERC721TxHash(s)
		if err != nil {
			// this error might comes from gas estimation, so it means we cannot send the real tx to the chain
			util.Logger.Warningf("[Engine.manageERC721ConfirmedRegitration]: failed to dry run tx of SwapPair %s", s.ID)

			s.State = erc721.SwapPairStateCreationTxDryRunFailed
			s.MessageLog = err.Error()
			if err := e.deps.DB.Save(s).Error; err != nil {
				util.Logger.Error(
					errors.Wrapf(err, "[Engine.manageERC721ConfirmedRegitration]: failed to update SwapPair %s to '%s' state", s.ID, s.State),
				)
			}

			continue
		}

		// We save the tx as our checkpoint to probe the stats later
		// It tells that this tx might be sent or might not, but it is okay
		// We will set the state to failed later
		s.State = erc721.SwapPairStateCreationTxCreated
		s.CreateTxHash = txHash
		s.CreateHeight = math.MaxInt64
		if err := e.deps.DB.Save(s).Error; err != nil {
			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC721ConfirmedRegitration]: failed to update SwapPair %s to '%s' state", s.ID, s.State),
			)

			continue
		}

		util.Logger.Infof(
			"[Engine.manageERC721ConfirmedRegitration]: sent dry run tx to chain id %s, %s",
			e.chainID(),
			txHash,
		)

		request, err := e.sendERC721CreatePairRequest(s, false)
		if err != nil {
			// retry when a transaction is attempted to be replaced
			// with a different one without the required price bump.
			if errors.Cause(err).Error() == core.ErrReplaceUnderpriced.Error() {
				s.State = erc721.SwapPairStateRegistrationConfirmed
				s.MessageLog = err.Error()
				if dbErr := e.deps.DB.Save(s).Error; dbErr != nil {
					util.Logger.Error(
						errors.Wrapf(dbErr, "[Engine.manageERC721ConfirmedRegitration]: failed to update SwapPair %s to '%s' state", s.ID, s.State),
					)

					continue
				}
			}

			s.State = erc721.SwapPairStateCreationTxFailed
			s.MessageLog = err.Error()
			if dbErr := e.deps.DB.Save(s).Error; dbErr != nil {
				util.Logger.Error(
					errors.Wrapf(dbErr, "[Engine.manageERC721ConfirmedRegitration]: failed to update SwapPair %s to '%s' state", s.ID, s.State),
				)

				continue
			}

			util.Logger.Error(
				errors.Wrapf(err, "[Engine.manageERC721ConfirmedRegitration]: failed to send a real tx %s of SwapPair %s", s.CreateTxHash, s.ID),
			)

			continue
		}

		util.Logger.Infof(
			"[Engine.manageERC721ConfirmedRegitration]: sent tx to chain id %s, %s/%s",
			e.chainID(),
			e.conf.ExplorerURL,
			request.Hash().String(),
		)

		// update tx hash again in case there are some parameters might change tx hash
		// for example, gas limit which comes from estimation
		s.CreateTxHash = request.Hash().String()
		if dbErr := e.deps.DB.Save(s).Error; dbErr != nil {
			util.Logger.Error(
				errors.Wrapf(dbErr, "[Engine.manageERC721ConfirmedRegitration]: failed to update SwapPair %s creation tx hash %s right after sending out", s.ID, s.CreateTxHash),
			)

			continue
		}
	}
}
