// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VidarCashCard

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
	_ = abi.ConvertType
)

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hexHash\",\"type\":\"string\"}],\"name\":\"changeOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orgMsg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_avaxAdd\",\"type\":\"string\"}],\"name\":\"completeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimateGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHexHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwnerTransferStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hex_hash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hexHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_orgMsg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_avaxAdd\",\"type\":\"string\"}],\"name\":\"modifyOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_orgMsg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_avaxAdd\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"sendToContract\",\"type\":\"bool\"}],\"name\":\"sendCash\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// CurrentOwner is a free data retrieval call binding the contract method 0xb387ef92.
//
// Solidity: function currentOwner() view returns(address)
func (_Main *MainCaller) CurrentOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "currentOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentOwner is a free data retrieval call binding the contract method 0xb387ef92.
//
// Solidity: function currentOwner() view returns(address)
func (_Main *MainSession) CurrentOwner() (common.Address, error) {
	return _Main.Contract.CurrentOwner(&_Main.CallOpts)
}

// CurrentOwner is a free data retrieval call binding the contract method 0xb387ef92.
//
// Solidity: function currentOwner() view returns(address)
func (_Main *MainCallerSession) CurrentOwner() (common.Address, error) {
	return _Main.Contract.CurrentOwner(&_Main.CallOpts)
}

// EstimateGasFee is a free data retrieval call binding the contract method 0x861458d8.
//
// Solidity: function estimateGasFee() view returns(uint256)
func (_Main *MainCaller) EstimateGasFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "estimateGasFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateGasFee is a free data retrieval call binding the contract method 0x861458d8.
//
// Solidity: function estimateGasFee() view returns(uint256)
func (_Main *MainSession) EstimateGasFee() (*big.Int, error) {
	return _Main.Contract.EstimateGasFee(&_Main.CallOpts)
}

// EstimateGasFee is a free data retrieval call binding the contract method 0x861458d8.
//
// Solidity: function estimateGasFee() view returns(uint256)
func (_Main *MainCallerSession) EstimateGasFee() (*big.Int, error) {
	return _Main.Contract.EstimateGasFee(&_Main.CallOpts)
}

// GetCurrentOwner is a free data retrieval call binding the contract method 0xa18a186b.
//
// Solidity: function getCurrentOwner() view returns(address)
func (_Main *MainCaller) GetCurrentOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getCurrentOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentOwner is a free data retrieval call binding the contract method 0xa18a186b.
//
// Solidity: function getCurrentOwner() view returns(address)
func (_Main *MainSession) GetCurrentOwner() (common.Address, error) {
	return _Main.Contract.GetCurrentOwner(&_Main.CallOpts)
}

// GetCurrentOwner is a free data retrieval call binding the contract method 0xa18a186b.
//
// Solidity: function getCurrentOwner() view returns(address)
func (_Main *MainCallerSession) GetCurrentOwner() (common.Address, error) {
	return _Main.Contract.GetCurrentOwner(&_Main.CallOpts)
}

