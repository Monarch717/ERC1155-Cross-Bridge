// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERC721SwapAgentMetaData contains all meta data concerning the ERC721SwapAgent contract.
var ERC721SwapAgentMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"swapTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"BackwardSwapFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mirroredTokenAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"BackwardSwapStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"swapTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromTokenAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mirroredTokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"SwapFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"registerTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromTokenAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mirroredTokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"}],\"name\":\"SwapPairCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"SwapPairRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"SwapStarted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"fromTokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"name\":\"createSwapPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"swapTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"fromTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"fill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filledSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"registerSwapPair\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registeredToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapMappingIncoming\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapMappingOutgoing\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC721SwapAgentABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721SwapAgentMetaData.ABI instead.
var ERC721SwapAgentABI = ERC721SwapAgentMetaData.ABI

// ERC721SwapAgent is an auto generated Go binding around an Ethereum contract.
type ERC721SwapAgent struct {
	ERC721SwapAgentCaller     // Read-only binding to the contract
	ERC721SwapAgentTransactor // Write-only binding to the contract
	ERC721SwapAgentFilterer   // Log filterer for contract events
}

// ERC721SwapAgentCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721SwapAgentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721SwapAgentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721SwapAgentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721SwapAgentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721SwapAgentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721SwapAgentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721SwapAgentSession struct {
	Contract     *ERC721SwapAgent  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721SwapAgentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721SwapAgentCallerSession struct {
	Contract *ERC721SwapAgentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ERC721SwapAgentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721SwapAgentTransactorSession struct {
	Contract     *ERC721SwapAgentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ERC721SwapAgentRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721SwapAgentRaw struct {
	Contract *ERC721SwapAgent // Generic contract binding to access the raw methods on
}

// ERC721SwapAgentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721SwapAgentCallerRaw struct {
	Contract *ERC721SwapAgentCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721SwapAgentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721SwapAgentTransactorRaw struct {
	Contract *ERC721SwapAgentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721SwapAgent creates a new instance of ERC721SwapAgent, bound to a specific deployed contract.
func NewERC721SwapAgent(address common.Address, backend bind.ContractBackend) (*ERC721SwapAgent, error) {
	contract, err := bindERC721SwapAgent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721SwapAgent{ERC721SwapAgentCaller: ERC721SwapAgentCaller{contract: contract}, ERC721SwapAgentTransactor: ERC721SwapAgentTransactor{contract: contract}, ERC721SwapAgentFilterer: ERC721SwapAgentFilterer{contract: contract}}, nil
}

// NewERC721SwapAgentCaller creates a new read-only instance of ERC721SwapAgent, bound to a specific deployed contract.
func NewERC721SwapAgentCaller(address common.Address, caller bind.ContractCaller) (*ERC721SwapAgentCaller, error) {
	contract, err := bindERC721SwapAgent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721SwapAgentCaller{contract: contract}, nil
}

// NewERC721SwapAgentTransactor creates a new write-only instance of ERC721SwapAgent, bound to a specific deployed contract.
func NewERC721SwapAgentTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721SwapAgentTransactor, error) {
	contract, err := bindERC721SwapAgent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721SwapAgentTransactor{contract: contract}, nil
}

// NewERC721SwapAgentFilterer creates a new log filterer instance of ERC721SwapAgent, bound to a specific deployed contract.
func NewERC721SwapAgentFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721SwapAgentFilterer, error) {
	contract, err := bindERC721SwapAgent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721SwapAgentFilterer{contract: contract}, nil
}

// bindERC721SwapAgent binds a generic wrapper to an already deployed contract.
func bindERC721SwapAgent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721SwapAgentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721SwapAgent *ERC721SwapAgentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721SwapAgent.Contract.ERC721SwapAgentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721SwapAgent *ERC721SwapAgentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.ERC721SwapAgentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721SwapAgent *ERC721SwapAgentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.ERC721SwapAgentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721SwapAgent *ERC721SwapAgentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721SwapAgent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721SwapAgent *ERC721SwapAgentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721SwapAgent *ERC721SwapAgentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.contract.Transact(opts, method, params...)
}

