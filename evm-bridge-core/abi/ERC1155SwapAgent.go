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

// ERC1155SwapAgentMetaData contains all meta data concerning the ERC1155SwapAgent contract.
var ERC1155SwapAgentMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"swapTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"BackwardSwapFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mirroredTokenAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"BackwardSwapStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"swapTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromTokenAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mirroredTokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"SwapFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"registerTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromTokenAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mirroredTokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"}],\"name\":\"SwapPairCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"SwapPairRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"SwapStarted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"registerTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"fromTokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"createSwapPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"swapTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"fromTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"fill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filledSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"registerSwapPair\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registeredToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapMappingIncoming\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapMappingOutgoing\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC1155SwapAgentABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1155SwapAgentMetaData.ABI instead.
var ERC1155SwapAgentABI = ERC1155SwapAgentMetaData.ABI

// ERC1155SwapAgent is an auto generated Go binding around an Ethereum contract.
type ERC1155SwapAgent struct {
	ERC1155SwapAgentCaller     // Read-only binding to the contract
	ERC1155SwapAgentTransactor // Write-only binding to the contract
	ERC1155SwapAgentFilterer   // Log filterer for contract events
}

// ERC1155SwapAgentCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1155SwapAgentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155SwapAgentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1155SwapAgentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155SwapAgentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1155SwapAgentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155SwapAgentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1155SwapAgentSession struct {
	Contract     *ERC1155SwapAgent // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1155SwapAgentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1155SwapAgentCallerSession struct {
	Contract *ERC1155SwapAgentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC1155SwapAgentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1155SwapAgentTransactorSession struct {
	Contract     *ERC1155SwapAgentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC1155SwapAgentRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1155SwapAgentRaw struct {
	Contract *ERC1155SwapAgent // Generic contract binding to access the raw methods on
}

// ERC1155SwapAgentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1155SwapAgentCallerRaw struct {
	Contract *ERC1155SwapAgentCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1155SwapAgentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1155SwapAgentTransactorRaw struct {
	Contract *ERC1155SwapAgentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1155SwapAgent creates a new instance of ERC1155SwapAgent, bound to a specific deployed contract.
func NewERC1155SwapAgent(address common.Address, backend bind.ContractBackend) (*ERC1155SwapAgent, error) {
	contract, err := bindERC1155SwapAgent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1155SwapAgent{ERC1155SwapAgentCaller: ERC1155SwapAgentCaller{contract: contract}, ERC1155SwapAgentTransactor: ERC1155SwapAgentTransactor{contract: contract}, ERC1155SwapAgentFilterer: ERC1155SwapAgentFilterer{contract: contract}}, nil
}

// NewERC1155SwapAgentCaller creates a new read-only instance of ERC1155SwapAgent, bound to a specific deployed contract.
func NewERC1155SwapAgentCaller(address common.Address, caller bind.ContractCaller) (*ERC1155SwapAgentCaller, error) {
	contract, err := bindERC1155SwapAgent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155SwapAgentCaller{contract: contract}, nil
}

// NewERC1155SwapAgentTransactor creates a new write-only instance of ERC1155SwapAgent, bound to a specific deployed contract.
func NewERC1155SwapAgentTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1155SwapAgentTransactor, error) {
	contract, err := bindERC1155SwapAgent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155SwapAgentTransactor{contract: contract}, nil
}

// NewERC1155SwapAgentFilterer creates a new log filterer instance of ERC1155SwapAgent, bound to a specific deployed contract.
func NewERC1155SwapAgentFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1155SwapAgentFilterer, error) {
	contract, err := bindERC1155SwapAgent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1155SwapAgentFilterer{contract: contract}, nil
}

// bindERC1155SwapAgent binds a generic wrapper to an already deployed contract.
func bindERC1155SwapAgent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC1155SwapAgentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155SwapAgent *ERC1155SwapAgentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155SwapAgent.Contract.ERC1155SwapAgentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155SwapAgent *ERC1155SwapAgentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.ERC1155SwapAgentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155SwapAgent *ERC1155SwapAgentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.ERC1155SwapAgentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155SwapAgent *ERC1155SwapAgentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155SwapAgent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.contract.Transact(opts, method, params...)
}

