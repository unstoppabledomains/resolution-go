// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package resolver

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"inputs\":[{\"internalType\":\"contractRegistry\",\"name\":\"registry\",\"type\":\"address\"},{\"internalType\":\"contractMintingController\",\"name\":\"mintingController\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"keyIndex\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"NewKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ResetRecords\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"keyIndex\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"valueIndex\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"Set\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keyHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getByHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getMany\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"keyHashes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getManyByHash\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keyHash\",\"type\":\"uint256\"}],\"name\":\"hashToKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"hashes\",\"type\":\"uint256[]\"}],\"name\":\"hashesToKeys\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"nonceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"preconfigure\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"reconfigure\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"reconfigureFor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"reset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"resetFor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"setFor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setMany\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"setManyFor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x1be5e7ed.
//
// Solidity: function get(string key, uint256 tokenId) view returns(string)
func (_Contract *ContractCaller) Get(opts *bind.CallOpts, key string, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "get", key, tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x1be5e7ed.
//
// Solidity: function get(string key, uint256 tokenId) view returns(string)
func (_Contract *ContractSession) Get(key string, tokenId *big.Int) (string, error) {
	return _Contract.Contract.Get(&_Contract.CallOpts, key, tokenId)
}

// Get is a free data retrieval call binding the contract method 0x1be5e7ed.
//
// Solidity: function get(string key, uint256 tokenId) view returns(string)
func (_Contract *ContractCallerSession) Get(key string, tokenId *big.Int) (string, error) {
	return _Contract.Contract.Get(&_Contract.CallOpts, key, tokenId)
}

// GetByHash is a free data retrieval call binding the contract method 0x672b9f81.
//
// Solidity: function getByHash(uint256 keyHash, uint256 tokenId) view returns(string key, string value)
func (_Contract *ContractCaller) GetByHash(opts *bind.CallOpts, keyHash *big.Int, tokenId *big.Int) (struct {
	Key   string
	Value string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getByHash", keyHash, tokenId)

	outstruct := new(struct {
		Key   string
		Value string
	})

	outstruct.Key = out[0].(string)
	outstruct.Value = out[1].(string)

	return *outstruct, err

}

// GetByHash is a free data retrieval call binding the contract method 0x672b9f81.
//
// Solidity: function getByHash(uint256 keyHash, uint256 tokenId) view returns(string key, string value)
func (_Contract *ContractSession) GetByHash(keyHash *big.Int, tokenId *big.Int) (struct {
	Key   string
	Value string
}, error) {
	return _Contract.Contract.GetByHash(&_Contract.CallOpts, keyHash, tokenId)
}

// GetByHash is a free data retrieval call binding the contract method 0x672b9f81.
//
// Solidity: function getByHash(uint256 keyHash, uint256 tokenId) view returns(string key, string value)
func (_Contract *ContractCallerSession) GetByHash(keyHash *big.Int, tokenId *big.Int) (struct {
	Key   string
	Value string
}, error) {
	return _Contract.Contract.GetByHash(&_Contract.CallOpts, keyHash, tokenId)
}

// GetMany is a free data retrieval call binding the contract method 0x1bd8cc1a.
//
// Solidity: function getMany(string[] keys, uint256 tokenId) view returns(string[])
func (_Contract *ContractCaller) GetMany(opts *bind.CallOpts, keys []string, tokenId *big.Int) ([]string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getMany", keys, tokenId)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetMany is a free data retrieval call binding the contract method 0x1bd8cc1a.
//
// Solidity: function getMany(string[] keys, uint256 tokenId) view returns(string[])
func (_Contract *ContractSession) GetMany(keys []string, tokenId *big.Int) ([]string, error) {
	return _Contract.Contract.GetMany(&_Contract.CallOpts, keys, tokenId)
}

// GetMany is a free data retrieval call binding the contract method 0x1bd8cc1a.
//
// Solidity: function getMany(string[] keys, uint256 tokenId) view returns(string[])
func (_Contract *ContractCallerSession) GetMany(keys []string, tokenId *big.Int) ([]string, error) {
	return _Contract.Contract.GetMany(&_Contract.CallOpts, keys, tokenId)
}

// GetManyByHash is a free data retrieval call binding the contract method 0xb85afd28.
//
// Solidity: function getManyByHash(uint256[] keyHashes, uint256 tokenId) view returns(string[] keys, string[] values)
func (_Contract *ContractCaller) GetManyByHash(opts *bind.CallOpts, keyHashes []*big.Int, tokenId *big.Int) (struct {
	Keys   []string
	Values []string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getManyByHash", keyHashes, tokenId)

	outstruct := new(struct {
		Keys   []string
		Values []string
	})

	outstruct.Keys = out[0].([]string)
	outstruct.Values = out[1].([]string)

	return *outstruct, err

}

// GetManyByHash is a free data retrieval call binding the contract method 0xb85afd28.
//
// Solidity: function getManyByHash(uint256[] keyHashes, uint256 tokenId) view returns(string[] keys, string[] values)
func (_Contract *ContractSession) GetManyByHash(keyHashes []*big.Int, tokenId *big.Int) (struct {
	Keys   []string
	Values []string
}, error) {
	return _Contract.Contract.GetManyByHash(&_Contract.CallOpts, keyHashes, tokenId)
}

// GetManyByHash is a free data retrieval call binding the contract method 0xb85afd28.
//
// Solidity: function getManyByHash(uint256[] keyHashes, uint256 tokenId) view returns(string[] keys, string[] values)
func (_Contract *ContractCallerSession) GetManyByHash(keyHashes []*big.Int, tokenId *big.Int) (struct {
	Keys   []string
	Values []string
}, error) {
	return _Contract.Contract.GetManyByHash(&_Contract.CallOpts, keyHashes, tokenId)
}

// HashToKey is a free data retrieval call binding the contract method 0x2c3d376d.
//
// Solidity: function hashToKey(uint256 keyHash) view returns(string)
func (_Contract *ContractCaller) HashToKey(opts *bind.CallOpts, keyHash *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "hashToKey", keyHash)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// HashToKey is a free data retrieval call binding the contract method 0x2c3d376d.
//
// Solidity: function hashToKey(uint256 keyHash) view returns(string)
func (_Contract *ContractSession) HashToKey(keyHash *big.Int) (string, error) {
	return _Contract.Contract.HashToKey(&_Contract.CallOpts, keyHash)
}

// HashToKey is a free data retrieval call binding the contract method 0x2c3d376d.
//
// Solidity: function hashToKey(uint256 keyHash) view returns(string)
func (_Contract *ContractCallerSession) HashToKey(keyHash *big.Int) (string, error) {
	return _Contract.Contract.HashToKey(&_Contract.CallOpts, keyHash)
}

// HashesToKeys is a free data retrieval call binding the contract method 0x8c87a4ea.
//
// Solidity: function hashesToKeys(uint256[] hashes) view returns(string[])
func (_Contract *ContractCaller) HashesToKeys(opts *bind.CallOpts, hashes []*big.Int) ([]string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "hashesToKeys", hashes)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// HashesToKeys is a free data retrieval call binding the contract method 0x8c87a4ea.
//
// Solidity: function hashesToKeys(uint256[] hashes) view returns(string[])
func (_Contract *ContractSession) HashesToKeys(hashes []*big.Int) ([]string, error) {
	return _Contract.Contract.HashesToKeys(&_Contract.CallOpts, hashes)
}

// HashesToKeys is a free data retrieval call binding the contract method 0x8c87a4ea.
//
// Solidity: function hashesToKeys(uint256[] hashes) view returns(string[])
func (_Contract *ContractCallerSession) HashesToKeys(hashes []*big.Int) ([]string, error) {
	return _Contract.Contract.HashesToKeys(&_Contract.CallOpts, hashes)
}

// NonceOf is a free data retrieval call binding the contract method 0x6ccbae5f.
//
// Solidity: function nonceOf(uint256 tokenId) view returns(uint256)
func (_Contract *ContractCaller) NonceOf(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nonceOf", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonceOf is a free data retrieval call binding the contract method 0x6ccbae5f.
//
// Solidity: function nonceOf(uint256 tokenId) view returns(uint256)
func (_Contract *ContractSession) NonceOf(tokenId *big.Int) (*big.Int, error) {
	return _Contract.Contract.NonceOf(&_Contract.CallOpts, tokenId)
}

// NonceOf is a free data retrieval call binding the contract method 0x6ccbae5f.
//
// Solidity: function nonceOf(uint256 tokenId) view returns(uint256)
func (_Contract *ContractCallerSession) NonceOf(tokenId *big.Int) (*big.Int, error) {
	return _Contract.Contract.NonceOf(&_Contract.CallOpts, tokenId)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Contract *ContractCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Contract *ContractSession) Registry() (common.Address, error) {
	return _Contract.Contract.Registry(&_Contract.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Contract *ContractCallerSession) Registry() (common.Address, error) {
	return _Contract.Contract.Registry(&_Contract.CallOpts)
}

// Preconfigure is a paid mutator transaction binding the contract method 0xe837ae74.
//
// Solidity: function preconfigure(string[] keys, string[] values, uint256 tokenId) returns()
func (_Contract *ContractTransactor) Preconfigure(opts *bind.TransactOpts, keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "preconfigure", keys, values, tokenId)
}

// Preconfigure is a paid mutator transaction binding the contract method 0xe837ae74.
//
// Solidity: function preconfigure(string[] keys, string[] values, uint256 tokenId) returns()
func (_Contract *ContractSession) Preconfigure(keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Preconfigure(&_Contract.TransactOpts, keys, values, tokenId)
}

// Preconfigure is a paid mutator transaction binding the contract method 0xe837ae74.
//
// Solidity: function preconfigure(string[] keys, string[] values, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) Preconfigure(keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Preconfigure(&_Contract.TransactOpts, keys, values, tokenId)
}

// Reconfigure is a paid mutator transaction binding the contract method 0xec129eea.
//
// Solidity: function reconfigure(string[] keys, string[] values, uint256 tokenId) returns()
func (_Contract *ContractTransactor) Reconfigure(opts *bind.TransactOpts, keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "reconfigure", keys, values, tokenId)
}

// Reconfigure is a paid mutator transaction binding the contract method 0xec129eea.
//
// Solidity: function reconfigure(string[] keys, string[] values, uint256 tokenId) returns()
func (_Contract *ContractSession) Reconfigure(keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Reconfigure(&_Contract.TransactOpts, keys, values, tokenId)
}

// Reconfigure is a paid mutator transaction binding the contract method 0xec129eea.
//
// Solidity: function reconfigure(string[] keys, string[] values, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) Reconfigure(keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Reconfigure(&_Contract.TransactOpts, keys, values, tokenId)
}

// ReconfigureFor is a paid mutator transaction binding the contract method 0xa3557e6c.
//
// Solidity: function reconfigureFor(string[] keys, string[] values, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactor) ReconfigureFor(opts *bind.TransactOpts, keys []string, values []string, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "reconfigureFor", keys, values, tokenId, signature)
}

// ReconfigureFor is a paid mutator transaction binding the contract method 0xa3557e6c.
//
// Solidity: function reconfigureFor(string[] keys, string[] values, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractSession) ReconfigureFor(keys []string, values []string, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.ReconfigureFor(&_Contract.TransactOpts, keys, values, tokenId, signature)
}

