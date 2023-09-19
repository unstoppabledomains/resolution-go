// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package namewrapperreader

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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractENS\",\"name\":\"_ens\",\"type\":\"address\"},{\"internalType\":\"contractIBaseRegistrar\",\"name\":\"_registrar\",\"type\":\"address\"},{\"internalType\":\"contractIMetadataService\",\"name\":\"_metadataService\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CannotUpgrade\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncompatibleParent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"IncorrectTargetOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectTokenType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"labelHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedLabelhash\",\"type\":\"bytes32\"}],\"name\":\"LabelMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"name\":\"LabelTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LabelTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NameIsNotWrapped\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"OperationProhibited\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Unauthorised\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"ControllerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"ExpiryExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"fuses\",\"type\":\"uint32\"}],\"name\":\"FusesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NameUnwrapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"fuses\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"NameWrapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_tokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"fuseMask\",\"type\":\"uint32\"}],\"name\":\"allFusesBurned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"canExtendSubnames\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"canModifyName\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"controllers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ens\",\"outputs\":[{\"internalType\":\"contractENS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentNode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"labelhash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"extendExpiry\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"fuses\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentNode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"labelhash\",\"type\":\"bytes32\"}],\"name\":\"isWrapped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"isWrapped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataService\",\"outputs\":[{\"internalType\":\"contractIMetadataService\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"names\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"wrappedOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"ownerControlledFuses\",\"type\":\"uint16\"}],\"name\":\"registerAndWrapETH2LD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"registrarExpiry\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registrar\",\"outputs\":[{\"internalType\":\"contractIBaseRegistrar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"renew\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentNode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"labelhash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"fuses\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"setChildFuses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"ownerControlledFuses\",\"type\":\"uint16\"}],\"name\":\"setFuses\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMetadataService\",\"name\":\"_metadataService\",\"type\":\"address\"}],\"name\":\"setMetadataService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentNode\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"fuses\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"setSubnodeOwner\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentNode\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"fuses\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"name\":\"setSubnodeRecord\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"ttl\",\"type\":\"uint64\"}],\"name\":\"setTTL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINameWrapperUpgrade\",\"name\":\"_upgradeAddress\",\"type\":\"address\"}],\"name\":\"setUpgradeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentNode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"labelhash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"labelhash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"registrant\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"unwrapETH2LD\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeContract\",\"outputs\":[{\"internalType\":\"contractINameWrapperUpgrade\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"wrappedOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"wrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"wrappedOwner\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"ownerControlledFuses\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"wrapETH2LD\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// Tokens is a free data retrieval call binding the contract method 0xed70554d.
//
// Solidity: function _tokens(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_tokens", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0xed70554d.
//
// Solidity: function _tokens(uint256 ) view returns(uint256)
func (_Contract *ContractSession) Tokens(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.Tokens(&_Contract.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xed70554d.
//
// Solidity: function _tokens(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) Tokens(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.Tokens(&_Contract.CallOpts, arg0)
}

// AllFusesBurned is a free data retrieval call binding the contract method 0xadf4960a.
//
// Solidity: function allFusesBurned(bytes32 node, uint32 fuseMask) view returns(bool)
func (_Contract *ContractCaller) AllFusesBurned(opts *bind.CallOpts, node [32]byte, fuseMask uint32) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allFusesBurned", node, fuseMask)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllFusesBurned is a free data retrieval call binding the contract method 0xadf4960a.
//
// Solidity: function allFusesBurned(bytes32 node, uint32 fuseMask) view returns(bool)
func (_Contract *ContractSession) AllFusesBurned(node [32]byte, fuseMask uint32) (bool, error) {
	return _Contract.Contract.AllFusesBurned(&_Contract.CallOpts, node, fuseMask)
}

// AllFusesBurned is a free data retrieval call binding the contract method 0xadf4960a.
//
// Solidity: function allFusesBurned(bytes32 node, uint32 fuseMask) view returns(bool)
func (_Contract *ContractCallerSession) AllFusesBurned(node [32]byte, fuseMask uint32) (bool, error) {
	return _Contract.Contract.AllFusesBurned(&_Contract.CallOpts, node, fuseMask)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Contract *ContractCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Contract *ContractSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Contract *ContractCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Contract *ContractCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Contract *ContractSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Contract.Contract.BalanceOfBatch(&_Contract.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Contract *ContractCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Contract.Contract.BalanceOfBatch(&_Contract.CallOpts, accounts, ids)
}

// CanExtendSubnames is a free data retrieval call binding the contract method 0x0e4cd725.
//
// Solidity: function canExtendSubnames(bytes32 node, address addr) view returns(bool)
func (_Contract *ContractCaller) CanExtendSubnames(opts *bind.CallOpts, node [32]byte, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "canExtendSubnames", node, addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanExtendSubnames is a free data retrieval call binding the contract method 0x0e4cd725.
//
// Solidity: function canExtendSubnames(bytes32 node, address addr) view returns(bool)
func (_Contract *ContractSession) CanExtendSubnames(node [32]byte, addr common.Address) (bool, error) {
	return _Contract.Contract.CanExtendSubnames(&_Contract.CallOpts, node, addr)
}

// CanExtendSubnames is a free data retrieval call binding the contract method 0x0e4cd725.
//
// Solidity: function canExtendSubnames(bytes32 node, address addr) view returns(bool)
func (_Contract *ContractCallerSession) CanExtendSubnames(node [32]byte, addr common.Address) (bool, error) {
	return _Contract.Contract.CanExtendSubnames(&_Contract.CallOpts, node, addr)
}

// CanModifyName is a free data retrieval call binding the contract method 0x41415eab.
//
// Solidity: function canModifyName(bytes32 node, address addr) view returns(bool)
func (_Contract *ContractCaller) CanModifyName(opts *bind.CallOpts, node [32]byte, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "canModifyName", node, addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanModifyName is a free data retrieval call binding the contract method 0x41415eab.
//
// Solidity: function canModifyName(bytes32 node, address addr) view returns(bool)
func (_Contract *ContractSession) CanModifyName(node [32]byte, addr common.Address) (bool, error) {
	return _Contract.Contract.CanModifyName(&_Contract.CallOpts, node, addr)
}

// CanModifyName is a free data retrieval call binding the contract method 0x41415eab.
//
// Solidity: function canModifyName(bytes32 node, address addr) view returns(bool)
func (_Contract *ContractCallerSession) CanModifyName(node [32]byte, addr common.Address) (bool, error) {
	return _Contract.Contract.CanModifyName(&_Contract.CallOpts, node, addr)
}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_Contract *ContractCaller) Controllers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "controllers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_Contract *ContractSession) Controllers(arg0 common.Address) (bool, error) {
	return _Contract.Contract.Controllers(&_Contract.CallOpts, arg0)
}

// Controllers is a free data retrieval call binding the contract method 0xda8c229e.
//
// Solidity: function controllers(address ) view returns(bool)
func (_Contract *ContractCallerSession) Controllers(arg0 common.Address) (bool, error) {
	return _Contract.Contract.Controllers(&_Contract.CallOpts, arg0)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() view returns(address)
func (_Contract *ContractCaller) Ens(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ens")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() view returns(address)
func (_Contract *ContractSession) Ens() (common.Address, error) {
	return _Contract.Contract.Ens(&_Contract.CallOpts)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() view returns(address)
func (_Contract *ContractCallerSession) Ens() (common.Address, error) {
	return _Contract.Contract.Ens(&_Contract.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 id) view returns(address operator)
func (_Contract *ContractCaller) GetApproved(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getApproved", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 id) view returns(address operator)
func (_Contract *ContractSession) GetApproved(id *big.Int) (common.Address, error) {
	return _Contract.Contract.GetApproved(&_Contract.CallOpts, id)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 id) view returns(address operator)
func (_Contract *ContractCallerSession) GetApproved(id *big.Int) (common.Address, error) {
	return _Contract.Contract.GetApproved(&_Contract.CallOpts, id)
}

// GetData is a free data retrieval call binding the contract method 0x0178fe3f.
//
// Solidity: function getData(uint256 id) view returns(address owner, uint32 fuses, uint64 expiry)
func (_Contract *ContractCaller) GetData(opts *bind.CallOpts, id *big.Int) (struct {
	Owner  common.Address
	Fuses  uint32
	Expiry uint64
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getData", id)

	outstruct := new(struct {
		Owner  common.Address
		Fuses  uint32
		Expiry uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Fuses = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Expiry = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetData is a free data retrieval call binding the contract method 0x0178fe3f.
//
// Solidity: function getData(uint256 id) view returns(address owner, uint32 fuses, uint64 expiry)
func (_Contract *ContractSession) GetData(id *big.Int) (struct {
	Owner  common.Address
	Fuses  uint32
	Expiry uint64
}, error) {
	return _Contract.Contract.GetData(&_Contract.CallOpts, id)
}

// GetData is a free data retrieval call binding the contract method 0x0178fe3f.
//
// Solidity: function getData(uint256 id) view returns(address owner, uint32 fuses, uint64 expiry)
func (_Contract *ContractCallerSession) GetData(id *big.Int) (struct {
	Owner  common.Address
	Fuses  uint32
	Expiry uint64
}, error) {
	return _Contract.Contract.GetData(&_Contract.CallOpts, id)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Contract *ContractCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Contract *ContractSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Contract.Contract.IsApprovedForAll(&_Contract.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Contract *ContractCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Contract.Contract.IsApprovedForAll(&_Contract.CallOpts, account, operator)
}

// IsWrapped is a free data retrieval call binding the contract method 0xd9a50c12.
//
// Solidity: function isWrapped(bytes32 parentNode, bytes32 labelhash) view returns(bool)
func (_Contract *ContractCaller) IsWrapped(opts *bind.CallOpts, parentNode [32]byte, labelhash [32]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isWrapped", parentNode, labelhash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWrapped is a free data retrieval call binding the contract method 0xd9a50c12.
//
// Solidity: function isWrapped(bytes32 parentNode, bytes32 labelhash) view returns(bool)
func (_Contract *ContractSession) IsWrapped(parentNode [32]byte, labelhash [32]byte) (bool, error) {
	return _Contract.Contract.IsWrapped(&_Contract.CallOpts, parentNode, labelhash)
}

// IsWrapped is a free data retrieval call binding the contract method 0xd9a50c12.
//
// Solidity: function isWrapped(bytes32 parentNode, bytes32 labelhash) view returns(bool)
func (_Contract *ContractCallerSession) IsWrapped(parentNode [32]byte, labelhash [32]byte) (bool, error) {
	return _Contract.Contract.IsWrapped(&_Contract.CallOpts, parentNode, labelhash)
}

// IsWrapped0 is a free data retrieval call binding the contract method 0xfd0cd0d9.
//
// Solidity: function isWrapped(bytes32 node) view returns(bool)
func (_Contract *ContractCaller) IsWrapped0(opts *bind.CallOpts, node [32]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isWrapped0", node)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWrapped0 is a free data retrieval call binding the contract method 0xfd0cd0d9.
//
// Solidity: function isWrapped(bytes32 node) view returns(bool)
func (_Contract *ContractSession) IsWrapped0(node [32]byte) (bool, error) {
	return _Contract.Contract.IsWrapped0(&_Contract.CallOpts, node)
}

// IsWrapped0 is a free data retrieval call binding the contract method 0xfd0cd0d9.
//
// Solidity: function isWrapped(bytes32 node) view returns(bool)
func (_Contract *ContractCallerSession) IsWrapped0(node [32]byte) (bool, error) {
	return _Contract.Contract.IsWrapped0(&_Contract.CallOpts, node)
}

// MetadataService is a free data retrieval call binding the contract method 0x53095467.
//
// Solidity: function metadataService() view returns(address)
func (_Contract *ContractCaller) MetadataService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "metadataService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetadataService is a free data retrieval call binding the contract method 0x53095467.
//
// Solidity: function metadataService() view returns(address)
func (_Contract *ContractSession) MetadataService() (common.Address, error) {
	return _Contract.Contract.MetadataService(&_Contract.CallOpts)
}

// MetadataService is a free data retrieval call binding the contract method 0x53095467.
//
// Solidity: function metadataService() view returns(address)
func (_Contract *ContractCallerSession) MetadataService() (common.Address, error) {
	return _Contract.Contract.MetadataService(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCallerSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Names is a free data retrieval call binding the contract method 0x20c38e2b.
//
// Solidity: function names(bytes32 ) view returns(bytes)
func (_Contract *ContractCaller) Names(opts *bind.CallOpts, arg0 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "names", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Names is a free data retrieval call binding the contract method 0x20c38e2b.
//
// Solidity: function names(bytes32 ) view returns(bytes)
func (_Contract *ContractSession) Names(arg0 [32]byte) ([]byte, error) {
	return _Contract.Contract.Names(&_Contract.CallOpts, arg0)
}

// Names is a free data retrieval call binding the contract method 0x20c38e2b.
//
// Solidity: function names(bytes32 ) view returns(bytes)
func (_Contract *ContractCallerSession) Names(arg0 [32]byte) ([]byte, error) {
	return _Contract.Contract.Names(&_Contract.CallOpts, arg0)
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
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Contract *ContractCaller) OwnerOf(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ownerOf", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Contract *ContractSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _Contract.Contract.OwnerOf(&_Contract.CallOpts, id)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Contract *ContractCallerSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _Contract.Contract.OwnerOf(&_Contract.CallOpts, id)
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_Contract *ContractCaller) Registrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "registrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_Contract *ContractSession) Registrar() (common.Address, error) {
	return _Contract.Contract.Registrar(&_Contract.CallOpts)
}

// Registrar is a free data retrieval call binding the contract method 0x2b20e397.
//
// Solidity: function registrar() view returns(address)
func (_Contract *ContractCallerSession) Registrar() (common.Address, error) {
	return _Contract.Contract.Registrar(&_Contract.CallOpts)
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

// UpgradeContract is a free data retrieval call binding the contract method 0x1f4e1504.
//
// Solidity: function upgradeContract() view returns(address)
func (_Contract *ContractCaller) UpgradeContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "upgradeContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradeContract is a free data retrieval call binding the contract method 0x1f4e1504.
//
// Solidity: function upgradeContract() view returns(address)
func (_Contract *ContractSession) UpgradeContract() (common.Address, error) {
	return _Contract.Contract.UpgradeContract(&_Contract.CallOpts)
}

// UpgradeContract is a free data retrieval call binding the contract method 0x1f4e1504.
//
// Solidity: function upgradeContract() view returns(address)
func (_Contract *ContractCallerSession) UpgradeContract() (common.Address, error) {
	return _Contract.Contract.UpgradeContract(&_Contract.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Contract *ContractCaller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "uri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Contract *ContractSession) Uri(tokenId *big.Int) (string, error) {
	return _Contract.Contract.Uri(&_Contract.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 tokenId) view returns(string)
func (_Contract *ContractCallerSession) Uri(tokenId *big.Int) (string, error) {
	return _Contract.Contract.Uri(&_Contract.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contract *ContractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contract *ContractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, to, tokenId)
}

// ExtendExpiry is a paid mutator transaction binding the contract method 0x6e5d6ad2.
//
// Solidity: function extendExpiry(bytes32 parentNode, bytes32 labelhash, uint64 expiry) returns(uint64)
func (_Contract *ContractTransactor) ExtendExpiry(opts *bind.TransactOpts, parentNode [32]byte, labelhash [32]byte, expiry uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "extendExpiry", parentNode, labelhash, expiry)
}

// ExtendExpiry is a paid mutator transaction binding the contract method 0x6e5d6ad2.
//
// Solidity: function extendExpiry(bytes32 parentNode, bytes32 labelhash, uint64 expiry) returns(uint64)
func (_Contract *ContractSession) ExtendExpiry(parentNode [32]byte, labelhash [32]byte, expiry uint64) (*types.Transaction, error) {
	return _Contract.Contract.ExtendExpiry(&_Contract.TransactOpts, parentNode, labelhash, expiry)
}

// ExtendExpiry is a paid mutator transaction binding the contract method 0x6e5d6ad2.
//
// Solidity: function extendExpiry(bytes32 parentNode, bytes32 labelhash, uint64 expiry) returns(uint64)
func (_Contract *ContractTransactorSession) ExtendExpiry(parentNode [32]byte, labelhash [32]byte, expiry uint64) (*types.Transaction, error) {
	return _Contract.Contract.ExtendExpiry(&_Contract.TransactOpts, parentNode, labelhash, expiry)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address to, address , uint256 tokenId, bytes data) returns(bytes4)
func (_Contract *ContractTransactor) OnERC721Received(opts *bind.TransactOpts, to common.Address, arg1 common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "onERC721Received", to, arg1, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address to, address , uint256 tokenId, bytes data) returns(bytes4)
func (_Contract *ContractSession) OnERC721Received(to common.Address, arg1 common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.OnERC721Received(&_Contract.TransactOpts, to, arg1, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address to, address , uint256 tokenId, bytes data) returns(bytes4)
func (_Contract *ContractTransactorSession) OnERC721Received(to common.Address, arg1 common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.OnERC721Received(&_Contract.TransactOpts, to, arg1, tokenId, data)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_Contract *ContractTransactor) RecoverFunds(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "recoverFunds", _token, _to, _amount)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_Contract *ContractSession) RecoverFunds(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RecoverFunds(&_Contract.TransactOpts, _token, _to, _amount)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_Contract *ContractTransactorSession) RecoverFunds(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RecoverFunds(&_Contract.TransactOpts, _token, _to, _amount)
}

// RegisterAndWrapETH2LD is a paid mutator transaction binding the contract method 0xa4014982.
//
// Solidity: function registerAndWrapETH2LD(string label, address wrappedOwner, uint256 duration, address resolver, uint16 ownerControlledFuses) returns(uint256 registrarExpiry)
func (_Contract *ContractTransactor) RegisterAndWrapETH2LD(opts *bind.TransactOpts, label string, wrappedOwner common.Address, duration *big.Int, resolver common.Address, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "registerAndWrapETH2LD", label, wrappedOwner, duration, resolver, ownerControlledFuses)
}

// RegisterAndWrapETH2LD is a paid mutator transaction binding the contract method 0xa4014982.
//
// Solidity: function registerAndWrapETH2LD(string label, address wrappedOwner, uint256 duration, address resolver, uint16 ownerControlledFuses) returns(uint256 registrarExpiry)
func (_Contract *ContractSession) RegisterAndWrapETH2LD(label string, wrappedOwner common.Address, duration *big.Int, resolver common.Address, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _Contract.Contract.RegisterAndWrapETH2LD(&_Contract.TransactOpts, label, wrappedOwner, duration, resolver, ownerControlledFuses)
}

// RegisterAndWrapETH2LD is a paid mutator transaction binding the contract method 0xa4014982.
//
// Solidity: function registerAndWrapETH2LD(string label, address wrappedOwner, uint256 duration, address resolver, uint16 ownerControlledFuses) returns(uint256 registrarExpiry)
func (_Contract *ContractTransactorSession) RegisterAndWrapETH2LD(label string, wrappedOwner common.Address, duration *big.Int, resolver common.Address, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _Contract.Contract.RegisterAndWrapETH2LD(&_Contract.TransactOpts, label, wrappedOwner, duration, resolver, ownerControlledFuses)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 tokenId, uint256 duration) returns(uint256 expires)
func (_Contract *ContractTransactor) Renew(opts *bind.TransactOpts, tokenId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renew", tokenId, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 tokenId, uint256 duration) returns(uint256 expires)
func (_Contract *ContractSession) Renew(tokenId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Renew(&_Contract.TransactOpts, tokenId, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 tokenId, uint256 duration) returns(uint256 expires)
func (_Contract *ContractTransactorSession) Renew(tokenId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Renew(&_Contract.TransactOpts, tokenId, duration)
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

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Contract *ContractTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Contract *ContractSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeBatchTransferFrom(&_Contract.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Contract *ContractTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeBatchTransferFrom(&_Contract.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Contract *ContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Contract *ContractSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom(&_Contract.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Contract *ContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom(&_Contract.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contract *ContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contract *ContractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetApprovalForAll(&_Contract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contract *ContractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetApprovalForAll(&_Contract.TransactOpts, operator, approved)
}

// SetChildFuses is a paid mutator transaction binding the contract method 0x33c69ea9.
//
// Solidity: function setChildFuses(bytes32 parentNode, bytes32 labelhash, uint32 fuses, uint64 expiry) returns()
func (_Contract *ContractTransactor) SetChildFuses(opts *bind.TransactOpts, parentNode [32]byte, labelhash [32]byte, fuses uint32, expiry uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setChildFuses", parentNode, labelhash, fuses, expiry)
}

// SetChildFuses is a paid mutator transaction binding the contract method 0x33c69ea9.
//
// Solidity: function setChildFuses(bytes32 parentNode, bytes32 labelhash, uint32 fuses, uint64 expiry) returns()
func (_Contract *ContractSession) SetChildFuses(parentNode [32]byte, labelhash [32]byte, fuses uint32, expiry uint64) (*types.Transaction, error) {
	return _Contract.Contract.SetChildFuses(&_Contract.TransactOpts, parentNode, labelhash, fuses, expiry)
}

// SetChildFuses is a paid mutator transaction binding the contract method 0x33c69ea9.
//
// Solidity: function setChildFuses(bytes32 parentNode, bytes32 labelhash, uint32 fuses, uint64 expiry) returns()
func (_Contract *ContractTransactorSession) SetChildFuses(parentNode [32]byte, labelhash [32]byte, fuses uint32, expiry uint64) (*types.Transaction, error) {
	return _Contract.Contract.SetChildFuses(&_Contract.TransactOpts, parentNode, labelhash, fuses, expiry)
}

// SetController is a paid mutator transaction binding the contract method 0xe0dba60f.
//
// Solidity: function setController(address controller, bool active) returns()
func (_Contract *ContractTransactor) SetController(opts *bind.TransactOpts, controller common.Address, active bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setController", controller, active)
}

// SetController is a paid mutator transaction binding the contract method 0xe0dba60f.
//
// Solidity: function setController(address controller, bool active) returns()
func (_Contract *ContractSession) SetController(controller common.Address, active bool) (*types.Transaction, error) {
	return _Contract.Contract.SetController(&_Contract.TransactOpts, controller, active)
}

// SetController is a paid mutator transaction binding the contract method 0xe0dba60f.
//
// Solidity: function setController(address controller, bool active) returns()
func (_Contract *ContractTransactorSession) SetController(controller common.Address, active bool) (*types.Transaction, error) {
	return _Contract.Contract.SetController(&_Contract.TransactOpts, controller, active)
}

// SetFuses is a paid mutator transaction binding the contract method 0x402906fc.
//
// Solidity: function setFuses(bytes32 node, uint16 ownerControlledFuses) returns(uint32)
func (_Contract *ContractTransactor) SetFuses(opts *bind.TransactOpts, node [32]byte, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setFuses", node, ownerControlledFuses)
}

// SetFuses is a paid mutator transaction binding the contract method 0x402906fc.
//
// Solidity: function setFuses(bytes32 node, uint16 ownerControlledFuses) returns(uint32)
func (_Contract *ContractSession) SetFuses(node [32]byte, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _Contract.Contract.SetFuses(&_Contract.TransactOpts, node, ownerControlledFuses)
}

// SetFuses is a paid mutator transaction binding the contract method 0x402906fc.
//
// Solidity: function setFuses(bytes32 node, uint16 ownerControlledFuses) returns(uint32)
func (_Contract *ContractTransactorSession) SetFuses(node [32]byte, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _Contract.Contract.SetFuses(&_Contract.TransactOpts, node, ownerControlledFuses)
}

// SetMetadataService is a paid mutator transaction binding the contract method 0x1534e177.
//
// Solidity: function setMetadataService(address _metadataService) returns()
func (_Contract *ContractTransactor) SetMetadataService(opts *bind.TransactOpts, _metadataService common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setMetadataService", _metadataService)
}

// SetMetadataService is a paid mutator transaction binding the contract method 0x1534e177.
//
// Solidity: function setMetadataService(address _metadataService) returns()
func (_Contract *ContractSession) SetMetadataService(_metadataService common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetMetadataService(&_Contract.TransactOpts, _metadataService)
}

// SetMetadataService is a paid mutator transaction binding the contract method 0x1534e177.
//
// Solidity: function setMetadataService(address _metadataService) returns()
func (_Contract *ContractTransactorSession) SetMetadataService(_metadataService common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetMetadataService(&_Contract.TransactOpts, _metadataService)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_Contract *ContractTransactor) SetRecord(opts *bind.TransactOpts, node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setRecord", node, owner, resolver, ttl)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_Contract *ContractSession) SetRecord(node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _Contract.Contract.SetRecord(&_Contract.TransactOpts, node, owner, resolver, ttl)
}

// SetRecord is a paid mutator transaction binding the contract method 0xcf408823.
//
// Solidity: function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) returns()
func (_Contract *ContractTransactorSession) SetRecord(node [32]byte, owner common.Address, resolver common.Address, ttl uint64) (*types.Transaction, error) {
	return _Contract.Contract.SetRecord(&_Contract.TransactOpts, node, owner, resolver, ttl)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_Contract *ContractTransactor) SetResolver(opts *bind.TransactOpts, node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setResolver", node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_Contract *ContractSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetResolver(&_Contract.TransactOpts, node, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x1896f70a.
//
// Solidity: function setResolver(bytes32 node, address resolver) returns()
func (_Contract *ContractTransactorSession) SetResolver(node [32]byte, resolver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetResolver(&_Contract.TransactOpts, node, resolver)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0xc658e086.
//
// Solidity: function setSubnodeOwner(bytes32 parentNode, string label, address owner, uint32 fuses, uint64 expiry) returns(bytes32 node)
func (_Contract *ContractTransactor) SetSubnodeOwner(opts *bind.TransactOpts, parentNode [32]byte, label string, owner common.Address, fuses uint32, expiry uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setSubnodeOwner", parentNode, label, owner, fuses, expiry)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0xc658e086.
//
// Solidity: function setSubnodeOwner(bytes32 parentNode, string label, address owner, uint32 fuses, uint64 expiry) returns(bytes32 node)
func (_Contract *ContractSession) SetSubnodeOwner(parentNode [32]byte, label string, owner common.Address, fuses uint32, expiry uint64) (*types.Transaction, error) {
	return _Contract.Contract.SetSubnodeOwner(&_Contract.TransactOpts, parentNode, label, owner, fuses, expiry)
}

// SetSubnodeOwner is a paid mutator transaction binding the contract method 0xc658e086.
//
// Solidity: function setSubnodeOwner(bytes32 parentNode, string label, address owner, uint32 fuses, uint64 expiry) returns(bytes32 node)
func (_Contract *ContractTransactorSession) SetSubnodeOwner(parentNode [32]byte, label string, owner common.Address, fuses uint32, expiry uint64) (*types.Transaction, error) {
	return _Contract.Contract.SetSubnodeOwner(&_Contract.TransactOpts, parentNode, label, owner, fuses, expiry)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x24c1af44.
//
// Solidity: function setSubnodeRecord(bytes32 parentNode, string label, address owner, address resolver, uint64 ttl, uint32 fuses, uint64 expiry) returns(bytes32 node)
func (_Contract *ContractTransactor) SetSubnodeRecord(opts *bind.TransactOpts, parentNode [32]byte, label string, owner common.Address, resolver common.Address, ttl uint64, fuses uint32, expiry uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setSubnodeRecord", parentNode, label, owner, resolver, ttl, fuses, expiry)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x24c1af44.
//
// Solidity: function setSubnodeRecord(bytes32 parentNode, string label, address owner, address resolver, uint64 ttl, uint32 fuses, uint64 expiry) returns(bytes32 node)
func (_Contract *ContractSession) SetSubnodeRecord(parentNode [32]byte, label string, owner common.Address, resolver common.Address, ttl uint64, fuses uint32, expiry uint64) (*types.Transaction, error) {
	return _Contract.Contract.SetSubnodeRecord(&_Contract.TransactOpts, parentNode, label, owner, resolver, ttl, fuses, expiry)
}

// SetSubnodeRecord is a paid mutator transaction binding the contract method 0x24c1af44.
//
// Solidity: function setSubnodeRecord(bytes32 parentNode, string label, address owner, address resolver, uint64 ttl, uint32 fuses, uint64 expiry) returns(bytes32 node)
func (_Contract *ContractTransactorSession) SetSubnodeRecord(parentNode [32]byte, label string, owner common.Address, resolver common.Address, ttl uint64, fuses uint32, expiry uint64) (*types.Transaction, error) {
	return _Contract.Contract.SetSubnodeRecord(&_Contract.TransactOpts, parentNode, label, owner, resolver, ttl, fuses, expiry)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_Contract *ContractTransactor) SetTTL(opts *bind.TransactOpts, node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setTTL", node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_Contract *ContractSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _Contract.Contract.SetTTL(&_Contract.TransactOpts, node, ttl)
}

// SetTTL is a paid mutator transaction binding the contract method 0x14ab9038.
//
// Solidity: function setTTL(bytes32 node, uint64 ttl) returns()
func (_Contract *ContractTransactorSession) SetTTL(node [32]byte, ttl uint64) (*types.Transaction, error) {
	return _Contract.Contract.SetTTL(&_Contract.TransactOpts, node, ttl)
}

// SetUpgradeContract is a paid mutator transaction binding the contract method 0xb6bcad26.
//
// Solidity: function setUpgradeContract(address _upgradeAddress) returns()
func (_Contract *ContractTransactor) SetUpgradeContract(opts *bind.TransactOpts, _upgradeAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setUpgradeContract", _upgradeAddress)
}

// SetUpgradeContract is a paid mutator transaction binding the contract method 0xb6bcad26.
//
// Solidity: function setUpgradeContract(address _upgradeAddress) returns()
func (_Contract *ContractSession) SetUpgradeContract(_upgradeAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetUpgradeContract(&_Contract.TransactOpts, _upgradeAddress)
}

// SetUpgradeContract is a paid mutator transaction binding the contract method 0xb6bcad26.
//
// Solidity: function setUpgradeContract(address _upgradeAddress) returns()
func (_Contract *ContractTransactorSession) SetUpgradeContract(_upgradeAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetUpgradeContract(&_Contract.TransactOpts, _upgradeAddress)
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

// Unwrap is a paid mutator transaction binding the contract method 0xd8c9921a.
//
// Solidity: function unwrap(bytes32 parentNode, bytes32 labelhash, address controller) returns()
func (_Contract *ContractTransactor) Unwrap(opts *bind.TransactOpts, parentNode [32]byte, labelhash [32]byte, controller common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unwrap", parentNode, labelhash, controller)
}

// Unwrap is a paid mutator transaction binding the contract method 0xd8c9921a.
//
// Solidity: function unwrap(bytes32 parentNode, bytes32 labelhash, address controller) returns()
func (_Contract *ContractSession) Unwrap(parentNode [32]byte, labelhash [32]byte, controller common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Unwrap(&_Contract.TransactOpts, parentNode, labelhash, controller)
}

// Unwrap is a paid mutator transaction binding the contract method 0xd8c9921a.
//
// Solidity: function unwrap(bytes32 parentNode, bytes32 labelhash, address controller) returns()
func (_Contract *ContractTransactorSession) Unwrap(parentNode [32]byte, labelhash [32]byte, controller common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Unwrap(&_Contract.TransactOpts, parentNode, labelhash, controller)
}

// UnwrapETH2LD is a paid mutator transaction binding the contract method 0x8b4dfa75.
//
// Solidity: function unwrapETH2LD(bytes32 labelhash, address registrant, address controller) returns()
func (_Contract *ContractTransactor) UnwrapETH2LD(opts *bind.TransactOpts, labelhash [32]byte, registrant common.Address, controller common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unwrapETH2LD", labelhash, registrant, controller)
}

// UnwrapETH2LD is a paid mutator transaction binding the contract method 0x8b4dfa75.
//
// Solidity: function unwrapETH2LD(bytes32 labelhash, address registrant, address controller) returns()
func (_Contract *ContractSession) UnwrapETH2LD(labelhash [32]byte, registrant common.Address, controller common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UnwrapETH2LD(&_Contract.TransactOpts, labelhash, registrant, controller)
}

// UnwrapETH2LD is a paid mutator transaction binding the contract method 0x8b4dfa75.
//
// Solidity: function unwrapETH2LD(bytes32 labelhash, address registrant, address controller) returns()
func (_Contract *ContractTransactorSession) UnwrapETH2LD(labelhash [32]byte, registrant common.Address, controller common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UnwrapETH2LD(&_Contract.TransactOpts, labelhash, registrant, controller)
}

// Upgrade is a paid mutator transaction binding the contract method 0xc93ab3fd.
//
// Solidity: function upgrade(bytes name, bytes extraData) returns()
func (_Contract *ContractTransactor) Upgrade(opts *bind.TransactOpts, name []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "upgrade", name, extraData)
}

// Upgrade is a paid mutator transaction binding the contract method 0xc93ab3fd.
//
// Solidity: function upgrade(bytes name, bytes extraData) returns()
func (_Contract *ContractSession) Upgrade(name []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.Upgrade(&_Contract.TransactOpts, name, extraData)
}

// Upgrade is a paid mutator transaction binding the contract method 0xc93ab3fd.
//
// Solidity: function upgrade(bytes name, bytes extraData) returns()
func (_Contract *ContractTransactorSession) Upgrade(name []byte, extraData []byte) (*types.Transaction, error) {
	return _Contract.Contract.Upgrade(&_Contract.TransactOpts, name, extraData)
}

// Wrap is a paid mutator transaction binding the contract method 0xeb8ae530.
//
// Solidity: function wrap(bytes name, address wrappedOwner, address resolver) returns()
func (_Contract *ContractTransactor) Wrap(opts *bind.TransactOpts, name []byte, wrappedOwner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "wrap", name, wrappedOwner, resolver)
}

// Wrap is a paid mutator transaction binding the contract method 0xeb8ae530.
//
// Solidity: function wrap(bytes name, address wrappedOwner, address resolver) returns()
func (_Contract *ContractSession) Wrap(name []byte, wrappedOwner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Wrap(&_Contract.TransactOpts, name, wrappedOwner, resolver)
}

// Wrap is a paid mutator transaction binding the contract method 0xeb8ae530.
//
// Solidity: function wrap(bytes name, address wrappedOwner, address resolver) returns()
func (_Contract *ContractTransactorSession) Wrap(name []byte, wrappedOwner common.Address, resolver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Wrap(&_Contract.TransactOpts, name, wrappedOwner, resolver)
}

// WrapETH2LD is a paid mutator transaction binding the contract method 0x8cf8b41e.
//
// Solidity: function wrapETH2LD(string label, address wrappedOwner, uint16 ownerControlledFuses, address resolver) returns(uint64 expiry)
func (_Contract *ContractTransactor) WrapETH2LD(opts *bind.TransactOpts, label string, wrappedOwner common.Address, ownerControlledFuses uint16, resolver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "wrapETH2LD", label, wrappedOwner, ownerControlledFuses, resolver)
}

// WrapETH2LD is a paid mutator transaction binding the contract method 0x8cf8b41e.
//
// Solidity: function wrapETH2LD(string label, address wrappedOwner, uint16 ownerControlledFuses, address resolver) returns(uint64 expiry)
func (_Contract *ContractSession) WrapETH2LD(label string, wrappedOwner common.Address, ownerControlledFuses uint16, resolver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.WrapETH2LD(&_Contract.TransactOpts, label, wrappedOwner, ownerControlledFuses, resolver)
}

// WrapETH2LD is a paid mutator transaction binding the contract method 0x8cf8b41e.
//
// Solidity: function wrapETH2LD(string label, address wrappedOwner, uint16 ownerControlledFuses, address resolver) returns(uint64 expiry)
func (_Contract *ContractTransactorSession) WrapETH2LD(label string, wrappedOwner common.Address, ownerControlledFuses uint16, resolver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.WrapETH2LD(&_Contract.TransactOpts, label, wrappedOwner, ownerControlledFuses, resolver)
}

// ContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contract contract.
type ContractApprovalIterator struct {
	Event *ContractApproval // Event containing the contract specifics and raw log

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
func (it *ContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApproval)
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
		it.Event = new(ContractApproval)
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
func (it *ContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApproval represents a Approval event raised by the Contract contract.
type ContractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contract *ContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalIterator{contract: _Contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contract *ContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApproval)
				if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contract *ContractFilterer) ParseApproval(log types.Log) (*ContractApproval, error) {
	event := new(ContractApproval)
	if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Contract contract.
type ContractApprovalForAllIterator struct {
	Event *ContractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApprovalForAll)
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
		it.Event = new(ContractApprovalForAll)
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
func (it *ContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApprovalForAll represents a ApprovalForAll event raised by the Contract contract.
type ContractApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Contract *ContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*ContractApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalForAllIterator{contract: _Contract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Contract *ContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ContractApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApprovalForAll)
				if err := _Contract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Contract *ContractFilterer) ParseApprovalForAll(log types.Log) (*ContractApprovalForAll, error) {
	event := new(ContractApprovalForAll)
	if err := _Contract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractControllerChangedIterator is returned from FilterControllerChanged and is used to iterate over the raw logs and unpacked data for ControllerChanged events raised by the Contract contract.
type ContractControllerChangedIterator struct {
	Event *ContractControllerChanged // Event containing the contract specifics and raw log

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
func (it *ContractControllerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractControllerChanged)
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
		it.Event = new(ContractControllerChanged)
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
func (it *ContractControllerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractControllerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractControllerChanged represents a ControllerChanged event raised by the Contract contract.
type ContractControllerChanged struct {
	Controller common.Address
	Active     bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterControllerChanged is a free log retrieval operation binding the contract event 0x4c97694570a07277810af7e5669ffd5f6a2d6b74b6e9a274b8b870fd5114cf87.
//
// Solidity: event ControllerChanged(address indexed controller, bool active)
func (_Contract *ContractFilterer) FilterControllerChanged(opts *bind.FilterOpts, controller []common.Address) (*ContractControllerChangedIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ControllerChanged", controllerRule)
	if err != nil {
		return nil, err
	}
	return &ContractControllerChangedIterator{contract: _Contract.contract, event: "ControllerChanged", logs: logs, sub: sub}, nil
}

// WatchControllerChanged is a free log subscription operation binding the contract event 0x4c97694570a07277810af7e5669ffd5f6a2d6b74b6e9a274b8b870fd5114cf87.
//
// Solidity: event ControllerChanged(address indexed controller, bool active)
func (_Contract *ContractFilterer) WatchControllerChanged(opts *bind.WatchOpts, sink chan<- *ContractControllerChanged, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ControllerChanged", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractControllerChanged)
				if err := _Contract.contract.UnpackLog(event, "ControllerChanged", log); err != nil {
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

// ParseControllerChanged is a log parse operation binding the contract event 0x4c97694570a07277810af7e5669ffd5f6a2d6b74b6e9a274b8b870fd5114cf87.
//
// Solidity: event ControllerChanged(address indexed controller, bool active)
func (_Contract *ContractFilterer) ParseControllerChanged(log types.Log) (*ContractControllerChanged, error) {
	event := new(ContractControllerChanged)
	if err := _Contract.contract.UnpackLog(event, "ControllerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractExpiryExtendedIterator is returned from FilterExpiryExtended and is used to iterate over the raw logs and unpacked data for ExpiryExtended events raised by the Contract contract.
type ContractExpiryExtendedIterator struct {
	Event *ContractExpiryExtended // Event containing the contract specifics and raw log

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
func (it *ContractExpiryExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractExpiryExtended)
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
		it.Event = new(ContractExpiryExtended)
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
func (it *ContractExpiryExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractExpiryExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractExpiryExtended represents a ExpiryExtended event raised by the Contract contract.
type ContractExpiryExtended struct {
	Node   [32]byte
	Expiry uint64
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExpiryExtended is a free log retrieval operation binding the contract event 0xf675815a0817338f93a7da433f6bd5f5542f1029b11b455191ac96c7f6a9b132.
//
// Solidity: event ExpiryExtended(bytes32 indexed node, uint64 expiry)
func (_Contract *ContractFilterer) FilterExpiryExtended(opts *bind.FilterOpts, node [][32]byte) (*ContractExpiryExtendedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ExpiryExtended", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ContractExpiryExtendedIterator{contract: _Contract.contract, event: "ExpiryExtended", logs: logs, sub: sub}, nil
}

// WatchExpiryExtended is a free log subscription operation binding the contract event 0xf675815a0817338f93a7da433f6bd5f5542f1029b11b455191ac96c7f6a9b132.
//
// Solidity: event ExpiryExtended(bytes32 indexed node, uint64 expiry)
func (_Contract *ContractFilterer) WatchExpiryExtended(opts *bind.WatchOpts, sink chan<- *ContractExpiryExtended, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ExpiryExtended", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractExpiryExtended)
				if err := _Contract.contract.UnpackLog(event, "ExpiryExtended", log); err != nil {
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

// ParseExpiryExtended is a log parse operation binding the contract event 0xf675815a0817338f93a7da433f6bd5f5542f1029b11b455191ac96c7f6a9b132.
//
// Solidity: event ExpiryExtended(bytes32 indexed node, uint64 expiry)
func (_Contract *ContractFilterer) ParseExpiryExtended(log types.Log) (*ContractExpiryExtended, error) {
	event := new(ContractExpiryExtended)
	if err := _Contract.contract.UnpackLog(event, "ExpiryExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractFusesSetIterator is returned from FilterFusesSet and is used to iterate over the raw logs and unpacked data for FusesSet events raised by the Contract contract.
type ContractFusesSetIterator struct {
	Event *ContractFusesSet // Event containing the contract specifics and raw log

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
func (it *ContractFusesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFusesSet)
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
		it.Event = new(ContractFusesSet)
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
func (it *ContractFusesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractFusesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractFusesSet represents a FusesSet event raised by the Contract contract.
type ContractFusesSet struct {
	Node  [32]byte
	Fuses uint32
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFusesSet is a free log retrieval operation binding the contract event 0x39873f00c80f4f94b7bd1594aebcf650f003545b74824d57ddf4939e3ff3a34b.
//
// Solidity: event FusesSet(bytes32 indexed node, uint32 fuses)
func (_Contract *ContractFilterer) FilterFusesSet(opts *bind.FilterOpts, node [][32]byte) (*ContractFusesSetIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FusesSet", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ContractFusesSetIterator{contract: _Contract.contract, event: "FusesSet", logs: logs, sub: sub}, nil
}

// WatchFusesSet is a free log subscription operation binding the contract event 0x39873f00c80f4f94b7bd1594aebcf650f003545b74824d57ddf4939e3ff3a34b.
//
// Solidity: event FusesSet(bytes32 indexed node, uint32 fuses)
func (_Contract *ContractFilterer) WatchFusesSet(opts *bind.WatchOpts, sink chan<- *ContractFusesSet, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FusesSet", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractFusesSet)
				if err := _Contract.contract.UnpackLog(event, "FusesSet", log); err != nil {
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

// ParseFusesSet is a log parse operation binding the contract event 0x39873f00c80f4f94b7bd1594aebcf650f003545b74824d57ddf4939e3ff3a34b.
//
// Solidity: event FusesSet(bytes32 indexed node, uint32 fuses)
func (_Contract *ContractFilterer) ParseFusesSet(log types.Log) (*ContractFusesSet, error) {
	event := new(ContractFusesSet)
	if err := _Contract.contract.UnpackLog(event, "FusesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNameUnwrappedIterator is returned from FilterNameUnwrapped and is used to iterate over the raw logs and unpacked data for NameUnwrapped events raised by the Contract contract.
type ContractNameUnwrappedIterator struct {
	Event *ContractNameUnwrapped // Event containing the contract specifics and raw log

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
func (it *ContractNameUnwrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNameUnwrapped)
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
		it.Event = new(ContractNameUnwrapped)
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
func (it *ContractNameUnwrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNameUnwrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNameUnwrapped represents a NameUnwrapped event raised by the Contract contract.
type ContractNameUnwrapped struct {
	Node  [32]byte
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNameUnwrapped is a free log retrieval operation binding the contract event 0xee2ba1195c65bcf218a83d874335c6bf9d9067b4c672f3c3bf16cf40de7586c4.
//
// Solidity: event NameUnwrapped(bytes32 indexed node, address owner)
func (_Contract *ContractFilterer) FilterNameUnwrapped(opts *bind.FilterOpts, node [][32]byte) (*ContractNameUnwrappedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NameUnwrapped", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ContractNameUnwrappedIterator{contract: _Contract.contract, event: "NameUnwrapped", logs: logs, sub: sub}, nil
}

// WatchNameUnwrapped is a free log subscription operation binding the contract event 0xee2ba1195c65bcf218a83d874335c6bf9d9067b4c672f3c3bf16cf40de7586c4.
//
// Solidity: event NameUnwrapped(bytes32 indexed node, address owner)
func (_Contract *ContractFilterer) WatchNameUnwrapped(opts *bind.WatchOpts, sink chan<- *ContractNameUnwrapped, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NameUnwrapped", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNameUnwrapped)
				if err := _Contract.contract.UnpackLog(event, "NameUnwrapped", log); err != nil {
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

// ParseNameUnwrapped is a log parse operation binding the contract event 0xee2ba1195c65bcf218a83d874335c6bf9d9067b4c672f3c3bf16cf40de7586c4.
//
// Solidity: event NameUnwrapped(bytes32 indexed node, address owner)
func (_Contract *ContractFilterer) ParseNameUnwrapped(log types.Log) (*ContractNameUnwrapped, error) {
	event := new(ContractNameUnwrapped)
	if err := _Contract.contract.UnpackLog(event, "NameUnwrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNameWrappedIterator is returned from FilterNameWrapped and is used to iterate over the raw logs and unpacked data for NameWrapped events raised by the Contract contract.
type ContractNameWrappedIterator struct {
	Event *ContractNameWrapped // Event containing the contract specifics and raw log

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
func (it *ContractNameWrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNameWrapped)
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
		it.Event = new(ContractNameWrapped)
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
func (it *ContractNameWrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNameWrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNameWrapped represents a NameWrapped event raised by the Contract contract.
type ContractNameWrapped struct {
	Node   [32]byte
	Name   []byte
	Owner  common.Address
	Fuses  uint32
	Expiry uint64
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNameWrapped is a free log retrieval operation binding the contract event 0x8ce7013e8abebc55c3890a68f5a27c67c3f7efa64e584de5fb22363c606fd340.
//
// Solidity: event NameWrapped(bytes32 indexed node, bytes name, address owner, uint32 fuses, uint64 expiry)
func (_Contract *ContractFilterer) FilterNameWrapped(opts *bind.FilterOpts, node [][32]byte) (*ContractNameWrappedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NameWrapped", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ContractNameWrappedIterator{contract: _Contract.contract, event: "NameWrapped", logs: logs, sub: sub}, nil
}

// WatchNameWrapped is a free log subscription operation binding the contract event 0x8ce7013e8abebc55c3890a68f5a27c67c3f7efa64e584de5fb22363c606fd340.
//
// Solidity: event NameWrapped(bytes32 indexed node, bytes name, address owner, uint32 fuses, uint64 expiry)
func (_Contract *ContractFilterer) WatchNameWrapped(opts *bind.WatchOpts, sink chan<- *ContractNameWrapped, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NameWrapped", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNameWrapped)
				if err := _Contract.contract.UnpackLog(event, "NameWrapped", log); err != nil {
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

// ParseNameWrapped is a log parse operation binding the contract event 0x8ce7013e8abebc55c3890a68f5a27c67c3f7efa64e584de5fb22363c606fd340.
//
// Solidity: event NameWrapped(bytes32 indexed node, bytes name, address owner, uint32 fuses, uint64 expiry)
func (_Contract *ContractFilterer) ParseNameWrapped(log types.Log) (*ContractNameWrapped, error) {
	event := new(ContractNameWrapped)
	if err := _Contract.contract.UnpackLog(event, "NameWrapped", log); err != nil {
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

// ContractTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Contract contract.
type ContractTransferBatchIterator struct {
	Event *ContractTransferBatch // Event containing the contract specifics and raw log

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
func (it *ContractTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransferBatch)
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
		it.Event = new(ContractTransferBatch)
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
func (it *ContractTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransferBatch represents a TransferBatch event raised by the Contract contract.
type ContractTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Contract *ContractFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ContractTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractTransferBatchIterator{contract: _Contract.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Contract *ContractFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ContractTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransferBatch)
				if err := _Contract.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Contract *ContractFilterer) ParseTransferBatch(log types.Log) (*ContractTransferBatch, error) {
	event := new(ContractTransferBatch)
	if err := _Contract.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Contract contract.
type ContractTransferSingleIterator struct {
	Event *ContractTransferSingle // Event containing the contract specifics and raw log

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
func (it *ContractTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransferSingle)
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
		it.Event = new(ContractTransferSingle)
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
func (it *ContractTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransferSingle represents a TransferSingle event raised by the Contract contract.
type ContractTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Contract *ContractFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ContractTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractTransferSingleIterator{contract: _Contract.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Contract *ContractFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ContractTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransferSingle)
				if err := _Contract.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Contract *ContractFilterer) ParseTransferSingle(log types.Log) (*ContractTransferSingle, error) {
	event := new(ContractTransferSingle)
	if err := _Contract.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Contract contract.
type ContractURIIterator struct {
	Event *ContractURI // Event containing the contract specifics and raw log

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
func (it *ContractURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractURI)
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
		it.Event = new(ContractURI)
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
func (it *ContractURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractURI represents a URI event raised by the Contract contract.
type ContractURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Contract *ContractFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ContractURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ContractURIIterator{contract: _Contract.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Contract *ContractFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ContractURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractURI)
				if err := _Contract.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Contract *ContractFilterer) ParseURI(log types.Log) (*ContractURI, error) {
	event := new(ContractURI)
	if err := _Contract.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