// FilledSwap is a free data retrieval call binding the contract method 0xa86894ca.
//
// Solidity: function filledSwap(bytes32 ) view returns(bool)
func (_ERC1155SwapAgent *ERC1155SwapAgentCaller) FilledSwap(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ERC1155SwapAgent.contract.Call(opts, &out, "filledSwap", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FilledSwap is a free data retrieval call binding the contract method 0xa86894ca.
//
// Solidity: function filledSwap(bytes32 ) view returns(bool)
func (_ERC1155SwapAgent *ERC1155SwapAgentSession) FilledSwap(arg0 [32]byte) (bool, error) {
	return _ERC1155SwapAgent.Contract.FilledSwap(&_ERC1155SwapAgent.CallOpts, arg0)
}

// FilledSwap is a free data retrieval call binding the contract method 0xa86894ca.
//
// Solidity: function filledSwap(bytes32 ) view returns(bool)
func (_ERC1155SwapAgent *ERC1155SwapAgentCallerSession) FilledSwap(arg0 [32]byte) (bool, error) {
	return _ERC1155SwapAgent.Contract.FilledSwap(&_ERC1155SwapAgent.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC1155SwapAgent *ERC1155SwapAgentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC1155SwapAgent.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC1155SwapAgent *ERC1155SwapAgentSession) Owner() (common.Address, error) {
	return _ERC1155SwapAgent.Contract.Owner(&_ERC1155SwapAgent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC1155SwapAgent *ERC1155SwapAgentCallerSession) Owner() (common.Address, error) {
	return _ERC1155SwapAgent.Contract.Owner(&_ERC1155SwapAgent.CallOpts)
}

// RegisteredToken is a free data retrieval call binding the contract method 0x0b4f43c1.
//
// Solidity: function registeredToken(uint256 , address ) view returns(bool)
func (_ERC1155SwapAgent *ERC1155SwapAgentCaller) RegisteredToken(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _ERC1155SwapAgent.contract.Call(opts, &out, "registeredToken", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredToken is a free data retrieval call binding the contract method 0x0b4f43c1.
//
// Solidity: function registeredToken(uint256 , address ) view returns(bool)
func (_ERC1155SwapAgent *ERC1155SwapAgentSession) RegisteredToken(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _ERC1155SwapAgent.Contract.RegisteredToken(&_ERC1155SwapAgent.CallOpts, arg0, arg1)
}

// RegisteredToken is a free data retrieval call binding the contract method 0x0b4f43c1.
//
// Solidity: function registeredToken(uint256 , address ) view returns(bool)
func (_ERC1155SwapAgent *ERC1155SwapAgentCallerSession) RegisteredToken(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _ERC1155SwapAgent.Contract.RegisteredToken(&_ERC1155SwapAgent.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155SwapAgent *ERC1155SwapAgentCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC1155SwapAgent.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155SwapAgent *ERC1155SwapAgentSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155SwapAgent.Contract.SupportsInterface(&_ERC1155SwapAgent.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC1155SwapAgent *ERC1155SwapAgentCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC1155SwapAgent.Contract.SupportsInterface(&_ERC1155SwapAgent.CallOpts, interfaceId)
}

// SwapMappingIncoming is a free data retrieval call binding the contract method 0xec686704.
//
// Solidity: function swapMappingIncoming(uint256 , address ) view returns(address)
func (_ERC1155SwapAgent *ERC1155SwapAgentCaller) SwapMappingIncoming(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ERC1155SwapAgent.contract.Call(opts, &out, "swapMappingIncoming", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapMappingIncoming is a free data retrieval call binding the contract method 0xec686704.
//
// Solidity: function swapMappingIncoming(uint256 , address ) view returns(address)
func (_ERC1155SwapAgent *ERC1155SwapAgentSession) SwapMappingIncoming(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _ERC1155SwapAgent.Contract.SwapMappingIncoming(&_ERC1155SwapAgent.CallOpts, arg0, arg1)
}

// SwapMappingIncoming is a free data retrieval call binding the contract method 0xec686704.
//
// Solidity: function swapMappingIncoming(uint256 , address ) view returns(address)
func (_ERC1155SwapAgent *ERC1155SwapAgentCallerSession) SwapMappingIncoming(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _ERC1155SwapAgent.Contract.SwapMappingIncoming(&_ERC1155SwapAgent.CallOpts, arg0, arg1)
}

// SwapMappingOutgoing is a free data retrieval call binding the contract method 0x0d43d992.
//
// Solidity: function swapMappingOutgoing(uint256 , address ) view returns(address)
func (_ERC1155SwapAgent *ERC1155SwapAgentCaller) SwapMappingOutgoing(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ERC1155SwapAgent.contract.Call(opts, &out, "swapMappingOutgoing", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapMappingOutgoing is a free data retrieval call binding the contract method 0x0d43d992.
//
// Solidity: function swapMappingOutgoing(uint256 , address ) view returns(address)
func (_ERC1155SwapAgent *ERC1155SwapAgentSession) SwapMappingOutgoing(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _ERC1155SwapAgent.Contract.SwapMappingOutgoing(&_ERC1155SwapAgent.CallOpts, arg0, arg1)
}

// SwapMappingOutgoing is a free data retrieval call binding the contract method 0x0d43d992.
//
// Solidity: function swapMappingOutgoing(uint256 , address ) view returns(address)
func (_ERC1155SwapAgent *ERC1155SwapAgentCallerSession) SwapMappingOutgoing(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _ERC1155SwapAgent.Contract.SwapMappingOutgoing(&_ERC1155SwapAgent.CallOpts, arg0, arg1)
}

// CreateSwapPair is a paid mutator transaction binding the contract method 0x4acbe1ca.
//
// Solidity: function createSwapPair(bytes32 registerTxHash, address fromTokenAddr, uint256 fromChainId, string uri) returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactor) CreateSwapPair(opts *bind.TransactOpts, registerTxHash [32]byte, fromTokenAddr common.Address, fromChainId *big.Int, uri string) (*types.Transaction, error) {
	return _ERC1155SwapAgent.contract.Transact(opts, "createSwapPair", registerTxHash, fromTokenAddr, fromChainId, uri)
}

// CreateSwapPair is a paid mutator transaction binding the contract method 0x4acbe1ca.
//
// Solidity: function createSwapPair(bytes32 registerTxHash, address fromTokenAddr, uint256 fromChainId, string uri) returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentSession) CreateSwapPair(registerTxHash [32]byte, fromTokenAddr common.Address, fromChainId *big.Int, uri string) (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.CreateSwapPair(&_ERC1155SwapAgent.TransactOpts, registerTxHash, fromTokenAddr, fromChainId, uri)
}

// CreateSwapPair is a paid mutator transaction binding the contract method 0x4acbe1ca.
//
// Solidity: function createSwapPair(bytes32 registerTxHash, address fromTokenAddr, uint256 fromChainId, string uri) returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactorSession) CreateSwapPair(registerTxHash [32]byte, fromTokenAddr common.Address, fromChainId *big.Int, uri string) (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.CreateSwapPair(&_ERC1155SwapAgent.TransactOpts, registerTxHash, fromTokenAddr, fromChainId, uri)
}

// Fill is a paid mutator transaction binding the contract method 0x04828122.
//
// Solidity: function fill(bytes32 swapTxHash, address fromTokenAddr, address recipient, uint256 fromChainId, uint256[] ids, uint256[] amounts) returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactor) Fill(opts *bind.TransactOpts, swapTxHash [32]byte, fromTokenAddr common.Address, recipient common.Address, fromChainId *big.Int, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _ERC1155SwapAgent.contract.Transact(opts, "fill", swapTxHash, fromTokenAddr, recipient, fromChainId, ids, amounts)
}

// Fill is a paid mutator transaction binding the contract method 0x04828122.
//
// Solidity: function fill(bytes32 swapTxHash, address fromTokenAddr, address recipient, uint256 fromChainId, uint256[] ids, uint256[] amounts) returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentSession) Fill(swapTxHash [32]byte, fromTokenAddr common.Address, recipient common.Address, fromChainId *big.Int, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.Fill(&_ERC1155SwapAgent.TransactOpts, swapTxHash, fromTokenAddr, recipient, fromChainId, ids, amounts)
}

// Fill is a paid mutator transaction binding the contract method 0x04828122.
//
// Solidity: function fill(bytes32 swapTxHash, address fromTokenAddr, address recipient, uint256 fromChainId, uint256[] ids, uint256[] amounts) returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactorSession) Fill(swapTxHash [32]byte, fromTokenAddr common.Address, recipient common.Address, fromChainId *big.Int, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.Fill(&_ERC1155SwapAgent.TransactOpts, swapTxHash, fromTokenAddr, recipient, fromChainId, ids, amounts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155SwapAgent.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentSession) Initialize() (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.Initialize(&_ERC1155SwapAgent.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactorSession) Initialize() (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.Initialize(&_ERC1155SwapAgent.TransactOpts)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155SwapAgent.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ERC1155SwapAgent *ERC1155SwapAgentSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.OnERC1155BatchReceived(&_ERC1155SwapAgent.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.OnERC1155BatchReceived(&_ERC1155SwapAgent.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155SwapAgent.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ERC1155SwapAgent *ERC1155SwapAgentSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.OnERC1155Received(&_ERC1155SwapAgent.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.OnERC1155Received(&_ERC1155SwapAgent.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// RegisterSwapPair is a paid mutator transaction binding the contract method 0x45b1ab1b.
//
// Solidity: function registerSwapPair(address tokenAddr, uint256 chainId) payable returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactor) RegisterSwapPair(opts *bind.TransactOpts, tokenAddr common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _ERC1155SwapAgent.contract.Transact(opts, "registerSwapPair", tokenAddr, chainId)
}

// RegisterSwapPair is a paid mutator transaction binding the contract method 0x45b1ab1b.
//
// Solidity: function registerSwapPair(address tokenAddr, uint256 chainId) payable returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentSession) RegisterSwapPair(tokenAddr common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.RegisterSwapPair(&_ERC1155SwapAgent.TransactOpts, tokenAddr, chainId)
}

// RegisterSwapPair is a paid mutator transaction binding the contract method 0x45b1ab1b.
//
// Solidity: function registerSwapPair(address tokenAddr, uint256 chainId) payable returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactorSession) RegisterSwapPair(tokenAddr common.Address, chainId *big.Int) (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.RegisterSwapPair(&_ERC1155SwapAgent.TransactOpts, tokenAddr, chainId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155SwapAgent.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.RenounceOwnership(&_ERC1155SwapAgent.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.RenounceOwnership(&_ERC1155SwapAgent.TransactOpts)
}

// Swap is a paid mutator transaction binding the contract method 0xa180639a.
//
// Solidity: function swap(address tokenAddr, address recipient, uint256[] ids, uint256[] amounts, uint256 dstChainId) payable returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactor) Swap(opts *bind.TransactOpts, tokenAddr common.Address, recipient common.Address, ids []*big.Int, amounts []*big.Int, dstChainId *big.Int) (*types.Transaction, error) {
	return _ERC1155SwapAgent.contract.Transact(opts, "swap", tokenAddr, recipient, ids, amounts, dstChainId)
}

// Swap is a paid mutator transaction binding the contract method 0xa180639a.
//
// Solidity: function swap(address tokenAddr, address recipient, uint256[] ids, uint256[] amounts, uint256 dstChainId) payable returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentSession) Swap(tokenAddr common.Address, recipient common.Address, ids []*big.Int, amounts []*big.Int, dstChainId *big.Int) (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.Swap(&_ERC1155SwapAgent.TransactOpts, tokenAddr, recipient, ids, amounts, dstChainId)
}

// Swap is a paid mutator transaction binding the contract method 0xa180639a.
//
// Solidity: function swap(address tokenAddr, address recipient, uint256[] ids, uint256[] amounts, uint256 dstChainId) payable returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactorSession) Swap(tokenAddr common.Address, recipient common.Address, ids []*big.Int, amounts []*big.Int, dstChainId *big.Int) (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.Swap(&_ERC1155SwapAgent.TransactOpts, tokenAddr, recipient, ids, amounts, dstChainId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC1155SwapAgent.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.TransferOwnership(&_ERC1155SwapAgent.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC1155SwapAgent *ERC1155SwapAgentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC1155SwapAgent.Contract.TransferOwnership(&_ERC1155SwapAgent.TransactOpts, newOwner)
}

// ERC1155SwapAgentBackwardSwapFilledIterator is returned from FilterBackwardSwapFilled and is used to iterate over the raw logs and unpacked data for BackwardSwapFilled events raised by the ERC1155SwapAgent contract.
type ERC1155SwapAgentBackwardSwapFilledIterator struct {
	Event *ERC1155SwapAgentBackwardSwapFilled // Event containing the contract specifics and raw log

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
func (it *ERC1155SwapAgentBackwardSwapFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155SwapAgentBackwardSwapFilled)
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
		it.Event = new(ERC1155SwapAgentBackwardSwapFilled)
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
func (it *ERC1155SwapAgentBackwardSwapFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155SwapAgentBackwardSwapFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155SwapAgentBackwardSwapFilled represents a BackwardSwapFilled event raised by the ERC1155SwapAgent contract.
type ERC1155SwapAgentBackwardSwapFilled struct {
	SwapTxHash  [32]byte
	TokenAddr   common.Address
	Recipient   common.Address
	FromChainId *big.Int
	Ids         []*big.Int
	Amounts     []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBackwardSwapFilled is a free log retrieval operation binding the contract event 0x58d9c075708eb3187538135716481600d99ff6eb833bdbac1fee1e54cebec1b3.
//
// Solidity: event BackwardSwapFilled(bytes32 indexed swapTxHash, address indexed tokenAddr, address indexed recipient, uint256 fromChainId, uint256[] ids, uint256[] amounts)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) FilterBackwardSwapFilled(opts *bind.FilterOpts, swapTxHash [][32]byte, tokenAddr []common.Address, recipient []common.Address) (*ERC1155SwapAgentBackwardSwapFilledIterator, error) {

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

	logs, sub, err := _ERC1155SwapAgent.contract.FilterLogs(opts, "BackwardSwapFilled", swapTxHashRule, tokenAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155SwapAgentBackwardSwapFilledIterator{contract: _ERC1155SwapAgent.contract, event: "BackwardSwapFilled", logs: logs, sub: sub}, nil
}

// WatchBackwardSwapFilled is a free log subscription operation binding the contract event 0x58d9c075708eb3187538135716481600d99ff6eb833bdbac1fee1e54cebec1b3.
//
// Solidity: event BackwardSwapFilled(bytes32 indexed swapTxHash, address indexed tokenAddr, address indexed recipient, uint256 fromChainId, uint256[] ids, uint256[] amounts)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) WatchBackwardSwapFilled(opts *bind.WatchOpts, sink chan<- *ERC1155SwapAgentBackwardSwapFilled, swapTxHash [][32]byte, tokenAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC1155SwapAgent.contract.WatchLogs(opts, "BackwardSwapFilled", swapTxHashRule, tokenAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155SwapAgentBackwardSwapFilled)
				if err := _ERC1155SwapAgent.contract.UnpackLog(event, "BackwardSwapFilled", log); err != nil {
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

// ParseBackwardSwapFilled is a log parse operation binding the contract event 0x58d9c075708eb3187538135716481600d99ff6eb833bdbac1fee1e54cebec1b3.
//
// Solidity: event BackwardSwapFilled(bytes32 indexed swapTxHash, address indexed tokenAddr, address indexed recipient, uint256 fromChainId, uint256[] ids, uint256[] amounts)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) ParseBackwardSwapFilled(log types.Log) (*ERC1155SwapAgentBackwardSwapFilled, error) {
	event := new(ERC1155SwapAgentBackwardSwapFilled)
	if err := _ERC1155SwapAgent.contract.UnpackLog(event, "BackwardSwapFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155SwapAgentBackwardSwapStartedIterator is returned from FilterBackwardSwapStarted and is used to iterate over the raw logs and unpacked data for BackwardSwapStarted events raised by the ERC1155SwapAgent contract.
type ERC1155SwapAgentBackwardSwapStartedIterator struct {
	Event *ERC1155SwapAgentBackwardSwapStarted // Event containing the contract specifics and raw log

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
func (it *ERC1155SwapAgentBackwardSwapStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155SwapAgentBackwardSwapStarted)
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
		it.Event = new(ERC1155SwapAgentBackwardSwapStarted)
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
func (it *ERC1155SwapAgentBackwardSwapStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155SwapAgentBackwardSwapStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155SwapAgentBackwardSwapStarted represents a BackwardSwapStarted event raised by the ERC1155SwapAgent contract.
type ERC1155SwapAgentBackwardSwapStarted struct {
	MirroredTokenAddr common.Address
	Sender            common.Address
	Recipient         common.Address
	DstChainId        *big.Int
	Ids               []*big.Int
	Amounts           []*big.Int
	FeeAmount         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBackwardSwapStarted is a free log retrieval operation binding the contract event 0x5c317e3669ab4c20e7362c3bdc16700c64c83aa52a3abdd32e17eb1179e6706e.
//
// Solidity: event BackwardSwapStarted(address indexed mirroredTokenAddr, address indexed sender, address indexed recipient, uint256 dstChainId, uint256[] ids, uint256[] amounts, uint256 feeAmount)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) FilterBackwardSwapStarted(opts *bind.FilterOpts, mirroredTokenAddr []common.Address, sender []common.Address, recipient []common.Address) (*ERC1155SwapAgentBackwardSwapStartedIterator, error) {

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

	logs, sub, err := _ERC1155SwapAgent.contract.FilterLogs(opts, "BackwardSwapStarted", mirroredTokenAddrRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155SwapAgentBackwardSwapStartedIterator{contract: _ERC1155SwapAgent.contract, event: "BackwardSwapStarted", logs: logs, sub: sub}, nil
}

// WatchBackwardSwapStarted is a free log subscription operation binding the contract event 0x5c317e3669ab4c20e7362c3bdc16700c64c83aa52a3abdd32e17eb1179e6706e.
//
// Solidity: event BackwardSwapStarted(address indexed mirroredTokenAddr, address indexed sender, address indexed recipient, uint256 dstChainId, uint256[] ids, uint256[] amounts, uint256 feeAmount)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) WatchBackwardSwapStarted(opts *bind.WatchOpts, sink chan<- *ERC1155SwapAgentBackwardSwapStarted, mirroredTokenAddr []common.Address, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC1155SwapAgent.contract.WatchLogs(opts, "BackwardSwapStarted", mirroredTokenAddrRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155SwapAgentBackwardSwapStarted)
				if err := _ERC1155SwapAgent.contract.UnpackLog(event, "BackwardSwapStarted", log); err != nil {
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

// ParseBackwardSwapStarted is a log parse operation binding the contract event 0x5c317e3669ab4c20e7362c3bdc16700c64c83aa52a3abdd32e17eb1179e6706e.
//
// Solidity: event BackwardSwapStarted(address indexed mirroredTokenAddr, address indexed sender, address indexed recipient, uint256 dstChainId, uint256[] ids, uint256[] amounts, uint256 feeAmount)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) ParseBackwardSwapStarted(log types.Log) (*ERC1155SwapAgentBackwardSwapStarted, error) {
	event := new(ERC1155SwapAgentBackwardSwapStarted)
	if err := _ERC1155SwapAgent.contract.UnpackLog(event, "BackwardSwapStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155SwapAgentOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC1155SwapAgent contract.
type ERC1155SwapAgentOwnershipTransferredIterator struct {
	Event *ERC1155SwapAgentOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC1155SwapAgentOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155SwapAgentOwnershipTransferred)
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
		it.Event = new(ERC1155SwapAgentOwnershipTransferred)
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
func (it *ERC1155SwapAgentOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155SwapAgentOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155SwapAgentOwnershipTransferred represents a OwnershipTransferred event raised by the ERC1155SwapAgent contract.
type ERC1155SwapAgentOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC1155SwapAgentOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC1155SwapAgent.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155SwapAgentOwnershipTransferredIterator{contract: _ERC1155SwapAgent.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC1155SwapAgentOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC1155SwapAgent.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155SwapAgentOwnershipTransferred)
				if err := _ERC1155SwapAgent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) ParseOwnershipTransferred(log types.Log) (*ERC1155SwapAgentOwnershipTransferred, error) {
	event := new(ERC1155SwapAgentOwnershipTransferred)
	if err := _ERC1155SwapAgent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155SwapAgentSwapFilledIterator is returned from FilterSwapFilled and is used to iterate over the raw logs and unpacked data for SwapFilled events raised by the ERC1155SwapAgent contract.
type ERC1155SwapAgentSwapFilledIterator struct {
	Event *ERC1155SwapAgentSwapFilled // Event containing the contract specifics and raw log

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
func (it *ERC1155SwapAgentSwapFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155SwapAgentSwapFilled)
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
		it.Event = new(ERC1155SwapAgentSwapFilled)
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
func (it *ERC1155SwapAgentSwapFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155SwapAgentSwapFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155SwapAgentSwapFilled represents a SwapFilled event raised by the ERC1155SwapAgent contract.
type ERC1155SwapAgentSwapFilled struct {
	SwapTxHash        [32]byte
	FromTokenAddr     common.Address
	Recipient         common.Address
	MirroredTokenAddr common.Address
	FromChainId       *big.Int
	Ids               []*big.Int
	Amounts           []*big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapFilled is a free log retrieval operation binding the contract event 0x295d1e2c3b0c279f7107336cf70913e43ee3b0e77dee0b1d04f74d733006d806.
//
// Solidity: event SwapFilled(bytes32 indexed swapTxHash, address indexed fromTokenAddr, address indexed recipient, address mirroredTokenAddr, uint256 fromChainId, uint256[] ids, uint256[] amounts)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) FilterSwapFilled(opts *bind.FilterOpts, swapTxHash [][32]byte, fromTokenAddr []common.Address, recipient []common.Address) (*ERC1155SwapAgentSwapFilledIterator, error) {

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

	logs, sub, err := _ERC1155SwapAgent.contract.FilterLogs(opts, "SwapFilled", swapTxHashRule, fromTokenAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155SwapAgentSwapFilledIterator{contract: _ERC1155SwapAgent.contract, event: "SwapFilled", logs: logs, sub: sub}, nil
}

// WatchSwapFilled is a free log subscription operation binding the contract event 0x295d1e2c3b0c279f7107336cf70913e43ee3b0e77dee0b1d04f74d733006d806.
//
// Solidity: event SwapFilled(bytes32 indexed swapTxHash, address indexed fromTokenAddr, address indexed recipient, address mirroredTokenAddr, uint256 fromChainId, uint256[] ids, uint256[] amounts)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) WatchSwapFilled(opts *bind.WatchOpts, sink chan<- *ERC1155SwapAgentSwapFilled, swapTxHash [][32]byte, fromTokenAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC1155SwapAgent.contract.WatchLogs(opts, "SwapFilled", swapTxHashRule, fromTokenAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155SwapAgentSwapFilled)
				if err := _ERC1155SwapAgent.contract.UnpackLog(event, "SwapFilled", log); err != nil {
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

// ParseSwapFilled is a log parse operation binding the contract event 0x295d1e2c3b0c279f7107336cf70913e43ee3b0e77dee0b1d04f74d733006d806.
//
// Solidity: event SwapFilled(bytes32 indexed swapTxHash, address indexed fromTokenAddr, address indexed recipient, address mirroredTokenAddr, uint256 fromChainId, uint256[] ids, uint256[] amounts)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) ParseSwapFilled(log types.Log) (*ERC1155SwapAgentSwapFilled, error) {
	event := new(ERC1155SwapAgentSwapFilled)
	if err := _ERC1155SwapAgent.contract.UnpackLog(event, "SwapFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155SwapAgentSwapPairCreatedIterator is returned from FilterSwapPairCreated and is used to iterate over the raw logs and unpacked data for SwapPairCreated events raised by the ERC1155SwapAgent contract.
type ERC1155SwapAgentSwapPairCreatedIterator struct {
	Event *ERC1155SwapAgentSwapPairCreated // Event containing the contract specifics and raw log

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
func (it *ERC1155SwapAgentSwapPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155SwapAgentSwapPairCreated)
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
		it.Event = new(ERC1155SwapAgentSwapPairCreated)
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
func (it *ERC1155SwapAgentSwapPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155SwapAgentSwapPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155SwapAgentSwapPairCreated represents a SwapPairCreated event raised by the ERC1155SwapAgent contract.
type ERC1155SwapAgentSwapPairCreated struct {
	RegisterTxHash    [32]byte
	FromTokenAddr     common.Address
	MirroredTokenAddr common.Address
	FromChainId       *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapPairCreated is a free log retrieval operation binding the contract event 0x621d9726b59bf7f3a9cbd292df8310172e6dec3a7279c906c7c22129f406708e.
//
// Solidity: event SwapPairCreated(bytes32 indexed registerTxHash, address indexed fromTokenAddr, address indexed mirroredTokenAddr, uint256 fromChainId)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) FilterSwapPairCreated(opts *bind.FilterOpts, registerTxHash [][32]byte, fromTokenAddr []common.Address, mirroredTokenAddr []common.Address) (*ERC1155SwapAgentSwapPairCreatedIterator, error) {

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

	logs, sub, err := _ERC1155SwapAgent.contract.FilterLogs(opts, "SwapPairCreated", registerTxHashRule, fromTokenAddrRule, mirroredTokenAddrRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155SwapAgentSwapPairCreatedIterator{contract: _ERC1155SwapAgent.contract, event: "SwapPairCreated", logs: logs, sub: sub}, nil
}

// WatchSwapPairCreated is a free log subscription operation binding the contract event 0x621d9726b59bf7f3a9cbd292df8310172e6dec3a7279c906c7c22129f406708e.
//
// Solidity: event SwapPairCreated(bytes32 indexed registerTxHash, address indexed fromTokenAddr, address indexed mirroredTokenAddr, uint256 fromChainId)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) WatchSwapPairCreated(opts *bind.WatchOpts, sink chan<- *ERC1155SwapAgentSwapPairCreated, registerTxHash [][32]byte, fromTokenAddr []common.Address, mirroredTokenAddr []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC1155SwapAgent.contract.WatchLogs(opts, "SwapPairCreated", registerTxHashRule, fromTokenAddrRule, mirroredTokenAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155SwapAgentSwapPairCreated)
				if err := _ERC1155SwapAgent.contract.UnpackLog(event, "SwapPairCreated", log); err != nil {
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

// ParseSwapPairCreated is a log parse operation binding the contract event 0x621d9726b59bf7f3a9cbd292df8310172e6dec3a7279c906c7c22129f406708e.
//
// Solidity: event SwapPairCreated(bytes32 indexed registerTxHash, address indexed fromTokenAddr, address indexed mirroredTokenAddr, uint256 fromChainId)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) ParseSwapPairCreated(log types.Log) (*ERC1155SwapAgentSwapPairCreated, error) {
	event := new(ERC1155SwapAgentSwapPairCreated)
	if err := _ERC1155SwapAgent.contract.UnpackLog(event, "SwapPairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155SwapAgentSwapPairRegisterIterator is returned from FilterSwapPairRegister and is used to iterate over the raw logs and unpacked data for SwapPairRegister events raised by the ERC1155SwapAgent contract.
type ERC1155SwapAgentSwapPairRegisterIterator struct {
	Event *ERC1155SwapAgentSwapPairRegister // Event containing the contract specifics and raw log

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
func (it *ERC1155SwapAgentSwapPairRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155SwapAgentSwapPairRegister)
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
		it.Event = new(ERC1155SwapAgentSwapPairRegister)
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
func (it *ERC1155SwapAgentSwapPairRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155SwapAgentSwapPairRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155SwapAgentSwapPairRegister represents a SwapPairRegister event raised by the ERC1155SwapAgent contract.
type ERC1155SwapAgentSwapPairRegister struct {
	Sponsor      common.Address
	TokenAddress common.Address
	ToChainId    *big.Int
	FeeAmount    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwapPairRegister is a free log retrieval operation binding the contract event 0x6745f5c9e689dd3b252d145964ac6fbb294d72cedae5a4d52b1774728289e606.
//
// Solidity: event SwapPairRegister(address indexed sponsor, address indexed tokenAddress, uint256 toChainId, uint256 feeAmount)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) FilterSwapPairRegister(opts *bind.FilterOpts, sponsor []common.Address, tokenAddress []common.Address) (*ERC1155SwapAgentSwapPairRegisterIterator, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _ERC1155SwapAgent.contract.FilterLogs(opts, "SwapPairRegister", sponsorRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155SwapAgentSwapPairRegisterIterator{contract: _ERC1155SwapAgent.contract, event: "SwapPairRegister", logs: logs, sub: sub}, nil
}

// WatchSwapPairRegister is a free log subscription operation binding the contract event 0x6745f5c9e689dd3b252d145964ac6fbb294d72cedae5a4d52b1774728289e606.
//
// Solidity: event SwapPairRegister(address indexed sponsor, address indexed tokenAddress, uint256 toChainId, uint256 feeAmount)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) WatchSwapPairRegister(opts *bind.WatchOpts, sink chan<- *ERC1155SwapAgentSwapPairRegister, sponsor []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _ERC1155SwapAgent.contract.WatchLogs(opts, "SwapPairRegister", sponsorRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155SwapAgentSwapPairRegister)
				if err := _ERC1155SwapAgent.contract.UnpackLog(event, "SwapPairRegister", log); err != nil {
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

// ParseSwapPairRegister is a log parse operation binding the contract event 0x6745f5c9e689dd3b252d145964ac6fbb294d72cedae5a4d52b1774728289e606.
//
// Solidity: event SwapPairRegister(address indexed sponsor, address indexed tokenAddress, uint256 toChainId, uint256 feeAmount)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) ParseSwapPairRegister(log types.Log) (*ERC1155SwapAgentSwapPairRegister, error) {
	event := new(ERC1155SwapAgentSwapPairRegister)
	if err := _ERC1155SwapAgent.contract.UnpackLog(event, "SwapPairRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155SwapAgentSwapStartedIterator is returned from FilterSwapStarted and is used to iterate over the raw logs and unpacked data for SwapStarted events raised by the ERC1155SwapAgent contract.
type ERC1155SwapAgentSwapStartedIterator struct {
	Event *ERC1155SwapAgentSwapStarted // Event containing the contract specifics and raw log

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
func (it *ERC1155SwapAgentSwapStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155SwapAgentSwapStarted)
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
		it.Event = new(ERC1155SwapAgentSwapStarted)
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
func (it *ERC1155SwapAgentSwapStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155SwapAgentSwapStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155SwapAgentSwapStarted represents a SwapStarted event raised by the ERC1155SwapAgent contract.
type ERC1155SwapAgentSwapStarted struct {
	TokenAddr  common.Address
	Sender     common.Address
	Recipient  common.Address
	DstChainId *big.Int
	Ids        []*big.Int
	Amounts    []*big.Int
	FeeAmount  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwapStarted is a free log retrieval operation binding the contract event 0x074135076f5fc18420e0de96ce28c6bfe93048607465c707bdf3c10cf1e92d3f.
//
// Solidity: event SwapStarted(address indexed tokenAddr, address indexed sender, address indexed recipient, uint256 dstChainId, uint256[] ids, uint256[] amounts, uint256 feeAmount)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) FilterSwapStarted(opts *bind.FilterOpts, tokenAddr []common.Address, sender []common.Address, recipient []common.Address) (*ERC1155SwapAgentSwapStartedIterator, error) {

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

	logs, sub, err := _ERC1155SwapAgent.contract.FilterLogs(opts, "SwapStarted", tokenAddrRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155SwapAgentSwapStartedIterator{contract: _ERC1155SwapAgent.contract, event: "SwapStarted", logs: logs, sub: sub}, nil
}

// WatchSwapStarted is a free log subscription operation binding the contract event 0x074135076f5fc18420e0de96ce28c6bfe93048607465c707bdf3c10cf1e92d3f.
//
// Solidity: event SwapStarted(address indexed tokenAddr, address indexed sender, address indexed recipient, uint256 dstChainId, uint256[] ids, uint256[] amounts, uint256 feeAmount)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) WatchSwapStarted(opts *bind.WatchOpts, sink chan<- *ERC1155SwapAgentSwapStarted, tokenAddr []common.Address, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC1155SwapAgent.contract.WatchLogs(opts, "SwapStarted", tokenAddrRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155SwapAgentSwapStarted)
				if err := _ERC1155SwapAgent.contract.UnpackLog(event, "SwapStarted", log); err != nil {
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

// ParseSwapStarted is a log parse operation binding the contract event 0x074135076f5fc18420e0de96ce28c6bfe93048607465c707bdf3c10cf1e92d3f.
//
// Solidity: event SwapStarted(address indexed tokenAddr, address indexed sender, address indexed recipient, uint256 dstChainId, uint256[] ids, uint256[] amounts, uint256 feeAmount)
func (_ERC1155SwapAgent *ERC1155SwapAgentFilterer) ParseSwapStarted(log types.Log) (*ERC1155SwapAgentSwapStarted, error) {
	event := new(ERC1155SwapAgentSwapStarted)
	if err := _ERC1155SwapAgent.contract.UnpackLog(event, "SwapStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