// ReconfigureFor is a paid mutator transaction binding the contract method 0xa3557e6c.
//
// Solidity: function reconfigureFor(string[] keys, string[] values, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactorSession) ReconfigureFor(keys []string, values []string, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.ReconfigureFor(&_Contract.TransactOpts, keys, values, tokenId, signature)
}

// Reset is a paid mutator transaction binding the contract method 0x310bd74b.
//
// Solidity: function reset(uint256 tokenId) returns()
func (_Contract *ContractTransactor) Reset(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "reset", tokenId)
}

// Reset is a paid mutator transaction binding the contract method 0x310bd74b.
//
// Solidity: function reset(uint256 tokenId) returns()
func (_Contract *ContractSession) Reset(tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Reset(&_Contract.TransactOpts, tokenId)
}

// Reset is a paid mutator transaction binding the contract method 0x310bd74b.
//
// Solidity: function reset(uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) Reset(tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Reset(&_Contract.TransactOpts, tokenId)
}

// ResetFor is a paid mutator transaction binding the contract method 0xb87abc11.
//
// Solidity: function resetFor(uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactor) ResetFor(opts *bind.TransactOpts, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "resetFor", tokenId, signature)
}

// ResetFor is a paid mutator transaction binding the contract method 0xb87abc11.
//
// Solidity: function resetFor(uint256 tokenId, bytes signature) returns()
func (_Contract *ContractSession) ResetFor(tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.ResetFor(&_Contract.TransactOpts, tokenId, signature)
}

