package resolution

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	kns "github.com/jgimeno/go-namehash"
	"github.com/unstoppabledomains/resolution-go/v3/ens/contracts/ensreader"
)

// Ens is a naming service
type EnsService struct {
	ensReaderContract     *ensreader.Contract
	contractBackend       bind.ContractBackend
	networkId             int
	blockchainProviderUrl string
}

func (e *EnsService) owner(domainName string) (string, error) {
	normalizedName := normalizeName(domainName)
	namehash := kns.NameHash(normalizedName)

	address, err := e.ensReaderContract.Owner(&bind.CallOpts{Pending: false}, namehash)
	if err != nil {
		return "", err
	}

	return address.Hex(), nil
}

func (e *EnsService) resolver(domainName string) (string, error) {
	normalizedName := normalizeName(domainName)
	namehash := kns.NameHash(normalizedName)

	resolverAddress, err := e.ensReaderContract.Resolver(&bind.CallOpts{Pending: false}, namehash)

	if err != nil {
		return "", err
	}

	return resolverAddress.Hex(), nil
}

func (e *EnsService) namehash(domainName string) (string, error) {
	namehash := kns.NameHash(domainName)
	return namehash.String(), nil
}
