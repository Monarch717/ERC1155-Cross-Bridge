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
	) (*contractabi.ERC721SwapAgentSwapPairRegisterIterator, error)

	FilterSwapPairCreated(
		opts *bind.FilterOpts, registerTxHash [][32]byte,
		fromTokenAddr []common.Address,
		mirroredTokenAddr []common.Address,
	) (*contractabi.ERC721SwapAgentSwapPairCreatedIterator, error)

	CreateSwapPair(
		opts *bind.TransactOpts,
		registerTxHash [32]byte,
		fromTokenAddr common.Address,
		fromChainId *big.Int, baseURI string,
		tokenName string,
		tokenSymbol string,
	) (*types.Transaction, error)

	FilterSwapStarted(
		opts *bind.FilterOpts,
		tokenAddr []common.Address,
		sender []common.Address,
		recipient []common.Address,
	) (*contractabi.ERC721SwapAgentSwapStartedIterator, error)

	FilterSwapFilled(
		opts *bind.FilterOpts,
		swapTxHash [][32]byte,
		fromTokenAddr []common.Address,
		recipient []common.Address,
	) (*contractabi.ERC721SwapAgentSwapFilledIterator, error)

	Fill(
		opts *bind.TransactOpts,
		swapTxHash [32]byte,
		fromTokenAddr common.Address,
		recipient common.Address,
		fromChainId *big.Int,
		tokenId *big.Int,
		tokenURI string,
	) (*types.Transaction, error)

	FilterBackwardSwapStarted(
		opts *bind.FilterOpts,
		mirroredTokenAddr []common.Address,
		sender []common.Address,
		recipient []common.Address,
	) (*contractabi.ERC721SwapAgentBackwardSwapStartedIterator, error)

	FilterBackwardSwapFilled(
		opts *bind.FilterOpts,
		swapTxHash [][32]byte,
		tokenAddr []common.Address,
		recipient []common.Address,
	) (*contractabi.ERC721SwapAgentBackwardSwapFilledIterator, error)
}
