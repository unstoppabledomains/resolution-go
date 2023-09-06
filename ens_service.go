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
	"github.com/unstoppabledomains/resolution-go/v3/utils"
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

func (e EnsService) labelNamehash(domainName string) common.Hash {
	label, _ := utils.SplitDomain(domainName)

	return kns.Erc721Hash(label)
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

type ownerResult struct {
	address string
	err     error
	source  string
}

func (e EnsService) getOwnerFromNameWrapper(namehash common.Hash, ch chan<- ownerResult) {
	address, err := e.nameWrapperContract.OwnerOf(&bind.CallOpts{Pending: false}, namehash.Big())
	ch <- ownerResult{address: address.Hex(), err: err, source: "nameWrapper"}
}

func (e EnsService) getOwnerFromRegistry(namehash common.Hash, ch chan<- ownerResult) {
	address, err := e.ensRegistryContract.Owner(&bind.CallOpts{Pending: false}, namehash)
	ch <- ownerResult{address: address.Hex(), err: err, source: "registry"}
}

func (e EnsService) ownerOf(namehash common.Hash) (string, error) {

	ch := make(chan ownerResult, 2)

	go e.getOwnerFromNameWrapper(namehash, ch)
	go e.getOwnerFromRegistry(namehash, ch)

	address1 := <-ch
	address2 := <-ch

	var registryResult ownerResult
	var nameWrapperResult ownerResult

	if address1.source == "registry" {
		registryResult = address1
		nameWrapperResult = address2
	} else {
		registryResult = address2
		nameWrapperResult = address1
	}

	if registryResult.err != nil {
		return "", registryResult.err
	}

	if nameWrapperResult.err != nil {
		return registryResult.address, nil
	}

	if nameWrapperResult.address == NullAddress {
		return registryResult.address, nil
	}

	return nameWrapperResult.address, nil
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

func (e EnsService) isWrapped(namehash common.Hash) (bool, error) {
	return e.nameWrapperContract.IsWrapped0(&bind.CallOpts{Pending: false}, namehash)
}
