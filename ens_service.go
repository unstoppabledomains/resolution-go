package resolution

import (
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	kns "github.com/jgimeno/go-namehash"
	"github.com/unstoppabledomains/resolution-go/v3/ens/contracts/namewrapperreader"
	"github.com/unstoppabledomains/resolution-go/v3/ens/contracts/registryreader"
	"github.com/unstoppabledomains/resolution-go/v3/ens/contracts/resolverreader"
)

const (
	NullAddress = "0x0000000000000000000000000000000000000000"
)

// Ens is a naming service
type EnsService struct {
	ensRegistryContract   *registryreader.Contract
	nameWrapperContract   *namewrapperreader.Contract
	ensResolverContract   *resolverreader.Contract
	contractBackend       bind.ContractBackend
	networkId             int
	blockchainProviderUrl string
}

func (e EnsService) domainExists(namehash common.Hash) (bool, error) {
	return e.ensRegistryContract.RecordExists(&bind.CallOpts{Pending: false}, namehash)
}

func (e EnsService) namehash(domainName string) common.Hash {
	namehash := kns.NameHash(domainName)
	return namehash
}

func (e EnsService) resolver(namehash common.Hash) (string, error) {
	resolverAddress, err := e.ensRegistryContract.Resolver(&bind.CallOpts{Pending: false}, namehash)

	if err != nil {
		return "", err
	}

	return resolverAddress.Hex(), nil
}

func (e EnsService) reverseOf(addr string) (string, error) {
	namehash := e.namehash(addr[2:] + ".addr.reverse")

	resolverAddress, err := e.resolver(namehash)

	if err != nil || resolverAddress == NullAddress {
		return "", err
	}

	resolverContract, err := resolverreader.NewContract(common.HexToAddress(resolverAddress), e.contractBackend)

	if err != nil {
		return "", err
	}

	name, err := resolverContract.Name(&bind.CallOpts{Pending: false}, namehash)

	if err != nil {
		return "", err
	}

	return name, nil
}

func (e EnsService) ownerOf(namehash common.Hash) (string, error) {
	address, err := e.nameWrapperContract.OwnerOf(&bind.CallOpts{Pending: false}, namehash.Big())

	if err != nil {
		return "", err
	}

	return address.Hex(), nil
}

func (e EnsService) addrRecord(resolverAddress string, namehash common.Hash) (string, error) {
	resolverContract, err := resolverreader.NewContract(common.HexToAddress(resolverAddress), e.contractBackend)

	if err != nil {
		return "", err
	}

	addr, err := resolverContract.Addr(&bind.CallOpts{Pending: false}, namehash)

	if err != nil {
		return "", err
	}

	return addr.Hex(), nil
}

func (e EnsService) addrCoinRecord(resolverAddress string, namehash common.Hash, coin *big.Int) (string, error) {
	resolverContract, err := resolverreader.NewContract(common.HexToAddress(resolverAddress), e.contractBackend)

	if err != nil {
		return "", err
	}

	addr, err := resolverContract.Addr0(&bind.CallOpts{Pending: false}, namehash, coin)

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(addr), nil
}

func (e EnsService) contenthashRecord(resolverAddress string, namehash common.Hash) (string, error) {
	resolverContract, err := resolverreader.NewContract(common.HexToAddress(resolverAddress), e.contractBackend)

	if err != nil {
		return "", err
	}

	contentHash, err := resolverContract.Contenthash(&bind.CallOpts{Pending: false}, namehash)

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(contentHash), nil
}

func (e EnsService) textRecord(resolverAddress string, namehash common.Hash, key string) (string, error) {
	resolverContract, err := resolverreader.NewContract(common.HexToAddress(resolverAddress), e.contractBackend)

	if err != nil {
		return "", err
	}

	return resolverContract.Text(&bind.CallOpts{Pending: false}, namehash, key)
}