// GetHexHash is a free data retrieval call binding the contract method 0xc944f107.
//
// Solidity: function getHexHash() view returns(string)
func (_Main *MainCaller) GetHexHash(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getHexHash")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHexHash is a free data retrieval call binding the contract method 0xc944f107.
//
// Solidity: function getHexHash() view returns(string)
func (_Main *MainSession) GetHexHash() (string, error) {
	return _Main.Contract.GetHexHash(&_Main.CallOpts)
}

// GetHexHash is a free data retrieval call binding the contract method 0xc944f107.
//
// Solidity: function getHexHash() view returns(string)
func (_Main *MainCallerSession) GetHexHash() (string, error) {
	return _Main.Contract.GetHexHash(&_Main.CallOpts)
}

// GetOwnerTransferStatus is a free data retrieval call binding the contract method 0x31b1638a.
//
// Solidity: function getOwnerTransferStatus() view returns(bool)
func (_Main *MainCaller) GetOwnerTransferStatus(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getOwnerTransferStatus")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetOwnerTransferStatus is a free data retrieval call binding the contract method 0x31b1638a.
//
// Solidity: function getOwnerTransferStatus() view returns(bool)
func (_Main *MainSession) GetOwnerTransferStatus() (bool, error) {
	return _Main.Contract.GetOwnerTransferStatus(&_Main.CallOpts)
}

// GetOwnerTransferStatus is a free data retrieval call binding the contract method 0x31b1638a.
//
// Solidity: function getOwnerTransferStatus() view returns(bool)
func (_Main *MainCallerSession) GetOwnerTransferStatus() (bool, error) {
	return _Main.Contract.GetOwnerTransferStatus(&_Main.CallOpts)
}

// HexHash is a free data retrieval call binding the contract method 0xe50b5280.
//
// Solidity: function hex_hash() view returns(string)
func (_Main *MainCaller) HexHash(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "hex_hash")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// HexHash is a free data retrieval call binding the contract method 0xe50b5280.
//
// Solidity: function hex_hash() view returns(string)
func (_Main *MainSession) HexHash() (string, error) {
	return _Main.Contract.HexHash(&_Main.CallOpts)
}

// HexHash is a free data retrieval call binding the contract method 0xe50b5280.
//
// Solidity: function hex_hash() view returns(string)
func (_Main *MainCallerSession) HexHash() (string, error) {
	return _Main.Contract.HexHash(&_Main.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xfaec4db4.
//
// Solidity: function changeOwner(string _hexHash) returns(bool)
func (_Main *MainTransactor) ChangeOwner(opts *bind.TransactOpts, _hexHash string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "changeOwner", _hexHash)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xfaec4db4.
//
// Solidity: function changeOwner(string _hexHash) returns(bool)
func (_Main *MainSession) ChangeOwner(_hexHash string) (*types.Transaction, error) {
	return _Main.Contract.ChangeOwner(&_Main.TransactOpts, _hexHash)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xfaec4db4.
//
// Solidity: function changeOwner(string _hexHash) returns(bool)
func (_Main *MainTransactorSession) ChangeOwner(_hexHash string) (*types.Transaction, error) {
	return _Main.Contract.ChangeOwner(&_Main.TransactOpts, _hexHash)
}

// CompleteTransfer is a paid mutator transaction binding the contract method 0xbb6d1f1e.
//
// Solidity: function completeTransfer(string _orgMsg, string _avaxAdd) returns()
func (_Main *MainTransactor) CompleteTransfer(opts *bind.TransactOpts, _orgMsg string, _avaxAdd string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "completeTransfer", _orgMsg, _avaxAdd)
}

// CompleteTransfer is a paid mutator transaction binding the contract method 0xbb6d1f1e.
//
// Solidity: function completeTransfer(string _orgMsg, string _avaxAdd) returns()
func (_Main *MainSession) CompleteTransfer(_orgMsg string, _avaxAdd string) (*types.Transaction, error) {
	return _Main.Contract.CompleteTransfer(&_Main.TransactOpts, _orgMsg, _avaxAdd)
}

// CompleteTransfer is a paid mutator transaction binding the contract method 0xbb6d1f1e.
//
// Solidity: function completeTransfer(string _orgMsg, string _avaxAdd) returns()
func (_Main *MainTransactorSession) CompleteTransfer(_orgMsg string, _avaxAdd string) (*types.Transaction, error) {
	return _Main.Contract.CompleteTransfer(&_Main.TransactOpts, _orgMsg, _avaxAdd)
}

// ModifyOwner is a paid mutator transaction binding the contract method 0xe1fb7724.
//
// Solidity: function modifyOwner(string _hexHash, string _orgMsg, string _avaxAdd) returns()
func (_Main *MainTransactor) ModifyOwner(opts *bind.TransactOpts, _hexHash string, _orgMsg string, _avaxAdd string) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "modifyOwner", _hexHash, _orgMsg, _avaxAdd)
}

// ModifyOwner is a paid mutator transaction binding the contract method 0xe1fb7724.
//
// Solidity: function modifyOwner(string _hexHash, string _orgMsg, string _avaxAdd) returns()
func (_Main *MainSession) ModifyOwner(_hexHash string, _orgMsg string, _avaxAdd string) (*types.Transaction, error) {
	return _Main.Contract.ModifyOwner(&_Main.TransactOpts, _hexHash, _orgMsg, _avaxAdd)
}

// ModifyOwner is a paid mutator transaction binding the contract method 0xe1fb7724.
//
// Solidity: function modifyOwner(string _hexHash, string _orgMsg, string _avaxAdd) returns()
func (_Main *MainTransactorSession) ModifyOwner(_hexHash string, _orgMsg string, _avaxAdd string) (*types.Transaction, error) {
	return _Main.Contract.ModifyOwner(&_Main.TransactOpts, _hexHash, _orgMsg, _avaxAdd)
}

// SendCash is a paid mutator transaction binding the contract method 0x06e5084d.
//
// Solidity: function sendCash(address to, uint256 amount, string _orgMsg, string _avaxAdd, bool sendToContract) payable returns()
func (_Main *MainTransactor) SendCash(opts *bind.TransactOpts, to common.Address, amount *big.Int, _orgMsg string, _avaxAdd string, sendToContract bool) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "sendCash", to, amount, _orgMsg, _avaxAdd, sendToContract)
}

// SendCash is a paid mutator transaction binding the contract method 0x06e5084d.
//
// Solidity: function sendCash(address to, uint256 amount, string _orgMsg, string _avaxAdd, bool sendToContract) payable returns()
func (_Main *MainSession) SendCash(to common.Address, amount *big.Int, _orgMsg string, _avaxAdd string, sendToContract bool) (*types.Transaction, error) {
	return _Main.Contract.SendCash(&_Main.TransactOpts, to, amount, _orgMsg, _avaxAdd, sendToContract)
}

// SendCash is a paid mutator transaction binding the contract method 0x06e5084d.
//
// Solidity: function sendCash(address to, uint256 amount, string _orgMsg, string _avaxAdd, bool sendToContract) payable returns()
func (_Main *MainTransactorSession) SendCash(to common.Address, amount *big.Int, _orgMsg string, _avaxAdd string, sendToContract bool) (*types.Transaction, error) {
	return _Main.Contract.SendCash(&_Main.TransactOpts, to, amount, _orgMsg, _avaxAdd, sendToContract)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Main *MainTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Main *MainSession) Receive() (*types.Transaction, error) {
	return _Main.Contract.Receive(&_Main.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Main *MainTransactorSession) Receive() (*types.Transaction, error) {
	return _Main.Contract.Receive(&_Main.TransactOpts)
}
// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, ApiBin string) (common.Address, *types.Transaction, *Api, error) {
	
	address, tx, contract, err := bind.DeployContract(auth, *MainMetaData.ABI, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}