// ResetFor is a paid mutator transaction binding the contract method 0xb87abc11.
//
// Solidity: function resetFor(uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactorSession) ResetFor(tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.ResetFor(&_Contract.TransactOpts, tokenId, signature)
}

// Set is a paid mutator transaction binding the contract method 0x47c81699.
//
// Solidity: function set(string key, string value, uint256 tokenId) returns()
func (_Contract *ContractTransactor) Set(opts *bind.TransactOpts, key string, value string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "set", key, value, tokenId)
}

// Set is a paid mutator transaction binding the contract method 0x47c81699.
//
// Solidity: function set(string key, string value, uint256 tokenId) returns()
func (_Contract *ContractSession) Set(key string, value string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Set(&_Contract.TransactOpts, key, value, tokenId)
}

// Set is a paid mutator transaction binding the contract method 0x47c81699.
//
// Solidity: function set(string key, string value, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) Set(key string, value string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Set(&_Contract.TransactOpts, key, value, tokenId)
}

// SetFor is a paid mutator transaction binding the contract method 0xc5974073.
//
// Solidity: function setFor(string key, string value, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactor) SetFor(opts *bind.TransactOpts, key string, value string, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setFor", key, value, tokenId, signature)
}

// SetFor is a paid mutator transaction binding the contract method 0xc5974073.
//
// Solidity: function setFor(string key, string value, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractSession) SetFor(key string, value string, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetFor(&_Contract.TransactOpts, key, value, tokenId, signature)
}

