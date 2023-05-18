// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// VidarUtilslice is an auto generated low-level Go binding around an user-defined struct.
type VidarUtilslice struct {
	Len *big.Int
	Ptr *big.Int
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"address_to_string\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToHexString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"_char\",\"type\":\"bytes1\"}],\"name\":\"charToHexChar\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"a\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"b\",\"type\":\"string\"}],\"name\":\"compare\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"estimateGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_string1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_string2\",\"type\":\"string\"}],\"name\":\"getHex_Hash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_str\",\"type\":\"string\"}],\"name\":\"getSlice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_len\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ptr\",\"type\":\"uint256\"}],\"internalType\":\"structVidarUtil.slice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isValidAVAXAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"toHexString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x610dcc61003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100b35760003560e01c806363e1cbea1161007b57806363e1cbea14610163578063861458d814610176578063908c0eea1461018c578063a3fa67ed146101e2578063c68b3787146101f5578063dd005a831461020857600080fd5b806317184d12146100b8578063274695d3146100e95780633a96fdd714610109578063531915031461013d57806357b7866d14610150575b600080fd5b6100cb6100c636600461085a565b61021b565b6040516001600160f81b031990911681526020015b60405180910390f35b6100fc6100f7366004610884565b610251565b6040516100e091906108d1565b61012d6101173660046109a7565b8051602091820120825192909101919091201490565b60405190151581526020016100e0565b6100fc61014b366004610884565b6102ac565b61012d61015e366004610a0b565b610445565b6100fc610171366004610a78565b6104e6565b61017e61068d565b6040519081526020016100e0565b6101c761019a366004610a9a565b60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518251815260209283015192810192909252016100e0565b61012d6101f0366004610884565b3b1590565b6100fc610203366004610acf565b6106c4565b6100fc6102163660046109a7565b6107be565b6000600a60f883901c10156102425761023960f883901c6030610afe565b60f81b92915050565b61023960f883901c6057610afe565b6060813b6102735761026d6001600160a01b03831660146104e6565b92915050565b505060408051808201909152601881527f4e6f7420612076616c6964204176617820616464726573730000000000000000602082015290565b60408051602a808252606082810190935260009190602082018180368337019050509050600360fc1b816000815181106102e8576102e8610b17565b60200101906001600160f81b031916908160001a905350600f60fb1b8160018151811061031757610317610b17565b60200101906001600160f81b031916908160001a90535060005b601481101561043e576000610347826013610b2d565b610352906008610b40565b61035d906002610c3b565b610370906001600160a01b038716610c5d565b9050600061037f601083610c71565b9050600061038e826010610c93565b6103989084610caf565b90506103a3826107f8565b856103af866002610b40565b6103ba906002610cc8565b815181106103ca576103ca610b17565b60200101906001600160f81b031916908160001a9053506103ea816107f8565b856103f6866002610b40565b610401906003610cc8565b8151811061041157610411610b17565b60200101906001600160f81b031916908160001a905350505050808061043690610cdb565b915050610331565b5092915050565b60006104de84805461045690610cf4565b80601f016020809104026020016040519081016040528092919081815260200182805461048290610cf4565b80156104cf5780601f106104a4576101008083540402835291602001916104cf565b820191906000526020600020905b8154815290600101906020018083116104b257829003601f168201915b505050505061011785856107be565b949350505050565b606060006104f5836002610b40565b610500906002610cc8565b67ffffffffffffffff81111561051857610518610904565b6040519080825280601f01601f191660200182016040528015610542576020820181803683370190505b509050600360fc1b8160008151811061055d5761055d610b17565b60200101906001600160f81b031916908160001a905350600f60fb1b8160018151811061058c5761058c610b17565b60200101906001600160f81b031916908160001a90535060006105b0846002610b40565b6105bb906001610cc8565b90505b6001811115610633576f181899199a1a9b1b9c1cb0b131b232b360811b85600f16601081106105ef576105ef610b17565b1a60f81b82828151811061060557610605610b17565b60200101906001600160f81b031916908160001a90535060049490941c9361062c81610d2e565b90506105be565b5083156106865760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640160405180910390fd5b9392505050565b6000805a90503a60005a6106a19084610b2d565b6106ad9061a410610cc8565b905060006106bb8383610b40565b95945050505050565b6040805181815260608181018352916000919060208201818036833701905050905060005b602081101561043e57600084826020811061070657610706610b17565b1a90506000610721610719601084610c71565b60f81b61021b565b90506000610733610719601085610d45565b90508185610742866002610b40565b8151811061075257610752610b17565b60200101906001600160f81b031916908160001a9053508085610776866002610b40565b610781906001610cc8565b8151811061079157610791610b17565b60200101906001600160f81b031916908160001a90535050505080806107b690610cdb565b9150506106e9565b6060600083836040516020016107d5929190610d67565b60408051601f19818403018152919052805160208201209091506106bb816106c4565b6000600a8260ff16101561082a576040805180820190915260018152600360fc1b602090910152610239826030610afe565b6040805180820190915260018152606160f81b602090910152600a610850836061610afe565b6102399190610caf565b60006020828403121561086c57600080fd5b81356001600160f81b03198116811461068657600080fd5b60006020828403121561089657600080fd5b81356001600160a01b038116811461068657600080fd5b60005b838110156108c85781810151838201526020016108b0565b50506000910152565b60208152600082518060208401526108f08160408501602087016108ad565b601f01601f19169190910160400192915050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261092b57600080fd5b813567ffffffffffffffff8082111561094657610946610904565b604051601f8301601f19908116603f0116810190828211818310171561096e5761096e610904565b8160405283815286602085880101111561098757600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080604083850312156109ba57600080fd5b823567ffffffffffffffff808211156109d257600080fd5b6109de8683870161091a565b935060208501359150808211156109f457600080fd5b50610a018582860161091a565b9150509250929050565b600080600060608486031215610a2057600080fd5b83359250602084013567ffffffffffffffff80821115610a3f57600080fd5b610a4b8783880161091a565b93506040860135915080821115610a6157600080fd5b50610a6e8682870161091a565b9150509250925092565b60008060408385031215610a8b57600080fd5b50508035926020909101359150565b600060208284031215610aac57600080fd5b813567ffffffffffffffff811115610ac357600080fd5b6104de8482850161091a565b600060208284031215610ae157600080fd5b5035919050565b634e487b7160e01b600052601160045260246000fd5b60ff818116838216019081111561026d5761026d610ae8565b634e487b7160e01b600052603260045260246000fd5b8181038181111561026d5761026d610ae8565b808202811582820484141761026d5761026d610ae8565b600181815b80851115610b92578160001904821115610b7857610b78610ae8565b80851615610b8557918102915b93841c9390800290610b5c565b509250929050565b600082610ba95750600161026d565b81610bb65750600061026d565b8160018114610bcc5760028114610bd657610bf2565b600191505061026d565b60ff841115610be757610be7610ae8565b50506001821b61026d565b5060208310610133831016604e8410600b8410161715610c15575081810a61026d565b610c1f8383610b57565b8060001904821115610c3357610c33610ae8565b029392505050565b60006106868383610b9a565b634e487b7160e01b600052601260045260246000fd5b600082610c6c57610c6c610c47565b500490565b600060ff831680610c8457610c84610c47565b8060ff84160491505092915050565b60ff818116838216029081169081811461043e5761043e610ae8565b60ff828116828216039081111561026d5761026d610ae8565b8082018082111561026d5761026d610ae8565b600060018201610ced57610ced610ae8565b5060010190565b600181811c90821680610d0857607f821691505b602082108103610d2857634e487b7160e01b600052602260045260246000fd5b50919050565b600081610d3d57610d3d610ae8565b506000190190565b600060ff831680610d5857610d58610c47565b8060ff84160691505092915050565b60008351610d798184602088016108ad565b835190830190610d8d8183602088016108ad565b0194935050505056fea264697066735822122021bb2e6d250e874f595f9eb5b4a500e2ce39be353d63b36d1bc6d14453b0658a64736f6c63430008120033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// AddToString is a free data retrieval call binding the contract method 0x53191503.
//
// Solidity: function addToString(address _address) pure returns(string)
func (_Api *ApiCaller) AddToString(opts *bind.CallOpts, _address common.Address) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "addToString", _address)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AddToString is a free data retrieval call binding the contract method 0x53191503.
//
// Solidity: function addToString(address _address) pure returns(string)
func (_Api *ApiSession) AddToString(_address common.Address) (string, error) {
	return _Api.Contract.AddToString(&_Api.CallOpts, _address)
}

