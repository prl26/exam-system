// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package targetAbi

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

// TargetMetaData contains all meta data concerning the Target contract.
var TargetMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"score\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TargetABI is the input ABI used to generate the binding from.
// Deprecated: Use TargetMetaData.ABI instead.
var TargetABI = TargetMetaData.ABI

// Target is an auto generated Go binding around an Ethereum contract.
type Target struct {
	TargetCaller     // Read-only binding to the contract
	TargetTransactor // Write-only binding to the contract
	TargetFilterer   // Log filterer for contract events
}

// TargetCaller is an auto generated read-only Go binding around an Ethereum contract.
type TargetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TargetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TargetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TargetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TargetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TargetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TargetSession struct {
	Contract     *Target           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TargetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TargetCallerSession struct {
	Contract *TargetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TargetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TargetTransactorSession struct {
	Contract     *TargetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TargetRaw is an auto generated low-level Go binding around an Ethereum contract.
type TargetRaw struct {
	Contract *Target // Generic contract binding to access the raw methods on
}

// TargetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TargetCallerRaw struct {
	Contract *TargetCaller // Generic read-only contract binding to access the raw methods on
}

// TargetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TargetTransactorRaw struct {
	Contract *TargetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTarget creates a new instance of Target, bound to a specific deployed contract.
func NewTarget(address common.Address, backend bind.ContractBackend) (*Target, error) {
	contract, err := bindTarget(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Target{TargetCaller: TargetCaller{contract: contract}, TargetTransactor: TargetTransactor{contract: contract}, TargetFilterer: TargetFilterer{contract: contract}}, nil
}

// NewTargetCaller creates a new read-only instance of Target, bound to a specific deployed contract.
func NewTargetCaller(address common.Address, caller bind.ContractCaller) (*TargetCaller, error) {
	contract, err := bindTarget(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TargetCaller{contract: contract}, nil
}

// NewTargetTransactor creates a new write-only instance of Target, bound to a specific deployed contract.
func NewTargetTransactor(address common.Address, transactor bind.ContractTransactor) (*TargetTransactor, error) {
	contract, err := bindTarget(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TargetTransactor{contract: contract}, nil
}

// NewTargetFilterer creates a new log filterer instance of Target, bound to a specific deployed contract.
func NewTargetFilterer(address common.Address, filterer bind.ContractFilterer) (*TargetFilterer, error) {
	contract, err := bindTarget(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TargetFilterer{contract: contract}, nil
}

// bindTarget binds a generic wrapper to an already deployed contract.
func bindTarget(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TargetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Target *TargetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Target.Contract.TargetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Target *TargetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Target.Contract.TargetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Target *TargetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Target.Contract.TargetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Target *TargetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Target.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Target *TargetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Target.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Target *TargetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Target.Contract.contract.Transact(opts, method, params...)
}

// Score is a free data retrieval call binding the contract method 0xefedc669.
//
// Solidity: function score() view returns(uint256)
func (_Target *TargetCaller) Score(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Target.contract.Call(opts, &out, "score")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Score is a free data retrieval call binding the contract method 0xefedc669.
//
// Solidity: function score() view returns(uint256)
func (_Target *TargetSession) Score() (*big.Int, error) {
	return _Target.Contract.Score(&_Target.CallOpts)
}

// Score is a free data retrieval call binding the contract method 0xefedc669.
//
// Solidity: function score() view returns(uint256)
func (_Target *TargetCallerSession) Score() (*big.Int, error) {
	return _Target.Contract.Score(&_Target.CallOpts)
}