// FilledSwap is a free data retrieval call binding the contract method 0xa86894ca.
//
// Solidity: function filledSwap(bytes32 ) view returns(bool)
func (_ERC721SwapAgent *ERC721SwapAgentCaller) FilledSwap(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ERC721SwapAgent.contract.Call(opts, &out, "filledSwap", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FilledSwap is a free data retrieval call binding the contract method 0xa86894ca.
//
// Solidity: function filledSwap(bytes32 ) view returns(bool)
func (_ERC721SwapAgent *ERC721SwapAgentSession) FilledSwap(arg0 [32]byte) (bool, error) {
	return _ERC721SwapAgent.Contract.FilledSwap(&_ERC721SwapAgent.CallOpts, arg0)
}

// FilledSwap is a free data retrieval call binding the contract method 0xa86894ca.
//
// Solidity: function filledSwap(bytes32 ) view returns(bool)
func (_ERC721SwapAgent *ERC721SwapAgentCallerSession) FilledSwap(arg0 [32]byte) (bool, error) {
	return _ERC721SwapAgent.Contract.FilledSwap(&_ERC721SwapAgent.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC721SwapAgent *ERC721SwapAgentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC721SwapAgent.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC721SwapAgent *ERC721SwapAgentSession) Owner() (common.Address, error) {
	return _ERC721SwapAgent.Contract.Owner(&_ERC721SwapAgent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC721SwapAgent *ERC721SwapAgentCallerSession) Owner() (common.Address, error) {
	return _ERC721SwapAgent.Contract.Owner(&_ERC721SwapAgent.CallOpts)
}

// RegisteredToken is a free data retrieval call binding the contract method 0x0b4f43c1.
//
// Solidity: function registeredToken(uint256 , address ) view returns(bool)
func (_ERC721SwapAgent *ERC721SwapAgentCaller) RegisteredToken(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC721SwapAgent.contract.Call(opts, &out, "registeredToken", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredToken is a free data retrieval call binding the contract method 0x0b4f43c1.
//
// Solidity: function registeredToken(uint256 , address ) view returns(bool)
func (_ERC721SwapAgent *ERC721SwapAgentSession) RegisteredToken(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _ERC721SwapAgent.Contract.RegisteredToken(&_ERC721SwapAgent.CallOpts, arg0, arg1)
}

// RegisteredToken is a free data retrieval call binding the contract method 0x0b4f43c1.
//
// Solidity: function registeredToken(uint256 , address ) view returns(bool)
func (_ERC721SwapAgent *ERC721SwapAgentCallerSession) RegisteredToken(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _ERC721SwapAgent.Contract.RegisteredToken(&_ERC721SwapAgent.CallOpts, arg0, arg1)
}

// SwapMappingIncoming is a free data retrieval call binding the contract method 0xec686704.
//
// Solidity: function swapMappingIncoming(uint256 , address ) view returns(address)
func (_ERC721SwapAgent *ERC721SwapAgentCaller) SwapMappingIncoming(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ERC721SwapAgent.contract.Call(opts, &out, "swapMappingIncoming", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapMappingIncoming is a free data retrieval call binding the contract method 0xec686704.
//
// Solidity: function swapMappingIncoming(uint256 , address ) view returns(address)
func (_ERC721SwapAgent *ERC721SwapAgentSession) SwapMappingIncoming(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _ERC721SwapAgent.Contract.SwapMappingIncoming(&_ERC721SwapAgent.CallOpts, arg0, arg1)
}

// SwapMappingIncoming is a free data retrieval call binding the contract method 0xec686704.
//
// Solidity: function swapMappingIncoming(uint256 , address ) view returns(address)
func (_ERC721SwapAgent *ERC721SwapAgentCallerSession) SwapMappingIncoming(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _ERC721SwapAgent.Contract.SwapMappingIncoming(&_ERC721SwapAgent.CallOpts, arg0, arg1)
}

// SwapMappingOutgoing is a free data retrieval call binding the contract method 0x0d43d992.
//
// Solidity: function swapMappingOutgoing(uint256 , address ) view returns(address)
func (_ERC721SwapAgent *ERC721SwapAgentCaller) SwapMappingOutgoing(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ERC721SwapAgent.contract.Call(opts, &out, "swapMappingOutgoing", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapMappingOutgoing is a free data retrieval call binding the contract method 0x0d43d992.
//
// Solidity: function swapMappingOutgoing(uint256 , address ) view returns(address)
func (_ERC721SwapAgent *ERC721SwapAgentSession) SwapMappingOutgoing(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _ERC721SwapAgent.Contract.SwapMappingOutgoing(&_ERC721SwapAgent.CallOpts, arg0, arg1)
}

// SwapMappingOutgoing is a free data retrieval call binding the contract method 0x0d43d992.
//
// Solidity: function swapMappingOutgoing(uint256 , address ) view returns(address)
func (_ERC721SwapAgent *ERC721SwapAgentCallerSession) SwapMappingOutgoing(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _ERC721SwapAgent.Contract.SwapMappingOutgoing(&_ERC721SwapAgent.CallOpts, arg0, arg1)
}

// CreateSwapPair is a paid mutator transaction binding the contract method 0x9df52edd.
//
// Solidity: function createSwapPair(bytes32 registerTxHash, address fromTokenAddr, uint256 fromChainId, string baseURI_, string tokenName, string tokenSymbol) returns()
func (_ERC721SwapAgent *ERC721SwapAgentTransactor) CreateSwapPair(opts *bind.TransactOpts, registerTxHash [32]byte, fromTokenAddr common.Address, fromChainId *big.Int, baseURI_ string, tokenName string, tokenSymbol string) (*types.Transaction, error) {
	return _ERC721SwapAgent.contract.Transact(opts, "createSwapPair", registerTxHash, fromTokenAddr, fromChainId, baseURI_, tokenName, tokenSymbol)
}

// CreateSwapPair is a paid mutator transaction binding the contract method 0x9df52edd.
//
// Solidity: function createSwapPair(bytes32 registerTxHash, address fromTokenAddr, uint256 fromChainId, string baseURI_, string tokenName, string tokenSymbol) returns()
func (_ERC721SwapAgent *ERC721SwapAgentSession) CreateSwapPair(registerTxHash [32]byte, fromTokenAddr common.Address, fromChainId *big.Int, baseURI_ string, tokenName string, tokenSymbol string) (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.CreateSwapPair(&_ERC721SwapAgent.TransactOpts, registerTxHash, fromTokenAddr, fromChainId, baseURI_, tokenName, tokenSymbol)
}

// CreateSwapPair is a paid mutator transaction binding the contract method 0x9df52edd.
//
// Solidity: function createSwapPair(bytes32 registerTxHash, address fromTokenAddr, uint256 fromChainId, string baseURI_, string tokenName, string tokenSymbol) returns()
func (_ERC721SwapAgent *ERC721SwapAgentTransactorSession) CreateSwapPair(registerTxHash [32]byte, fromTokenAddr common.Address, fromChainId *big.Int, baseURI_ string, tokenName string, tokenSymbol string) (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.CreateSwapPair(&_ERC721SwapAgent.TransactOpts, registerTxHash, fromTokenAddr, fromChainId, baseURI_, tokenName, tokenSymbol)
}

// Fill is a paid mutator transaction binding the contract method 0x79e7db59.
//
// Solidity: function fill(bytes32 swapTxHash, address fromTokenAddr, address recipient, uint256 fromChainId, uint256 tokenId, string tokenURI) returns()
func (_ERC721SwapAgent *ERC721SwapAgentTransactor) Fill(opts *bind.TransactOpts, swapTxHash [32]byte, fromTokenAddr common.Address, recipient common.Address, fromChainId *big.Int, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _ERC721SwapAgent.contract.Transact(opts, "fill", swapTxHash, fromTokenAddr, recipient, fromChainId, tokenId, tokenURI)
}

// Fill is a paid mutator transaction binding the contract method 0x79e7db59.
//
// Solidity: function fill(bytes32 swapTxHash, address fromTokenAddr, address recipient, uint256 fromChainId, uint256 tokenId, string tokenURI) returns()
func (_ERC721SwapAgent *ERC721SwapAgentSession) Fill(swapTxHash [32]byte, fromTokenAddr common.Address, recipient common.Address, fromChainId *big.Int, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.Fill(&_ERC721SwapAgent.TransactOpts, swapTxHash, fromTokenAddr, recipient, fromChainId, tokenId, tokenURI)
}

// Fill is a paid mutator transaction binding the contract method 0x79e7db59.
//
// Solidity: function fill(bytes32 swapTxHash, address fromTokenAddr, address recipient, uint256 fromChainId, uint256 tokenId, string tokenURI) returns()
func (_ERC721SwapAgent *ERC721SwapAgentTransactorSession) Fill(swapTxHash [32]byte, fromTokenAddr common.Address, recipient common.Address, fromChainId *big.Int, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.Fill(&_ERC721SwapAgent.TransactOpts, swapTxHash, fromTokenAddr, recipient, fromChainId, tokenId, tokenURI)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ERC721SwapAgent *ERC721SwapAgentTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721SwapAgent.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ERC721SwapAgent *ERC721SwapAgentSession) Initialize() (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.Initialize(&_ERC721SwapAgent.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ERC721SwapAgent *ERC721SwapAgentTransactorSession) Initialize() (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.Initialize(&_ERC721SwapAgent.TransactOpts)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_ERC721SwapAgent *ERC721SwapAgentTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _ERC721SwapAgent.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_ERC721SwapAgent *ERC721SwapAgentSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.OnERC721Received(&_ERC721SwapAgent.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_ERC721SwapAgent *ERC721SwapAgentTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.OnERC721Received(&_ERC721SwapAgent.TransactOpts, arg0, arg1, arg2, arg3)
}

// RegisterSwapPair is a paid mutator transaction binding the contract method 0x45b1ab1b.
//
// Solidity: function registerSwapPair(address tokenAddr, uint256 chainId) payable returns()
func (_ERC721SwapAgent *ERC721SwapAgentTransactor) RegisterSwapPair(opts *bind.TransactOpts, tokenAddr common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _ERC721SwapAgent.contract.Transact(opts, "registerSwapPair", tokenAddr, chainId)
}

// RegisterSwapPair is a paid mutator transaction binding the contract method 0x45b1ab1b.
//
// Solidity: function registerSwapPair(address tokenAddr, uint256 chainId) payable returns()
func (_ERC721SwapAgent *ERC721SwapAgentSession) RegisterSwapPair(tokenAddr common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.RegisterSwapPair(&_ERC721SwapAgent.TransactOpts, tokenAddr, chainId)
}

// RegisterSwapPair is a paid mutator transaction binding the contract method 0x45b1ab1b.
//
// Solidity: function registerSwapPair(address tokenAddr, uint256 chainId) payable returns()
func (_ERC721SwapAgent *ERC721SwapAgentTransactorSession) RegisterSwapPair(tokenAddr common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.RegisterSwapPair(&_ERC721SwapAgent.TransactOpts, tokenAddr, chainId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC721SwapAgent *ERC721SwapAgentTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721SwapAgent.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC721SwapAgent *ERC721SwapAgentSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.RenounceOwnership(&_ERC721SwapAgent.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC721SwapAgent *ERC721SwapAgentTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.RenounceOwnership(&_ERC721SwapAgent.TransactOpts)
}

// Swap is a paid mutator transaction binding the contract method 0xfe029156.
//
// Solidity: function swap(address tokenAddr, address recipient, uint256 tokenId, uint256 dstChainId) payable returns()
func (_ERC721SwapAgent *ERC721SwapAgentTransactor) Swap(opts *bind.TransactOpts, tokenAddr common.Address, recipient common.Address, tokenId *big.Int, dstChainId *big.Int) (*types.Transaction, error) {
	return _ERC721SwapAgent.contract.Transact(opts, "swap", tokenAddr, recipient, tokenId, dstChainId)
}

// Swap is a paid mutator transaction binding the contract method 0xfe029156.
//
// Solidity: function swap(address tokenAddr, address recipient, uint256 tokenId, uint256 dstChainId) payable returns()
func (_ERC721SwapAgent *ERC721SwapAgentSession) Swap(tokenAddr common.Address, recipient common.Address, tokenId *big.Int, dstChainId *big.Int) (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.Swap(&_ERC721SwapAgent.TransactOpts, tokenAddr, recipient, tokenId, dstChainId)
}

// Swap is a paid mutator transaction binding the contract method 0xfe029156.
//
// Solidity: function swap(address tokenAddr, address recipient, uint256 tokenId, uint256 dstChainId) payable returns()
func (_ERC721SwapAgent *ERC721SwapAgentTransactorSession) Swap(tokenAddr common.Address, recipient common.Address, tokenId *big.Int, dstChainId *big.Int) (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.Swap(&_ERC721SwapAgent.TransactOpts, tokenAddr, recipient, tokenId, dstChainId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC721SwapAgent *ERC721SwapAgentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC721SwapAgent.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC721SwapAgent *ERC721SwapAgentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.TransferOwnership(&_ERC721SwapAgent.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC721SwapAgent *ERC721SwapAgentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC721SwapAgent.Contract.TransferOwnership(&_ERC721SwapAgent.TransactOpts, newOwner)
}

// ERC721SwapAgentBackwardSwapFilledIterator is returned from FilterBackwardSwapFilled and is used to iterate over the raw logs and unpacked data for BackwardSwapFilled events raised by the ERC721SwapAgent contract.
type ERC721SwapAgentBackwardSwapFilledIterator struct {
	Event *ERC721SwapAgentBackwardSwapFilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC721SwapAgentBackwardSwapFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721SwapAgentBackwardSwapFilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC721SwapAgentBackwardSwapFilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC721SwapAgentBackwardSwapFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721SwapAgentBackwardSwapFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721SwapAgentBackwardSwapFilled represents a BackwardSwapFilled event raised by the ERC721SwapAgent contract.
type ERC721SwapAgentBackwardSwapFilled struct {
	SwapTxHash  [32]byte
	TokenAddr   common.Address
	Recipient   common.Address
	FromChainId *big.Int
	TokenId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBackwardSwapFilled is a free log retrieval operation binding the contract event 0x3465ebb4fdfd4cdb5a9c6020ffa12bd1a621dca4e158f14497eff75a33fda45f.
//
// Solidity: event BackwardSwapFilled(bytes32 indexed swapTxHash, address indexed tokenAddr, address indexed recipient, uint256 fromChainId, uint256 tokenId)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) FilterBackwardSwapFilled(opts *bind.FilterOpts, swapTxHash [][32]byte, tokenAddr []common.Address, recipient []common.Address) (*ERC721SwapAgentBackwardSwapFilledIterator, error) {

	var swapTxHashRule []interface{}
	for _, swapTxHashItem := range swapTxHash {
		swapTxHashRule = append(swapTxHashRule, swapTxHashItem)
	}
	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC721SwapAgent.contract.FilterLogs(opts, "BackwardSwapFilled", swapTxHashRule, tokenAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC721SwapAgentBackwardSwapFilledIterator{contract: _ERC721SwapAgent.contract, event: "BackwardSwapFilled", logs: logs, sub: sub}, nil
}

// WatchBackwardSwapFilled is a free log subscription operation binding the contract event 0x3465ebb4fdfd4cdb5a9c6020ffa12bd1a621dca4e158f14497eff75a33fda45f.
//
// Solidity: event BackwardSwapFilled(bytes32 indexed swapTxHash, address indexed tokenAddr, address indexed recipient, uint256 fromChainId, uint256 tokenId)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) WatchBackwardSwapFilled(opts *bind.WatchOpts, sink chan<- *ERC721SwapAgentBackwardSwapFilled, swapTxHash [][32]byte, tokenAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var swapTxHashRule []interface{}
	for _, swapTxHashItem := range swapTxHash {
		swapTxHashRule = append(swapTxHashRule, swapTxHashItem)
	}
	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC721SwapAgent.contract.WatchLogs(opts, "BackwardSwapFilled", swapTxHashRule, tokenAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721SwapAgentBackwardSwapFilled)
				if err := _ERC721SwapAgent.contract.UnpackLog(event, "BackwardSwapFilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBackwardSwapFilled is a log parse operation binding the contract event 0x3465ebb4fdfd4cdb5a9c6020ffa12bd1a621dca4e158f14497eff75a33fda45f.
//
// Solidity: event BackwardSwapFilled(bytes32 indexed swapTxHash, address indexed tokenAddr, address indexed recipient, uint256 fromChainId, uint256 tokenId)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) ParseBackwardSwapFilled(log types.Log) (*ERC721SwapAgentBackwardSwapFilled, error) {
	event := new(ERC721SwapAgentBackwardSwapFilled)
	if err := _ERC721SwapAgent.contract.UnpackLog(event, "BackwardSwapFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721SwapAgentBackwardSwapStartedIterator is returned from FilterBackwardSwapStarted and is used to iterate over the raw logs and unpacked data for BackwardSwapStarted events raised by the ERC721SwapAgent contract.
type ERC721SwapAgentBackwardSwapStartedIterator struct {
	Event *ERC721SwapAgentBackwardSwapStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC721SwapAgentBackwardSwapStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721SwapAgentBackwardSwapStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC721SwapAgentBackwardSwapStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC721SwapAgentBackwardSwapStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721SwapAgentBackwardSwapStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721SwapAgentBackwardSwapStarted represents a BackwardSwapStarted event raised by the ERC721SwapAgent contract.
type ERC721SwapAgentBackwardSwapStarted struct {
	MirroredTokenAddr common.Address
	Sender            common.Address
	Recipient         common.Address
	DstChainId        *big.Int
	TokenId           *big.Int
	FeeAmount         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBackwardSwapStarted is a free log retrieval operation binding the contract event 0x3fb82ea212f026f013c4a9982b7422161274b487e5f7f3499ae555390a054662.
//
// Solidity: event BackwardSwapStarted(address indexed mirroredTokenAddr, address indexed sender, address indexed recipient, uint256 dstChainId, uint256 tokenId, uint256 feeAmount)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) FilterBackwardSwapStarted(opts *bind.FilterOpts, mirroredTokenAddr []common.Address, sender []common.Address, recipient []common.Address) (*ERC721SwapAgentBackwardSwapStartedIterator, error) {

	var mirroredTokenAddrRule []interface{}
	for _, mirroredTokenAddrItem := range mirroredTokenAddr {
		mirroredTokenAddrRule = append(mirroredTokenAddrRule, mirroredTokenAddrItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC721SwapAgent.contract.FilterLogs(opts, "BackwardSwapStarted", mirroredTokenAddrRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC721SwapAgentBackwardSwapStartedIterator{contract: _ERC721SwapAgent.contract, event: "BackwardSwapStarted", logs: logs, sub: sub}, nil
}

// WatchBackwardSwapStarted is a free log subscription operation binding the contract event 0x3fb82ea212f026f013c4a9982b7422161274b487e5f7f3499ae555390a054662.
//
// Solidity: event BackwardSwapStarted(address indexed mirroredTokenAddr, address indexed sender, address indexed recipient, uint256 dstChainId, uint256 tokenId, uint256 feeAmount)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) WatchBackwardSwapStarted(opts *bind.WatchOpts, sink chan<- *ERC721SwapAgentBackwardSwapStarted, mirroredTokenAddr []common.Address, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var mirroredTokenAddrRule []interface{}
	for _, mirroredTokenAddrItem := range mirroredTokenAddr {
		mirroredTokenAddrRule = append(mirroredTokenAddrRule, mirroredTokenAddrItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC721SwapAgent.contract.WatchLogs(opts, "BackwardSwapStarted", mirroredTokenAddrRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721SwapAgentBackwardSwapStarted)
				if err := _ERC721SwapAgent.contract.UnpackLog(event, "BackwardSwapStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBackwardSwapStarted is a log parse operation binding the contract event 0x3fb82ea212f026f013c4a9982b7422161274b487e5f7f3499ae555390a054662.
//
// Solidity: event BackwardSwapStarted(address indexed mirroredTokenAddr, address indexed sender, address indexed recipient, uint256 dstChainId, uint256 tokenId, uint256 feeAmount)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) ParseBackwardSwapStarted(log types.Log) (*ERC721SwapAgentBackwardSwapStarted, error) {
	event := new(ERC721SwapAgentBackwardSwapStarted)
	if err := _ERC721SwapAgent.contract.UnpackLog(event, "BackwardSwapStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721SwapAgentOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC721SwapAgent contract.
type ERC721SwapAgentOwnershipTransferredIterator struct {
	Event *ERC721SwapAgentOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC721SwapAgentOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721SwapAgentOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC721SwapAgentOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC721SwapAgentOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721SwapAgentOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721SwapAgentOwnershipTransferred represents a OwnershipTransferred event raised by the ERC721SwapAgent contract.
type ERC721SwapAgentOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC721SwapAgentOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC721SwapAgent.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC721SwapAgentOwnershipTransferredIterator{contract: _ERC721SwapAgent.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC721SwapAgentOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC721SwapAgent.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721SwapAgentOwnershipTransferred)
				if err := _ERC721SwapAgent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) ParseOwnershipTransferred(log types.Log) (*ERC721SwapAgentOwnershipTransferred, error) {
	event := new(ERC721SwapAgentOwnershipTransferred)
	if err := _ERC721SwapAgent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721SwapAgentSwapFilledIterator is returned from FilterSwapFilled and is used to iterate over the raw logs and unpacked data for SwapFilled events raised by the ERC721SwapAgent contract.
type ERC721SwapAgentSwapFilledIterator struct {
	Event *ERC721SwapAgentSwapFilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC721SwapAgentSwapFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721SwapAgentSwapFilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC721SwapAgentSwapFilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC721SwapAgentSwapFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721SwapAgentSwapFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721SwapAgentSwapFilled represents a SwapFilled event raised by the ERC721SwapAgent contract.
type ERC721SwapAgentSwapFilled struct {
	SwapTxHash        [32]byte
	FromTokenAddr     common.Address
	Recipient         common.Address
	MirroredTokenAddr common.Address
	FromChainId       *big.Int
	TokenId           *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapFilled is a free log retrieval operation binding the contract event 0xf1af0abbd42bfb09f51c7192f406e6cdc03db09a8dcdcdfcd99e954477c60601.
//
// Solidity: event SwapFilled(bytes32 indexed swapTxHash, address indexed fromTokenAddr, address indexed recipient, address mirroredTokenAddr, uint256 fromChainId, uint256 tokenId)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) FilterSwapFilled(opts *bind.FilterOpts, swapTxHash [][32]byte, fromTokenAddr []common.Address, recipient []common.Address) (*ERC721SwapAgentSwapFilledIterator, error) {

	var swapTxHashRule []interface{}
	for _, swapTxHashItem := range swapTxHash {
		swapTxHashRule = append(swapTxHashRule, swapTxHashItem)
	}
	var fromTokenAddrRule []interface{}
	for _, fromTokenAddrItem := range fromTokenAddr {
		fromTokenAddrRule = append(fromTokenAddrRule, fromTokenAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC721SwapAgent.contract.FilterLogs(opts, "SwapFilled", swapTxHashRule, fromTokenAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC721SwapAgentSwapFilledIterator{contract: _ERC721SwapAgent.contract, event: "SwapFilled", logs: logs, sub: sub}, nil
}

// WatchSwapFilled is a free log subscription operation binding the contract event 0xf1af0abbd42bfb09f51c7192f406e6cdc03db09a8dcdcdfcd99e954477c60601.
//
// Solidity: event SwapFilled(bytes32 indexed swapTxHash, address indexed fromTokenAddr, address indexed recipient, address mirroredTokenAddr, uint256 fromChainId, uint256 tokenId)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) WatchSwapFilled(opts *bind.WatchOpts, sink chan<- *ERC721SwapAgentSwapFilled, swapTxHash [][32]byte, fromTokenAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var swapTxHashRule []interface{}
	for _, swapTxHashItem := range swapTxHash {
		swapTxHashRule = append(swapTxHashRule, swapTxHashItem)
	}
	var fromTokenAddrRule []interface{}
	for _, fromTokenAddrItem := range fromTokenAddr {
		fromTokenAddrRule = append(fromTokenAddrRule, fromTokenAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC721SwapAgent.contract.WatchLogs(opts, "SwapFilled", swapTxHashRule, fromTokenAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721SwapAgentSwapFilled)
				if err := _ERC721SwapAgent.contract.UnpackLog(event, "SwapFilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwapFilled is a log parse operation binding the contract event 0xf1af0abbd42bfb09f51c7192f406e6cdc03db09a8dcdcdfcd99e954477c60601.
//
// Solidity: event SwapFilled(bytes32 indexed swapTxHash, address indexed fromTokenAddr, address indexed recipient, address mirroredTokenAddr, uint256 fromChainId, uint256 tokenId)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) ParseSwapFilled(log types.Log) (*ERC721SwapAgentSwapFilled, error) {
	event := new(ERC721SwapAgentSwapFilled)
	if err := _ERC721SwapAgent.contract.UnpackLog(event, "SwapFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721SwapAgentSwapPairCreatedIterator is returned from FilterSwapPairCreated and is used to iterate over the raw logs and unpacked data for SwapPairCreated events raised by the ERC721SwapAgent contract.
type ERC721SwapAgentSwapPairCreatedIterator struct {
	Event *ERC721SwapAgentSwapPairCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC721SwapAgentSwapPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721SwapAgentSwapPairCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC721SwapAgentSwapPairCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC721SwapAgentSwapPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721SwapAgentSwapPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721SwapAgentSwapPairCreated represents a SwapPairCreated event raised by the ERC721SwapAgent contract.
type ERC721SwapAgentSwapPairCreated struct {
	RegisterTxHash    [32]byte
	FromTokenAddr     common.Address
	MirroredTokenAddr common.Address
	FromChainId       *big.Int
	TokenSymbol       string
	TokenName         string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapPairCreated is a free log retrieval operation binding the contract event 0xf7346649f06f58e0489664a33d3cddd434424d3b49ca70c1a8808f488e4ecd49.
//
// Solidity: event SwapPairCreated(bytes32 indexed registerTxHash, address indexed fromTokenAddr, address indexed mirroredTokenAddr, uint256 fromChainId, string tokenSymbol, string tokenName)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) FilterSwapPairCreated(opts *bind.FilterOpts, registerTxHash [][32]byte, fromTokenAddr []common.Address, mirroredTokenAddr []common.Address) (*ERC721SwapAgentSwapPairCreatedIterator, error) {

	var registerTxHashRule []interface{}
	for _, registerTxHashItem := range registerTxHash {
		registerTxHashRule = append(registerTxHashRule, registerTxHashItem)
	}
	var fromTokenAddrRule []interface{}
	for _, fromTokenAddrItem := range fromTokenAddr {
		fromTokenAddrRule = append(fromTokenAddrRule, fromTokenAddrItem)
	}
	var mirroredTokenAddrRule []interface{}
	for _, mirroredTokenAddrItem := range mirroredTokenAddr {
		mirroredTokenAddrRule = append(mirroredTokenAddrRule, mirroredTokenAddrItem)
	}

	logs, sub, err := _ERC721SwapAgent.contract.FilterLogs(opts, "SwapPairCreated", registerTxHashRule, fromTokenAddrRule, mirroredTokenAddrRule)
	if err != nil {
		return nil, err
	}
	return &ERC721SwapAgentSwapPairCreatedIterator{contract: _ERC721SwapAgent.contract, event: "SwapPairCreated", logs: logs, sub: sub}, nil
}

// WatchSwapPairCreated is a free log subscription operation binding the contract event 0xf7346649f06f58e0489664a33d3cddd434424d3b49ca70c1a8808f488e4ecd49.
//
// Solidity: event SwapPairCreated(bytes32 indexed registerTxHash, address indexed fromTokenAddr, address indexed mirroredTokenAddr, uint256 fromChainId, string tokenSymbol, string tokenName)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) WatchSwapPairCreated(opts *bind.WatchOpts, sink chan<- *ERC721SwapAgentSwapPairCreated, registerTxHash [][32]byte, fromTokenAddr []common.Address, mirroredTokenAddr []common.Address) (event.Subscription, error) {

	var registerTxHashRule []interface{}
	for _, registerTxHashItem := range registerTxHash {
		registerTxHashRule = append(registerTxHashRule, registerTxHashItem)
	}
	var fromTokenAddrRule []interface{}
	for _, fromTokenAddrItem := range fromTokenAddr {
		fromTokenAddrRule = append(fromTokenAddrRule, fromTokenAddrItem)
	}
	var mirroredTokenAddrRule []interface{}
	for _, mirroredTokenAddrItem := range mirroredTokenAddr {
		mirroredTokenAddrRule = append(mirroredTokenAddrRule, mirroredTokenAddrItem)
	}

	logs, sub, err := _ERC721SwapAgent.contract.WatchLogs(opts, "SwapPairCreated", registerTxHashRule, fromTokenAddrRule, mirroredTokenAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721SwapAgentSwapPairCreated)
				if err := _ERC721SwapAgent.contract.UnpackLog(event, "SwapPairCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwapPairCreated is a log parse operation binding the contract event 0xf7346649f06f58e0489664a33d3cddd434424d3b49ca70c1a8808f488e4ecd49.
//
// Solidity: event SwapPairCreated(bytes32 indexed registerTxHash, address indexed fromTokenAddr, address indexed mirroredTokenAddr, uint256 fromChainId, string tokenSymbol, string tokenName)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) ParseSwapPairCreated(log types.Log) (*ERC721SwapAgentSwapPairCreated, error) {
	event := new(ERC721SwapAgentSwapPairCreated)
	if err := _ERC721SwapAgent.contract.UnpackLog(event, "SwapPairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721SwapAgentSwapPairRegisterIterator is returned from FilterSwapPairRegister and is used to iterate over the raw logs and unpacked data for SwapPairRegister events raised by the ERC721SwapAgent contract.
type ERC721SwapAgentSwapPairRegisterIterator struct {
	Event *ERC721SwapAgentSwapPairRegister // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC721SwapAgentSwapPairRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721SwapAgentSwapPairRegister)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC721SwapAgentSwapPairRegister)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC721SwapAgentSwapPairRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721SwapAgentSwapPairRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721SwapAgentSwapPairRegister represents a SwapPairRegister event raised by the ERC721SwapAgent contract.
type ERC721SwapAgentSwapPairRegister struct {
	Sponsor      common.Address
	TokenAddress common.Address
	TokenName    string
	TokenSymbol  string
	ToChainId    *big.Int
	FeeAmount    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwapPairRegister is a free log retrieval operation binding the contract event 0x254796a39d303c3ef102d83626b1cca9284dcb7e0bc4d33ca798f536017868dc.
//
// Solidity: event SwapPairRegister(address indexed sponsor, address indexed tokenAddress, string tokenName, string tokenSymbol, uint256 toChainId, uint256 feeAmount)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) FilterSwapPairRegister(opts *bind.FilterOpts, sponsor []common.Address, tokenAddress []common.Address) (*ERC721SwapAgentSwapPairRegisterIterator, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _ERC721SwapAgent.contract.FilterLogs(opts, "SwapPairRegister", sponsorRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC721SwapAgentSwapPairRegisterIterator{contract: _ERC721SwapAgent.contract, event: "SwapPairRegister", logs: logs, sub: sub}, nil
}

// WatchSwapPairRegister is a free log subscription operation binding the contract event 0x254796a39d303c3ef102d83626b1cca9284dcb7e0bc4d33ca798f536017868dc.
//
// Solidity: event SwapPairRegister(address indexed sponsor, address indexed tokenAddress, string tokenName, string tokenSymbol, uint256 toChainId, uint256 feeAmount)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) WatchSwapPairRegister(opts *bind.WatchOpts, sink chan<- *ERC721SwapAgentSwapPairRegister, sponsor []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _ERC721SwapAgent.contract.WatchLogs(opts, "SwapPairRegister", sponsorRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721SwapAgentSwapPairRegister)
				if err := _ERC721SwapAgent.contract.UnpackLog(event, "SwapPairRegister", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwapPairRegister is a log parse operation binding the contract event 0x254796a39d303c3ef102d83626b1cca9284dcb7e0bc4d33ca798f536017868dc.
//
// Solidity: event SwapPairRegister(address indexed sponsor, address indexed tokenAddress, string tokenName, string tokenSymbol, uint256 toChainId, uint256 feeAmount)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) ParseSwapPairRegister(log types.Log) (*ERC721SwapAgentSwapPairRegister, error) {
	event := new(ERC721SwapAgentSwapPairRegister)
	if err := _ERC721SwapAgent.contract.UnpackLog(event, "SwapPairRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721SwapAgentSwapStartedIterator is returned from FilterSwapStarted and is used to iterate over the raw logs and unpacked data for SwapStarted events raised by the ERC721SwapAgent contract.
type ERC721SwapAgentSwapStartedIterator struct {
	Event *ERC721SwapAgentSwapStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC721SwapAgentSwapStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721SwapAgentSwapStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC721SwapAgentSwapStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC721SwapAgentSwapStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721SwapAgentSwapStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721SwapAgentSwapStarted represents a SwapStarted event raised by the ERC721SwapAgent contract.
type ERC721SwapAgentSwapStarted struct {
	TokenAddr  common.Address
	Sender     common.Address
	Recipient  common.Address
	DstChainId *big.Int
	TokenId    *big.Int
	FeeAmount  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwapStarted is a free log retrieval operation binding the contract event 0x18e4fc12755e4744b0dd88b81c6c56d69b85b73c98d90ff1afbc8cc9166834f8.
//
// Solidity: event SwapStarted(address indexed tokenAddr, address indexed sender, address indexed recipient, uint256 dstChainId, uint256 tokenId, uint256 feeAmount)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) FilterSwapStarted(opts *bind.FilterOpts, tokenAddr []common.Address, sender []common.Address, recipient []common.Address) (*ERC721SwapAgentSwapStartedIterator, error) {

	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC721SwapAgent.contract.FilterLogs(opts, "SwapStarted", tokenAddrRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC721SwapAgentSwapStartedIterator{contract: _ERC721SwapAgent.contract, event: "SwapStarted", logs: logs, sub: sub}, nil
}

// WatchSwapStarted is a free log subscription operation binding the contract event 0x18e4fc12755e4744b0dd88b81c6c56d69b85b73c98d90ff1afbc8cc9166834f8.
//
// Solidity: event SwapStarted(address indexed tokenAddr, address indexed sender, address indexed recipient, uint256 dstChainId, uint256 tokenId, uint256 feeAmount)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) WatchSwapStarted(opts *bind.WatchOpts, sink chan<- *ERC721SwapAgentSwapStarted, tokenAddr []common.Address, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC721SwapAgent.contract.WatchLogs(opts, "SwapStarted", tokenAddrRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721SwapAgentSwapStarted)
				if err := _ERC721SwapAgent.contract.UnpackLog(event, "SwapStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwapStarted is a log parse operation binding the contract event 0x18e4fc12755e4744b0dd88b81c6c56d69b85b73c98d90ff1afbc8cc9166834f8.
//
// Solidity: event SwapStarted(address indexed tokenAddr, address indexed sender, address indexed recipient, uint256 dstChainId, uint256 tokenId, uint256 feeAmount)
func (_ERC721SwapAgent *ERC721SwapAgentFilterer) ParseSwapStarted(log types.Log) (*ERC721SwapAgentSwapStarted, error) {
	event := new(ERC721SwapAgentSwapStarted)
	if err := _ERC721SwapAgent.contract.UnpackLog(event, "SwapStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