// AddToString is a free data retrieval call binding the contract method 0x53191503.
//
// Solidity: function addToString(address _address) pure returns(string)
func (_Api *ApiCallerSession) AddToString(_address common.Address) (string, error) {
	return _Api.Contract.AddToString(&_Api.CallOpts, _address)
}

// AddressToString is a free data retrieval call binding the contract method 0x274695d3.
//
// Solidity: function address_to_string(address addr) view returns(string)
func (_Api *ApiCaller) AddressToString(opts *bind.CallOpts, addr common.Address) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "address_to_string", addr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AddressToString is a free data retrieval call binding the contract method 0x274695d3.
//
// Solidity: function address_to_string(address addr) view returns(string)
func (_Api *ApiSession) AddressToString(addr common.Address) (string, error) {
	return _Api.Contract.AddressToString(&_Api.CallOpts, addr)
}

// AddressToString is a free data retrieval call binding the contract method 0x274695d3.
//
// Solidity: function address_to_string(address addr) view returns(string)
func (_Api *ApiCallerSession) AddressToString(addr common.Address) (string, error) {
	return _Api.Contract.AddressToString(&_Api.CallOpts, addr)
}

// Bytes32ToHexString is a free data retrieval call binding the contract method 0xc68b3787.
//
// Solidity: function bytes32ToHexString(bytes32 _bytes32) pure returns(string)
func (_Api *ApiCaller) Bytes32ToHexString(opts *bind.CallOpts, _bytes32 [32]byte) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "bytes32ToHexString", _bytes32)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bytes32ToHexString is a free data retrieval call binding the contract method 0xc68b3787.
//
// Solidity: function bytes32ToHexString(bytes32 _bytes32) pure returns(string)
func (_Api *ApiSession) Bytes32ToHexString(_bytes32 [32]byte) (string, error) {
	return _Api.Contract.Bytes32ToHexString(&_Api.CallOpts, _bytes32)
}

