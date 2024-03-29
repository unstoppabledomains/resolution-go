package resolution

import (
	"encoding/hex"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/unstoppabledomains/resolution-go/v3/ens/contracts/namewrapperreader"
	"github.com/unstoppabledomains/resolution-go/v3/ens/contracts/registrarreader"
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
	baseRegistrarContract *registrarreader.Contract
	metadataClient        MetadataClient
	contractBackend       bind.ContractBackend
	registryAddress       common.Address
	networkId             int
	blockchainProviderUrl string
}

type ensGenericResult struct {
	result any
	err    error
	source string
}

///////////////////////////
// exist, expiry funtions//
///////////////////////////

func (e EnsService) domainExists(namehash common.Hash) (bool, error) {
	return e.ensRegistryContract.RecordExists(&bind.CallOpts{Pending: false}, namehash)
}

func (e EnsService) domainExpiry(domain string) (time.Time, error) {
	if utils.IsSubdomain(domain) { // if is Subdomain, return the expiration of the parent domain
		domain = utils.GetParentDomain(domain)
	}

	registrarAddress, err := e.getRegistrarAddress(domain)

	if err != nil {
		return time.Unix(0, 0), err
	}

	registrarContract, err := registrarreader.NewContract(common.HexToAddress(registrarAddress), e.contractBackend)

	if err != nil {
		return time.Unix(0, 0), err
	}

	expiryTS, err := registrarContract.NameExpires(&bind.CallOpts{Pending: false}, e.labelHash(domain).Big())

	if err != nil {
		return time.Unix(0, 0), err
	}

	if expiryTS.Int64() == 0 {
		return time.Unix(0, 0), &DomainNotConfiguredError{DomainName: domain}
	}

	return time.Unix(expiryTS.Int64(), 0), nil
}

//////////////////////////
// namehash functions 	//
//////////////////////////

func (e EnsService) namehash(domainName string) common.Hash {
	return utils.UnsEnsNameHash(domainName)
}

func (e EnsService) labelHash(domainName string) common.Hash {
	label, _ := utils.SplitDomain(domainName)
	return utils.Erc721Hash(label)
}

//////////////////////////
// resolver functions 	//
//////////////////////////

func (e EnsService) resolver(namehash common.Hash) (string, error) {
	resolverAddress, err := e.ensRegistryContract.Resolver(&bind.CallOpts{Pending: false}, namehash)

	if err != nil {
		return "", err
	}

	if resolverAddress.Hex() == NullAddress {
		return "", &DomainNotConfiguredError{DomainName: namehash.Hex()}
	}

	return resolverAddress.Hex(), nil
}

//////////////////////////
// registrar functions 	//
//////////////////////////

func (e EnsService) getRegistrarAddress(domainName string) (string, error) {
	parent := utils.GetParentDomain(domainName)

	if parent == "" {
		return "", errors.New("invalid domain")
	}

	parentNamehash := e.namehash(parent)

	registrarAddress, err := e.ensRegistryContract.Owner(&bind.CallOpts{Pending: false}, parentNamehash)

	if err != nil || registrarAddress.Hex() == NullAddress {
		return "", err
	}

	return registrarAddress.Hex(), nil
}

//////////////////////////
// owner functions 		//
//////////////////////////

func (e EnsService) getOwnerFromNameWrapper(namehash common.Hash, ch chan<- ensGenericResult) {
	address, err := e.nameWrapperContract.OwnerOf(&bind.CallOpts{Pending: false}, namehash.Big())

	if err != nil || address.Hex() == NullAddress {
		ch <- ensGenericResult{nil, err, "nameWrapper"}
		return
	}
	ch <- ensGenericResult{result: address.Hex(), err: err, source: "nameWrapper"}
}

func (e EnsService) getOwnerFromRegistry(namehash common.Hash, ch chan<- ensGenericResult) {
	address, err := e.ensRegistryContract.Owner(&bind.CallOpts{Pending: false}, namehash)

	if err != nil || address.Hex() == NullAddress {
		ch <- ensGenericResult{nil, err, "registry"}
		return
	}
	ch <- ensGenericResult{result: address.Hex(), err: err, source: "registry"}
}

func (e EnsService) ownerOf(namehash common.Hash) (string, error) {

	ch := make(chan ensGenericResult, 2)

	go e.getOwnerFromNameWrapper(namehash, ch)
	go e.getOwnerFromRegistry(namehash, ch)

	address1 := <-ch
	address2 := <-ch

	var registryResult ensGenericResult
	var nameWrapperResult ensGenericResult

	if address1.source == "registry" {
		registryResult = address1
		nameWrapperResult = address2
	} else {
		registryResult = address2
		nameWrapperResult = address1
	}

	if nameWrapperResult.result != nil {
		return nameWrapperResult.result.(string), nil
	}

	if registryResult.result != nil {
		return registryResult.result.(string), nil
	}

	return "", nil
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

	// contentHash := registrarAddress.Hex()

	decode, err := utils.DecodeENSContentHash(hex.EncodeToString(contentHash))

	if err != nil {
		return "", err
	}

	return decode, nil
}

func (e EnsService) textRecord(resolverAddress string, namehash common.Hash, key string) (string, error) {
	resolverContract, err := resolverreader.NewContract(common.HexToAddress(resolverAddress), e.contractBackend)

	if err != nil {
		return "", err
	}

	return resolverContract.Text(&bind.CallOpts{Pending: false}, namehash, key)
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

func (e EnsService) isWrapped(namehash common.Hash) (bool, error) {
	return e.nameWrapperContract.IsWrapped0(&bind.CallOpts{Pending: false}, namehash)
}
