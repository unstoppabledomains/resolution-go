// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxyreader

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenKey\",\"type\":\"string\"}],\"name\":\"SetLegacyRecords\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"}],\"name\":\"SetNetworkFamily\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"networks\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"families\",\"type\":\"string[]\"}],\"name\":\"addBlockchainNetworks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"networks\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"family\",\"type\":\"string\"}],\"name\":\"addBlockchainNetworks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[][]\",\"name\":\"legacyKeys\",\"type\":\"string[][]\"}],\"name\":\"addLegacyRecords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"token\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"token\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getAddressKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"token\",\"type\":\"string\"}],\"name\":\"getAddressKeys\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keyHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getByHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"keyHashes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getDataByHash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"keyHashes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"getDataByHashForMany\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"resolvers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"keys\",\"type\":\"string[][]\"},{\"internalType\":\"string[][]\",\"name\":\"values\",\"type\":\"string[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"getDataForMany\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"resolvers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"string[][]\",\"name\":\"values\",\"type\":\"string[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getMany\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"keyHashes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getManyByHash\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUNSRegistry\",\"name\":\"unsRegistry\",\"type\":\"address\"},{\"internalType\":\"contractICNSRegistry\",\"name\":\"cnsRegistry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isApprovedOrOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"labels\",\"type\":\"string[]\"}],\"name\":\"namehash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"ownerOfForMany\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"registryOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"resolverOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"reverseNameOf\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"reverseOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

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
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Contract *ContractCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Contract *ContractSession) NAME() (string, error) {
	return _Contract.Contract.NAME(&_Contract.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Contract *ContractCallerSession) NAME() (string, error) {
	return _Contract.Contract.NAME(&_Contract.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Contract *ContractCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Contract *ContractSession) VERSION() (string, error) {
	return _Contract.Contract.VERSION(&_Contract.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Contract *ContractCallerSession) VERSION() (string, error) {
	return _Contract.Contract.VERSION(&_Contract.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contract *ContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contract *ContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contract *ContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Contract *ContractCaller) Exists(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "exists", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Contract *ContractSession) Exists(tokenId *big.Int) (bool, error) {
	return _Contract.Contract.Exists(&_Contract.CallOpts, tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Contract *ContractCallerSession) Exists(tokenId *big.Int) (bool, error) {
	return _Contract.Contract.Exists(&_Contract.CallOpts, tokenId)
}

// Get is a free data retrieval call binding the contract method 0x1be5e7ed.
//
// Solidity: function get(string key, uint256 tokenId) view returns(string value)
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
// Solidity: function get(string key, uint256 tokenId) view returns(string value)
func (_Contract *ContractSession) Get(key string, tokenId *big.Int) (string, error) {
	return _Contract.Contract.Get(&_Contract.CallOpts, key, tokenId)
}

// Get is a free data retrieval call binding the contract method 0x1be5e7ed.
//
// Solidity: function get(string key, uint256 tokenId) view returns(string value)
func (_Contract *ContractCallerSession) Get(key string, tokenId *big.Int) (string, error) {
	return _Contract.Contract.Get(&_Contract.CallOpts, key, tokenId)
}

// GetAddress is a free data retrieval call binding the contract method 0xbfc5429b.
//
// Solidity: function getAddress(string network, string token, uint256 tokenId) view returns(string addr)
func (_Contract *ContractCaller) GetAddress(opts *bind.CallOpts, network string, token string, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAddress", network, token, tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0xbfc5429b.
//
// Solidity: function getAddress(string network, string token, uint256 tokenId) view returns(string addr)
func (_Contract *ContractSession) GetAddress(network string, token string, tokenId *big.Int) (string, error) {
	return _Contract.Contract.GetAddress(&_Contract.CallOpts, network, token, tokenId)
}

// GetAddress is a free data retrieval call binding the contract method 0xbfc5429b.
//
// Solidity: function getAddress(string network, string token, uint256 tokenId) view returns(string addr)
func (_Contract *ContractCallerSession) GetAddress(network string, token string, tokenId *big.Int) (string, error) {
	return _Contract.Contract.GetAddress(&_Contract.CallOpts, network, token, tokenId)
}

// GetAddressKey is a free data retrieval call binding the contract method 0xb98b3526.
//
// Solidity: function getAddressKey(string network, string token, uint256 tokenId) view returns(string key)
func (_Contract *ContractCaller) GetAddressKey(opts *bind.CallOpts, network string, token string, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAddressKey", network, token, tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAddressKey is a free data retrieval call binding the contract method 0xb98b3526.
//
// Solidity: function getAddressKey(string network, string token, uint256 tokenId) view returns(string key)
func (_Contract *ContractSession) GetAddressKey(network string, token string, tokenId *big.Int) (string, error) {
	return _Contract.Contract.GetAddressKey(&_Contract.CallOpts, network, token, tokenId)
}

// GetAddressKey is a free data retrieval call binding the contract method 0xb98b3526.
//
// Solidity: function getAddressKey(string network, string token, uint256 tokenId) view returns(string key)
func (_Contract *ContractCallerSession) GetAddressKey(network string, token string, tokenId *big.Int) (string, error) {
	return _Contract.Contract.GetAddressKey(&_Contract.CallOpts, network, token, tokenId)
}

// GetAddressKeys is a free data retrieval call binding the contract method 0x83caffbe.
//
// Solidity: function getAddressKeys(string network, string token) view returns(string[] keys)
func (_Contract *ContractCaller) GetAddressKeys(opts *bind.CallOpts, network string, token string) ([]string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAddressKeys", network, token)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAddressKeys is a free data retrieval call binding the contract method 0x83caffbe.
//
// Solidity: function getAddressKeys(string network, string token) view returns(string[] keys)
func (_Contract *ContractSession) GetAddressKeys(network string, token string) ([]string, error) {
	return _Contract.Contract.GetAddressKeys(&_Contract.CallOpts, network, token)
}

// GetAddressKeys is a free data retrieval call binding the contract method 0x83caffbe.
//
// Solidity: function getAddressKeys(string network, string token) view returns(string[] keys)
func (_Contract *ContractCallerSession) GetAddressKeys(network string, token string) ([]string, error) {
	return _Contract.Contract.GetAddressKeys(&_Contract.CallOpts, network, token)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contract *ContractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetApproved(&_Contract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetApproved(&_Contract.CallOpts, tokenId)
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
	if err != nil {
		return *outstruct, err
	}

	outstruct.Key = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Value = *abi.ConvertType(out[1], new(string)).(*string)

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

// GetData is a free data retrieval call binding the contract method 0x91015f6b.
//
// Solidity: function getData(string[] keys, uint256 tokenId) view returns(address resolver, address owner, string[] values)
func (_Contract *ContractCaller) GetData(opts *bind.CallOpts, keys []string, tokenId *big.Int) (struct {
	Resolver common.Address
	Owner    common.Address
	Values   []string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getData", keys, tokenId)

	outstruct := new(struct {
		Resolver common.Address
		Owner    common.Address
		Values   []string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Resolver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Values = *abi.ConvertType(out[2], new([]string)).(*[]string)

	return *outstruct, err

}

// GetData is a free data retrieval call binding the contract method 0x91015f6b.
//
// Solidity: function getData(string[] keys, uint256 tokenId) view returns(address resolver, address owner, string[] values)
func (_Contract *ContractSession) GetData(keys []string, tokenId *big.Int) (struct {
	Resolver common.Address
	Owner    common.Address
	Values   []string
}, error) {
	return _Contract.Contract.GetData(&_Contract.CallOpts, keys, tokenId)
}

// GetData is a free data retrieval call binding the contract method 0x91015f6b.
//
// Solidity: function getData(string[] keys, uint256 tokenId) view returns(address resolver, address owner, string[] values)
func (_Contract *ContractCallerSession) GetData(keys []string, tokenId *big.Int) (struct {
	Resolver common.Address
	Owner    common.Address
	Values   []string
}, error) {
	return _Contract.Contract.GetData(&_Contract.CallOpts, keys, tokenId)
}

// GetDataByHash is a free data retrieval call binding the contract method 0x03280755.
//
// Solidity: function getDataByHash(uint256[] keyHashes, uint256 tokenId) view returns(address resolver, address owner, string[] keys, string[] values)
func (_Contract *ContractCaller) GetDataByHash(opts *bind.CallOpts, keyHashes []*big.Int, tokenId *big.Int) (struct {
	Resolver common.Address
	Owner    common.Address
	Keys     []string
	Values   []string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDataByHash", keyHashes, tokenId)

	outstruct := new(struct {
		Resolver common.Address
		Owner    common.Address
		Keys     []string
		Values   []string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Resolver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Keys = *abi.ConvertType(out[2], new([]string)).(*[]string)
	outstruct.Values = *abi.ConvertType(out[3], new([]string)).(*[]string)

	return *outstruct, err

}

// GetDataByHash is a free data retrieval call binding the contract method 0x03280755.
//
// Solidity: function getDataByHash(uint256[] keyHashes, uint256 tokenId) view returns(address resolver, address owner, string[] keys, string[] values)
func (_Contract *ContractSession) GetDataByHash(keyHashes []*big.Int, tokenId *big.Int) (struct {
	Resolver common.Address
	Owner    common.Address
	Keys     []string
	Values   []string
}, error) {
	return _Contract.Contract.GetDataByHash(&_Contract.CallOpts, keyHashes, tokenId)
}

// GetDataByHash is a free data retrieval call binding the contract method 0x03280755.
//
// Solidity: function getDataByHash(uint256[] keyHashes, uint256 tokenId) view returns(address resolver, address owner, string[] keys, string[] values)
func (_Contract *ContractCallerSession) GetDataByHash(keyHashes []*big.Int, tokenId *big.Int) (struct {
	Resolver common.Address
	Owner    common.Address
	Keys     []string
	Values   []string
}, error) {
	return _Contract.Contract.GetDataByHash(&_Contract.CallOpts, keyHashes, tokenId)
}

// GetDataByHashForMany is a free data retrieval call binding the contract method 0x869b8884.
//
// Solidity: function getDataByHashForMany(uint256[] keyHashes, uint256[] tokenIds) view returns(address[] resolvers, address[] owners, string[][] keys, string[][] values)
func (_Contract *ContractCaller) GetDataByHashForMany(opts *bind.CallOpts, keyHashes []*big.Int, tokenIds []*big.Int) (struct {
	Resolvers []common.Address
	Owners    []common.Address
	Keys      [][]string
	Values    [][]string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDataByHashForMany", keyHashes, tokenIds)

	outstruct := new(struct {
		Resolvers []common.Address
		Owners    []common.Address
		Keys      [][]string
		Values    [][]string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Resolvers = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Owners = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Keys = *abi.ConvertType(out[2], new([][]string)).(*[][]string)
	outstruct.Values = *abi.ConvertType(out[3], new([][]string)).(*[][]string)

	return *outstruct, err

}

// GetDataByHashForMany is a free data retrieval call binding the contract method 0x869b8884.
//
// Solidity: function getDataByHashForMany(uint256[] keyHashes, uint256[] tokenIds) view returns(address[] resolvers, address[] owners, string[][] keys, string[][] values)
func (_Contract *ContractSession) GetDataByHashForMany(keyHashes []*big.Int, tokenIds []*big.Int) (struct {
	Resolvers []common.Address
	Owners    []common.Address
	Keys      [][]string
	Values    [][]string
}, error) {
	return _Contract.Contract.GetDataByHashForMany(&_Contract.CallOpts, keyHashes, tokenIds)
}

// GetDataByHashForMany is a free data retrieval call binding the contract method 0x869b8884.
//
// Solidity: function getDataByHashForMany(uint256[] keyHashes, uint256[] tokenIds) view returns(address[] resolvers, address[] owners, string[][] keys, string[][] values)
func (_Contract *ContractCallerSession) GetDataByHashForMany(keyHashes []*big.Int, tokenIds []*big.Int) (struct {
	Resolvers []common.Address
	Owners    []common.Address
	Keys      [][]string
	Values    [][]string
}, error) {
	return _Contract.Contract.GetDataByHashForMany(&_Contract.CallOpts, keyHashes, tokenIds)
}

// GetDataForMany is a free data retrieval call binding the contract method 0x933c051d.
//
// Solidity: function getDataForMany(string[] keys, uint256[] tokenIds) view returns(address[] resolvers, address[] owners, string[][] values)
func (_Contract *ContractCaller) GetDataForMany(opts *bind.CallOpts, keys []string, tokenIds []*big.Int) (struct {
	Resolvers []common.Address
	Owners    []common.Address
	Values    [][]string
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDataForMany", keys, tokenIds)

	outstruct := new(struct {
		Resolvers []common.Address
		Owners    []common.Address
		Values    [][]string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Resolvers = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Owners = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Values = *abi.ConvertType(out[2], new([][]string)).(*[][]string)

	return *outstruct, err

}

// GetDataForMany is a free data retrieval call binding the contract method 0x933c051d.
//
// Solidity: function getDataForMany(string[] keys, uint256[] tokenIds) view returns(address[] resolvers, address[] owners, string[][] values)
func (_Contract *ContractSession) GetDataForMany(keys []string, tokenIds []*big.Int) (struct {
	Resolvers []common.Address
	Owners    []common.Address
	Values    [][]string
}, error) {
	return _Contract.Contract.GetDataForMany(&_Contract.CallOpts, keys, tokenIds)
}

// GetDataForMany is a free data retrieval call binding the contract method 0x933c051d.
//
// Solidity: function getDataForMany(string[] keys, uint256[] tokenIds) view returns(address[] resolvers, address[] owners, string[][] values)
func (_Contract *ContractCallerSession) GetDataForMany(keys []string, tokenIds []*big.Int) (struct {
	Resolvers []common.Address
	Owners    []common.Address
	Values    [][]string
}, error) {
	return _Contract.Contract.GetDataForMany(&_Contract.CallOpts, keys, tokenIds)
}

// GetMany is a free data retrieval call binding the contract method 0x1bd8cc1a.
//
// Solidity: function getMany(string[] keys, uint256 tokenId) view returns(string[] values)
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
// Solidity: function getMany(string[] keys, uint256 tokenId) view returns(string[] values)
func (_Contract *ContractSession) GetMany(keys []string, tokenId *big.Int) ([]string, error) {
	return _Contract.Contract.GetMany(&_Contract.CallOpts, keys, tokenId)
}

// GetMany is a free data retrieval call binding the contract method 0x1bd8cc1a.
//
// Solidity: function getMany(string[] keys, uint256 tokenId) view returns(string[] values)
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
	if err != nil {
		return *outstruct, err
	}

	outstruct.Keys = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Values = *abi.ConvertType(out[1], new([]string)).(*[]string)

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

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) pure returns(bool)
func (_Contract *ContractCaller) IsApprovedForAll(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isApprovedForAll", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) pure returns(bool)
func (_Contract *ContractSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Contract.Contract.IsApprovedForAll(&_Contract.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) pure returns(bool)
func (_Contract *ContractCallerSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Contract.Contract.IsApprovedForAll(&_Contract.CallOpts, arg0, arg1)
}

// IsApprovedOrOwner is a free data retrieval call binding the contract method 0x430c2081.
//
// Solidity: function isApprovedOrOwner(address spender, uint256 tokenId) view returns(bool)
func (_Contract *ContractCaller) IsApprovedOrOwner(opts *bind.CallOpts, spender common.Address, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isApprovedOrOwner", spender, tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedOrOwner is a free data retrieval call binding the contract method 0x430c2081.
//
// Solidity: function isApprovedOrOwner(address spender, uint256 tokenId) view returns(bool)
func (_Contract *ContractSession) IsApprovedOrOwner(spender common.Address, tokenId *big.Int) (bool, error) {
	return _Contract.Contract.IsApprovedOrOwner(&_Contract.CallOpts, spender, tokenId)
}

// IsApprovedOrOwner is a free data retrieval call binding the contract method 0x430c2081.
//
// Solidity: function isApprovedOrOwner(address spender, uint256 tokenId) view returns(bool)
func (_Contract *ContractCallerSession) IsApprovedOrOwner(spender common.Address, tokenId *big.Int) (bool, error) {
	return _Contract.Contract.IsApprovedOrOwner(&_Contract.CallOpts, spender, tokenId)
}

// Namehash is a free data retrieval call binding the contract method 0x276fabb1.
//
// Solidity: function namehash(string[] labels) view returns(uint256)
func (_Contract *ContractCaller) Namehash(opts *bind.CallOpts, labels []string) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "namehash", labels)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Namehash is a free data retrieval call binding the contract method 0x276fabb1.
//
// Solidity: function namehash(string[] labels) view returns(uint256)
func (_Contract *ContractSession) Namehash(labels []string) (*big.Int, error) {
	return _Contract.Contract.Namehash(&_Contract.CallOpts, labels)
}

// Namehash is a free data retrieval call binding the contract method 0x276fabb1.
//
// Solidity: function namehash(string[] labels) view returns(uint256)
func (_Contract *ContractCallerSession) Namehash(labels []string) (*big.Int, error) {
	return _Contract.Contract.Namehash(&_Contract.CallOpts, labels)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contract *ContractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.OwnerOf(&_Contract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.OwnerOf(&_Contract.CallOpts, tokenId)
}

// OwnerOfForMany is a free data retrieval call binding the contract method 0xc15ae7cf.
//
// Solidity: function ownerOfForMany(uint256[] tokenIds) view returns(address[] owners)
func (_Contract *ContractCaller) OwnerOfForMany(opts *bind.CallOpts, tokenIds []*big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ownerOfForMany", tokenIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// OwnerOfForMany is a free data retrieval call binding the contract method 0xc15ae7cf.
//
// Solidity: function ownerOfForMany(uint256[] tokenIds) view returns(address[] owners)
func (_Contract *ContractSession) OwnerOfForMany(tokenIds []*big.Int) ([]common.Address, error) {
	return _Contract.Contract.OwnerOfForMany(&_Contract.CallOpts, tokenIds)
}

// OwnerOfForMany is a free data retrieval call binding the contract method 0xc15ae7cf.
//
// Solidity: function ownerOfForMany(uint256[] tokenIds) view returns(address[] owners)
func (_Contract *ContractCallerSession) OwnerOfForMany(tokenIds []*big.Int) ([]common.Address, error) {
	return _Contract.Contract.OwnerOfForMany(&_Contract.CallOpts, tokenIds)
}

// RegistryOf is a free data retrieval call binding the contract method 0xa81ce6f9.
//
// Solidity: function registryOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) RegistryOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "registryOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryOf is a free data retrieval call binding the contract method 0xa81ce6f9.
//
// Solidity: function registryOf(uint256 tokenId) view returns(address)
func (_Contract *ContractSession) RegistryOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.RegistryOf(&_Contract.CallOpts, tokenId)
}

// RegistryOf is a free data retrieval call binding the contract method 0xa81ce6f9.
//
// Solidity: function registryOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) RegistryOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.RegistryOf(&_Contract.CallOpts, tokenId)
}

// ResolverOf is a free data retrieval call binding the contract method 0xb3f9e4cb.
//
// Solidity: function resolverOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCaller) ResolverOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolverOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResolverOf is a free data retrieval call binding the contract method 0xb3f9e4cb.
//
// Solidity: function resolverOf(uint256 tokenId) view returns(address)
func (_Contract *ContractSession) ResolverOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.ResolverOf(&_Contract.CallOpts, tokenId)
}

// ResolverOf is a free data retrieval call binding the contract method 0xb3f9e4cb.
//
// Solidity: function resolverOf(uint256 tokenId) view returns(address)
func (_Contract *ContractCallerSession) ResolverOf(tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.ResolverOf(&_Contract.CallOpts, tokenId)
}

// ReverseNameOf is a free data retrieval call binding the contract method 0xbebec6b4.
//
// Solidity: function reverseNameOf(address addr) view returns(string)
func (_Contract *ContractCaller) ReverseNameOf(opts *bind.CallOpts, addr common.Address) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "reverseNameOf", addr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReverseNameOf is a free data retrieval call binding the contract method 0xbebec6b4.
//
// Solidity: function reverseNameOf(address addr) view returns(string)
func (_Contract *ContractSession) ReverseNameOf(addr common.Address) (string, error) {
	return _Contract.Contract.ReverseNameOf(&_Contract.CallOpts, addr)
}

// ReverseNameOf is a free data retrieval call binding the contract method 0xbebec6b4.
//
// Solidity: function reverseNameOf(address addr) view returns(string)
func (_Contract *ContractCallerSession) ReverseNameOf(addr common.Address) (string, error) {
	return _Contract.Contract.ReverseNameOf(&_Contract.CallOpts, addr)
}

// ReverseOf is a free data retrieval call binding the contract method 0x7e37479e.
//
// Solidity: function reverseOf(address addr) view returns(uint256)
func (_Contract *ContractCaller) ReverseOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "reverseOf", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReverseOf is a free data retrieval call binding the contract method 0x7e37479e.
//
// Solidity: function reverseOf(address addr) view returns(uint256)
func (_Contract *ContractSession) ReverseOf(addr common.Address) (*big.Int, error) {
	return _Contract.Contract.ReverseOf(&_Contract.CallOpts, addr)
}

// ReverseOf is a free data retrieval call binding the contract method 0x7e37479e.
//
// Solidity: function reverseOf(address addr) view returns(uint256)
func (_Contract *ContractCallerSession) ReverseOf(addr common.Address) (*big.Int, error) {
	return _Contract.Contract.ReverseOf(&_Contract.CallOpts, addr)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Contract *ContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Contract *ContractCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Contract *ContractSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Contract.Contract.TokenURI(&_Contract.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Contract *ContractCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Contract.Contract.TokenURI(&_Contract.CallOpts, tokenId)
}

// AddBlockchainNetworks is a paid mutator transaction binding the contract method 0xf0592359.
//
// Solidity: function addBlockchainNetworks(string[] networks, string[] families) returns()
func (_Contract *ContractTransactor) AddBlockchainNetworks(opts *bind.TransactOpts, networks []string, families []string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addBlockchainNetworks", networks, families)
}

// AddBlockchainNetworks is a paid mutator transaction binding the contract method 0xf0592359.
//
// Solidity: function addBlockchainNetworks(string[] networks, string[] families) returns()
func (_Contract *ContractSession) AddBlockchainNetworks(networks []string, families []string) (*types.Transaction, error) {
	return _Contract.Contract.AddBlockchainNetworks(&_Contract.TransactOpts, networks, families)
}

// AddBlockchainNetworks is a paid mutator transaction binding the contract method 0xf0592359.
//
// Solidity: function addBlockchainNetworks(string[] networks, string[] families) returns()
func (_Contract *ContractTransactorSession) AddBlockchainNetworks(networks []string, families []string) (*types.Transaction, error) {
	return _Contract.Contract.AddBlockchainNetworks(&_Contract.TransactOpts, networks, families)
}

// AddBlockchainNetworks0 is a paid mutator transaction binding the contract method 0xffad6f55.
//
// Solidity: function addBlockchainNetworks(string[] networks, string family) returns()
func (_Contract *ContractTransactor) AddBlockchainNetworks0(opts *bind.TransactOpts, networks []string, family string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addBlockchainNetworks0", networks, family)
}

// AddBlockchainNetworks0 is a paid mutator transaction binding the contract method 0xffad6f55.
//
// Solidity: function addBlockchainNetworks(string[] networks, string family) returns()
func (_Contract *ContractSession) AddBlockchainNetworks0(networks []string, family string) (*types.Transaction, error) {
	return _Contract.Contract.AddBlockchainNetworks0(&_Contract.TransactOpts, networks, family)
}

// AddBlockchainNetworks0 is a paid mutator transaction binding the contract method 0xffad6f55.
//
// Solidity: function addBlockchainNetworks(string[] networks, string family) returns()
func (_Contract *ContractTransactorSession) AddBlockchainNetworks0(networks []string, family string) (*types.Transaction, error) {
	return _Contract.Contract.AddBlockchainNetworks0(&_Contract.TransactOpts, networks, family)
}

// AddLegacyRecords is a paid mutator transaction binding the contract method 0x4bd79ed0.
//
// Solidity: function addLegacyRecords(string[] keys, string[][] legacyKeys) returns()
func (_Contract *ContractTransactor) AddLegacyRecords(opts *bind.TransactOpts, keys []string, legacyKeys [][]string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addLegacyRecords", keys, legacyKeys)
}

// AddLegacyRecords is a paid mutator transaction binding the contract method 0x4bd79ed0.
//
// Solidity: function addLegacyRecords(string[] keys, string[][] legacyKeys) returns()
func (_Contract *ContractSession) AddLegacyRecords(keys []string, legacyKeys [][]string) (*types.Transaction, error) {
	return _Contract.Contract.AddLegacyRecords(&_Contract.TransactOpts, keys, legacyKeys)
}

// AddLegacyRecords is a paid mutator transaction binding the contract method 0x4bd79ed0.
//
// Solidity: function addLegacyRecords(string[] keys, string[][] legacyKeys) returns()
func (_Contract *ContractTransactorSession) AddLegacyRecords(keys []string, legacyKeys [][]string) (*types.Transaction, error) {
	return _Contract.Contract.AddLegacyRecords(&_Contract.TransactOpts, keys, legacyKeys)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address unsRegistry, address cnsRegistry) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, unsRegistry common.Address, cnsRegistry common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", unsRegistry, cnsRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address unsRegistry, address cnsRegistry) returns()
func (_Contract *ContractSession) Initialize(unsRegistry common.Address, cnsRegistry common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, unsRegistry, cnsRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address unsRegistry, address cnsRegistry) returns()
func (_Contract *ContractTransactorSession) Initialize(unsRegistry common.Address, cnsRegistry common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, unsRegistry, cnsRegistry)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
// func (_Contract *ContractTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
// 	return _Contract.contract.Transact(opts, "multicall", data)
// }

// Multicall is a free data retrieval call binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) view returns(bytes[] results)
func (_Contract *ContractSession) Multicall(data [][]byte) ([][]byte, error) {
	return _Contract.Contract.Multicall(&_Contract.CallOpts, data)
}


// Multicall is a free data retrieval call binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) view returns(bytes[] results)
func (_Contract *ContractCallerSession) Multicall(data [][]byte) ([][]byte, error) {
	return _Contract.Contract.Multicall(&_Contract.CallOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
// func (_Contract *ContractTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
// 	return _Contract.Contract.Multicall(&_Contract.TransactOpts, data)
// }

// Multicall is a free data retrieval call binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) view returns(bytes[] results)
func (_Contract *ContractCaller) Multicall(opts *bind.CallOpts, data [][]byte) ([][]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "multicall", data)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}


// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address addr) returns()
func (_Contract *ContractTransactor) SetOwner(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setOwner", addr)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address addr) returns()
func (_Contract *ContractSession) SetOwner(addr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetOwner(&_Contract.TransactOpts, addr)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address addr) returns()
func (_Contract *ContractTransactorSession) SetOwner(addr common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetOwner(&_Contract.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// ContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contract contract.
type ContractInitializedIterator struct {
	Event *ContractInitialized // Event containing the contract specifics and raw log

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
func (it *ContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInitialized)
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
		it.Event = new(ContractInitialized)
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
func (it *ContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInitialized represents a Initialized event raised by the Contract contract.
type ContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractInitializedIterator{contract: _Contract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInitialized)
				if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) ParseInitialized(log types.Log) (*ContractInitialized, error) {
	event := new(ContractInitialized)
	if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetLegacyRecordsIterator is returned from FilterSetLegacyRecords and is used to iterate over the raw logs and unpacked data for SetLegacyRecords events raised by the Contract contract.
type ContractSetLegacyRecordsIterator struct {
	Event *ContractSetLegacyRecords // Event containing the contract specifics and raw log

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
func (it *ContractSetLegacyRecordsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetLegacyRecords)
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
		it.Event = new(ContractSetLegacyRecords)
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
func (it *ContractSetLegacyRecordsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetLegacyRecordsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetLegacyRecords represents a SetLegacyRecords event raised by the Contract contract.
type ContractSetLegacyRecords struct {
	TokenKey string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetLegacyRecords is a free log retrieval operation binding the contract event 0xb2b57b54285120d17f49f2490a39ef791441fd0ad1e6b6818c23262bf91f061c.
//
// Solidity: event SetLegacyRecords(string tokenKey)
func (_Contract *ContractFilterer) FilterSetLegacyRecords(opts *bind.FilterOpts) (*ContractSetLegacyRecordsIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetLegacyRecords")
	if err != nil {
		return nil, err
	}
	return &ContractSetLegacyRecordsIterator{contract: _Contract.contract, event: "SetLegacyRecords", logs: logs, sub: sub}, nil
}

// WatchSetLegacyRecords is a free log subscription operation binding the contract event 0xb2b57b54285120d17f49f2490a39ef791441fd0ad1e6b6818c23262bf91f061c.
//
// Solidity: event SetLegacyRecords(string tokenKey)
func (_Contract *ContractFilterer) WatchSetLegacyRecords(opts *bind.WatchOpts, sink chan<- *ContractSetLegacyRecords) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetLegacyRecords")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetLegacyRecords)
				if err := _Contract.contract.UnpackLog(event, "SetLegacyRecords", log); err != nil {
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

// ParseSetLegacyRecords is a log parse operation binding the contract event 0xb2b57b54285120d17f49f2490a39ef791441fd0ad1e6b6818c23262bf91f061c.
//
// Solidity: event SetLegacyRecords(string tokenKey)
func (_Contract *ContractFilterer) ParseSetLegacyRecords(log types.Log) (*ContractSetLegacyRecords, error) {
	event := new(ContractSetLegacyRecords)
	if err := _Contract.contract.UnpackLog(event, "SetLegacyRecords", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSetNetworkFamilyIterator is returned from FilterSetNetworkFamily and is used to iterate over the raw logs and unpacked data for SetNetworkFamily events raised by the Contract contract.
type ContractSetNetworkFamilyIterator struct {
	Event *ContractSetNetworkFamily // Event containing the contract specifics and raw log

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
func (it *ContractSetNetworkFamilyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetNetworkFamily)
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
		it.Event = new(ContractSetNetworkFamily)
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
func (it *ContractSetNetworkFamilyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetNetworkFamilyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetNetworkFamily represents a SetNetworkFamily event raised by the Contract contract.
type ContractSetNetworkFamily struct {
	Network string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetNetworkFamily is a free log retrieval operation binding the contract event 0x0bf4b04a0f6d7d0800284e1abb0f58f795c0a25b9088634b6ab847be51001dec.
//
// Solidity: event SetNetworkFamily(string network)
func (_Contract *ContractFilterer) FilterSetNetworkFamily(opts *bind.FilterOpts) (*ContractSetNetworkFamilyIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetNetworkFamily")
	if err != nil {
		return nil, err
	}
	return &ContractSetNetworkFamilyIterator{contract: _Contract.contract, event: "SetNetworkFamily", logs: logs, sub: sub}, nil
}

// WatchSetNetworkFamily is a free log subscription operation binding the contract event 0x0bf4b04a0f6d7d0800284e1abb0f58f795c0a25b9088634b6ab847be51001dec.
//
// Solidity: event SetNetworkFamily(string network)
func (_Contract *ContractFilterer) WatchSetNetworkFamily(opts *bind.WatchOpts, sink chan<- *ContractSetNetworkFamily) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetNetworkFamily")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetNetworkFamily)
				if err := _Contract.contract.UnpackLog(event, "SetNetworkFamily", log); err != nil {
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

// ParseSetNetworkFamily is a log parse operation binding the contract event 0x0bf4b04a0f6d7d0800284e1abb0f58f795c0a25b9088634b6ab847be51001dec.
//
// Solidity: event SetNetworkFamily(string network)
func (_Contract *ContractFilterer) ParseSetNetworkFamily(log types.Log) (*ContractSetNetworkFamily, error) {
	event := new(ContractSetNetworkFamily)
	if err := _Contract.contract.UnpackLog(event, "SetNetworkFamily", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