// Bytes32ToHexString is a free data retrieval call binding the contract method 0xc68b3787.
//
// Solidity: function bytes32ToHexString(bytes32 _bytes32) pure returns(string)
func (_Api *ApiCallerSession) Bytes32ToHexString(_bytes32 [32]byte) (string, error) {
	return _Api.Contract.Bytes32ToHexString(&_Api.CallOpts, _bytes32)
}

// CharToHexChar is a free data retrieval call binding the contract method 0x17184d12.
//
// Solidity: function charToHexChar(bytes1 _char) pure returns(bytes1)
func (_Api *ApiCaller) CharToHexChar(opts *bind.CallOpts, _char [1]byte) ([1]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "charToHexChar", _char)

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// CharToHexChar is a free data retrieval call binding the contract method 0x17184d12.
//
// Solidity: function charToHexChar(bytes1 _char) pure returns(bytes1)
func (_Api *ApiSession) CharToHexChar(_char [1]byte) ([1]byte, error) {
	return _Api.Contract.CharToHexChar(&_Api.CallOpts, _char)
}

// CharToHexChar is a free data retrieval call binding the contract method 0x17184d12.
//
// Solidity: function charToHexChar(bytes1 _char) pure returns(bytes1)
func (_Api *ApiCallerSession) CharToHexChar(_char [1]byte) ([1]byte, error) {
	return _Api.Contract.CharToHexChar(&_Api.CallOpts, _char)
}

// Compare is a free data retrieval call binding the contract method 0x3a96fdd7.
//
// Solidity: function compare(string a, string b) pure returns(bool)
func (_Api *ApiCaller) Compare(opts *bind.CallOpts, a string, b string) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "compare", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Compare is a free data retrieval call binding the contract method 0x3a96fdd7.
//
// Solidity: function compare(string a, string b) pure returns(bool)
func (_Api *ApiSession) Compare(a string, b string) (bool, error) {
	return _Api.Contract.Compare(&_Api.CallOpts, a, b)
}

// Compare is a free data retrieval call binding the contract method 0x3a96fdd7.
//
// Solidity: function compare(string a, string b) pure returns(bool)
func (_Api *ApiCallerSession) Compare(a string, b string) (bool, error) {
	return _Api.Contract.Compare(&_Api.CallOpts, a, b)
}

// EstimateGasFee is a free data retrieval call binding the contract method 0x861458d8.
//
// Solidity: function estimateGasFee() view returns(uint256)
func (_Api *ApiCaller) EstimateGasFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "estimateGasFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateGasFee is a free data retrieval call binding the contract method 0x861458d8.
//
// Solidity: function estimateGasFee() view returns(uint256)
func (_Api *ApiSession) EstimateGasFee() (*big.Int, error) {
	return _Api.Contract.EstimateGasFee(&_Api.CallOpts)
}