// SetFor is a paid mutator transaction binding the contract method 0xc5974073.
//
// Solidity: function setFor(string key, string value, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactorSession) SetFor(key string, value string, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetFor(&_Contract.TransactOpts, key, value, tokenId, signature)
}

// SetMany is a paid mutator transaction binding the contract method 0xce92b33e.
//
// Solidity: function setMany(string[] keys, string[] values, uint256 tokenId) returns()
func (_Contract *ContractTransactor) SetMany(opts *bind.TransactOpts, keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setMany", keys, values, tokenId)
}

// SetMany is a paid mutator transaction binding the contract method 0xce92b33e.
//
// Solidity: function setMany(string[] keys, string[] values, uint256 tokenId) returns()
func (_Contract *ContractSession) SetMany(keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMany(&_Contract.TransactOpts, keys, values, tokenId)
}

// SetMany is a paid mutator transaction binding the contract method 0xce92b33e.
//
// Solidity: function setMany(string[] keys, string[] values, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) SetMany(keys []string, values []string, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMany(&_Contract.TransactOpts, keys, values, tokenId)
}

// SetManyFor is a paid mutator transaction binding the contract method 0x8f69c188.
//
// Solidity: function setManyFor(string[] keys, string[] values, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactor) SetManyFor(opts *bind.TransactOpts, keys []string, values []string, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setManyFor", keys, values, tokenId, signature)
}

