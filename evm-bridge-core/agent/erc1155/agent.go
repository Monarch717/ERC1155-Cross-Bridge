package agent

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	contractabi "github.com/synycboom/bsc-evm-compatible-bridge-core/abi"
)

type SwapAgent interface {
	FilterSwapPairRegister(
		opts *bind.FilterOpts,
		sponsor []common.Address,
		tokenAddress []common.Address,
	) (*contractabi.ERC1155SwapAgentSwapPairRegisterIterator, error)

	FilterSwapPairCreated(
		opts *bind.FilterOpts,
		registerTxHash [][32]byte,
		fromTokenAddr []common.Address,
		mirroredTokenAddr []common.Address,
	) (*contractabi.ERC1155SwapAgentSwapPairCreatedIterator, error)

	CreateSwapPair(
		opts *bind.TransactOpts,
		registerTxHash [32]byte,
		fromTokenAddr common.Address,
		fromChainId *big.Int,
		uri string,
	) (*types.Transaction, error)

	FilterSwapStarted(
		opts *bind.FilterOpts,
		tokenAddr []common.Address,
		sender []common.Address,
		recipient []common.Address,
	) (*contractabi.ERC1155SwapAgentSwapStartedIterator, error)

	FilterSwapFilled(
		opts *bind.FilterOpts,
		swapTxHash [][32]byte,
		fromTokenAddr []common.Address,
		recipient []common.Address,
	) (*contractabi.ERC1155SwapAgentSwapFilledIterator, error)

	Fill(
		opts *bind.TransactOpts,
		swapTxHash [32]byte,
		fromTokenAddr common.Address,
		recipient common.Address,
		fromChainId *big.Int,
		ids []*big.Int,
		amounts []*big.Int,
	) (*types.Transaction, error)

	FilterBackwardSwapStarted(
		opts *bind.FilterOpts,
		mirroredTokenAddr []common.Address,
		sender []common.Address,
		recipient []common.Address,
	) (*contractabi.ERC1155SwapAgentBackwardSwapStartedIterator, error)

	FilterBackwardSwapFilled(
		opts *bind.FilterOpts,
		swapTxHash [][32]byte,
		tokenAddr []common.Address,
		recipient []common.Address,
	) (*contractabi.ERC1155SwapAgentBackwardSwapFilledIterator, error)
}