// EstimateGasFee is a free data retrieval call binding the contract method 0x861458d8.
//
// Solidity: function estimateGasFee() view returns(uint256)
func (_Api *ApiCallerSession) EstimateGasFee() (*big.Int, error) {
	return _Api.Contract.EstimateGasFee(&_Api.CallOpts)
}

// GetHexHash is a free data retrieval call binding the contract method 0xdd005a83.
//
// Solidity: function getHex_Hash(string _string1, string _string2) pure returns(string)
func (_Api *ApiCaller) GetHexHash(opts *bind.CallOpts, _string1 string, _string2 string) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getHex_Hash", _string1, _string2)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHexHash is a free data retrieval call binding the contract method 0xdd005a83.
//
// Solidity: function getHex_Hash(string _string1, string _string2) pure returns(string)
func (_Api *ApiSession) GetHexHash(_string1 string, _string2 string) (string, error) {
	return _Api.Contract.GetHexHash(&_Api.CallOpts, _string1, _string2)
}

// GetHexHash is a free data retrieval call binding the contract method 0xdd005a83.
//
// Solidity: function getHex_Hash(string _string1, string _string2) pure returns(string)
func (_Api *ApiCallerSession) GetHexHash(_string1 string, _string2 string) (string, error) {
	return _Api.Contract.GetHexHash(&_Api.CallOpts, _string1, _string2)
}

// GetSlice is a free data retrieval call binding the contract method 0x908c0eea.
//
// Solidity: function getSlice(string _str) pure returns((uint256,uint256))
func (_Api *ApiCaller) GetSlice(opts *bind.CallOpts, _str string) (VidarUtilslice, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getSlice", _str)

	if err != nil {
		return *new(VidarUtilslice), err
	}

	out0 := *abi.ConvertType(out[0], new(VidarUtilslice)).(*VidarUtilslice)

	return out0, err

}

// GetSlice is a free data retrieval call binding the contract method 0x908c0eea.
//
// Solidity: function getSlice(string _str) pure returns((uint256,uint256))
func (_Api *ApiSession) GetSlice(_str string) (VidarUtilslice, error) {
	return _Api.Contract.GetSlice(&_Api.CallOpts, _str)
}

// GetSlice is a free data retrieval call binding the contract method 0x908c0eea.
//
// Solidity: function getSlice(string _str) pure returns((uint256,uint256))
func (_Api *ApiCallerSession) GetSlice(_str string) (VidarUtilslice, error) {
	return _Api.Contract.GetSlice(&_Api.CallOpts, _str)
}

// IsValidAVAXAddress is a free data retrieval call binding the contract method 0xa3fa67ed.
//
// Solidity: function isValidAVAXAddress(address _address) view returns(bool)
func (_Api *ApiCaller) IsValidAVAXAddress(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "isValidAVAXAddress", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidAVAXAddress is a free data retrieval call binding the contract method 0xa3fa67ed.
//
// Solidity: function isValidAVAXAddress(address _address) view returns(bool)
func (_Api *ApiSession) IsValidAVAXAddress(_address common.Address) (bool, error) {
	return _Api.Contract.IsValidAVAXAddress(&_Api.CallOpts, _address)
}

// IsValidAVAXAddress is a free data retrieval call binding the contract method 0xa3fa67ed.
//
// Solidity: function isValidAVAXAddress(address _address) view returns(bool)
func (_Api *ApiCallerSession) IsValidAVAXAddress(_address common.Address) (bool, error) {
	return _Api.Contract.IsValidAVAXAddress(&_Api.CallOpts, _address)
}

// ToHexString is a free data retrieval call binding the contract method 0x63e1cbea.
//
// Solidity: function toHexString(uint256 value, uint256 length) pure returns(string)
func (_Api *ApiCaller) ToHexString(opts *bind.CallOpts, value *big.Int, length *big.Int) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "toHexString", value, length)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToHexString is a free data retrieval call binding the contract method 0x63e1cbea.
//
// Solidity: function toHexString(uint256 value, uint256 length) pure returns(string)
func (_Api *ApiSession) ToHexString(value *big.Int, length *big.Int) (string, error) {
	return _Api.Contract.ToHexString(&_Api.CallOpts, value, length)
}

// ToHexString is a free data retrieval call binding the contract method 0x63e1cbea.
//
// Solidity: function toHexString(uint256 value, uint256 length) pure returns(string)
func (_Api *ApiCallerSession) ToHexString(value *big.Int, length *big.Int) (string, error) {
	return _Api.Contract.ToHexString(&_Api.CallOpts, value, length)
}