// SetManyFor is a paid mutator transaction binding the contract method 0x8f69c188.
//
// Solidity: function setManyFor(string[] keys, string[] values, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractSession) SetManyFor(keys []string, values []string, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetManyFor(&_Contract.TransactOpts, keys, values, tokenId, signature)
}

// SetManyFor is a paid mutator transaction binding the contract method 0x8f69c188.
//
// Solidity: function setManyFor(string[] keys, string[] values, uint256 tokenId, bytes signature) returns()
func (_Contract *ContractTransactorSession) SetManyFor(keys []string, values []string, tokenId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetManyFor(&_Contract.TransactOpts, keys, values, tokenId, signature)
}

// ContractNewKeyIterator is returned from FilterNewKey and is used to iterate over the raw logs and unpacked data for NewKey events raised by the Contract contract.
type ContractNewKeyIterator struct {
	Event *ContractNewKey // Event containing the contract specifics and raw log

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
func (it *ContractNewKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewKey)
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
		it.Event = new(ContractNewKey)
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
func (it *ContractNewKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewKey represents a NewKey event raised by the Contract contract.
type ContractNewKey struct {
	TokenId  *big.Int
	KeyIndex common.Hash
	Key      string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewKey is a free log retrieval operation binding the contract event 0x7ae4f661958fbecc2f77be6b0eb280d2a6f604b29e1e7221c82b9da0c4af7f86.
//
// Solidity: event NewKey(uint256 indexed tokenId, string indexed keyIndex, string key)
func (_Contract *ContractFilterer) FilterNewKey(opts *bind.FilterOpts, tokenId []*big.Int, keyIndex []string) (*ContractNewKeyIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyIndexRule []interface{}
	for _, keyIndexItem := range keyIndex {
		keyIndexRule = append(keyIndexRule, keyIndexItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewKey", tokenIdRule, keyIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractNewKeyIterator{contract: _Contract.contract, event: "NewKey", logs: logs, sub: sub}, nil
}

// WatchNewKey is a free log subscription operation binding the contract event 0x7ae4f661958fbecc2f77be6b0eb280d2a6f604b29e1e7221c82b9da0c4af7f86.
//
// Solidity: event NewKey(uint256 indexed tokenId, string indexed keyIndex, string key)
func (_Contract *ContractFilterer) WatchNewKey(opts *bind.WatchOpts, sink chan<- *ContractNewKey, tokenId []*big.Int, keyIndex []string) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyIndexRule []interface{}
	for _, keyIndexItem := range keyIndex {
		keyIndexRule = append(keyIndexRule, keyIndexItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewKey", tokenIdRule, keyIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewKey)
				if err := _Contract.contract.UnpackLog(event, "NewKey", log); err != nil {
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

// ParseNewKey is a log parse operation binding the contract event 0x7ae4f661958fbecc2f77be6b0eb280d2a6f604b29e1e7221c82b9da0c4af7f86.
//
// Solidity: event NewKey(uint256 indexed tokenId, string indexed keyIndex, string key)
func (_Contract *ContractFilterer) ParseNewKey(log types.Log) (*ContractNewKey, error) {
	event := new(ContractNewKey)
	if err := _Contract.contract.UnpackLog(event, "NewKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractResetRecordsIterator is returned from FilterResetRecords and is used to iterate over the raw logs and unpacked data for ResetRecords events raised by the Contract contract.
type ContractResetRecordsIterator struct {
	Event *ContractResetRecords // Event containing the contract specifics and raw log

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
func (it *ContractResetRecordsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractResetRecords)
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
		it.Event = new(ContractResetRecords)
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
func (it *ContractResetRecordsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractResetRecordsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractResetRecords represents a ResetRecords event raised by the Contract contract.
type ContractResetRecords struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterResetRecords is a free log retrieval operation binding the contract event 0x185c30856dadb58bf097c1f665a52ada7029752dbcad008ea3fefc73bee8c9fe.
//
// Solidity: event ResetRecords(uint256 indexed tokenId)
func (_Contract *ContractFilterer) FilterResetRecords(opts *bind.FilterOpts, tokenId []*big.Int) (*ContractResetRecordsIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ResetRecords", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractResetRecordsIterator{contract: _Contract.contract, event: "ResetRecords", logs: logs, sub: sub}, nil
}

// WatchResetRecords is a free log subscription operation binding the contract event 0x185c30856dadb58bf097c1f665a52ada7029752dbcad008ea3fefc73bee8c9fe.
//
// Solidity: event ResetRecords(uint256 indexed tokenId)
func (_Contract *ContractFilterer) WatchResetRecords(opts *bind.WatchOpts, sink chan<- *ContractResetRecords, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ResetRecords", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractResetRecords)
				if err := _Contract.contract.UnpackLog(event, "ResetRecords", log); err != nil {
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

// ParseResetRecords is a log parse operation binding the contract event 0x185c30856dadb58bf097c1f665a52ada7029752dbcad008ea3fefc73bee8c9fe.
//
// Solidity: event ResetRecords(uint256 indexed tokenId)
func (_Contract *ContractFilterer) ParseResetRecords(log types.Log) (*ContractResetRecords, error) {
	event := new(ContractResetRecords)
	if err := _Contract.contract.UnpackLog(event, "ResetRecords", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the Contract contract.
type ContractSetIterator struct {
	Event *ContractSet // Event containing the contract specifics and raw log

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
func (it *ContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSet)
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
		it.Event = new(ContractSet)
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
func (it *ContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSet represents a Set event raised by the Contract contract.
type ContractSet struct {
	TokenId    *big.Int
	KeyIndex   common.Hash
	ValueIndex common.Hash
	Key        string
	Value      string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0x851ffe8e74d5015261dba0a1f9e1b0e5d42c5af5d2ad1908fee897c7d80a0d92.
//
// Solidity: event Set(uint256 indexed tokenId, string indexed keyIndex, string indexed valueIndex, string key, string value)
func (_Contract *ContractFilterer) FilterSet(opts *bind.FilterOpts, tokenId []*big.Int, keyIndex []string, valueIndex []string) (*ContractSetIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyIndexRule []interface{}
	for _, keyIndexItem := range keyIndex {
		keyIndexRule = append(keyIndexRule, keyIndexItem)
	}
	var valueIndexRule []interface{}
	for _, valueIndexItem := range valueIndex {
		valueIndexRule = append(valueIndexRule, valueIndexItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Set", tokenIdRule, keyIndexRule, valueIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractSetIterator{contract: _Contract.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0x851ffe8e74d5015261dba0a1f9e1b0e5d42c5af5d2ad1908fee897c7d80a0d92.
//
// Solidity: event Set(uint256 indexed tokenId, string indexed keyIndex, string indexed valueIndex, string key, string value)
func (_Contract *ContractFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *ContractSet, tokenId []*big.Int, keyIndex []string, valueIndex []string) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var keyIndexRule []interface{}
	for _, keyIndexItem := range keyIndex {
		keyIndexRule = append(keyIndexRule, keyIndexItem)
	}
	var valueIndexRule []interface{}
	for _, valueIndexItem := range valueIndex {
		valueIndexRule = append(valueIndexRule, valueIndexItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Set", tokenIdRule, keyIndexRule, valueIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSet)
				if err := _Contract.contract.UnpackLog(event, "Set", log); err != nil {
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

// ParseSet is a log parse operation binding the contract event 0x851ffe8e74d5015261dba0a1f9e1b0e5d42c5af5d2ad1908fee897c7d80a0d92.
//
// Solidity: event Set(uint256 indexed tokenId, string indexed keyIndex, string indexed valueIndex, string key, string value)
func (_Contract *ContractFilterer) ParseSet(log types.Log) (*ContractSet, error) {
	event := new(ContractSet)
	if err := _Contract.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
