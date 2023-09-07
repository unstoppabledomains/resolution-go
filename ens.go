package resolution

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/unstoppabledomains/resolution-go/v3/dnsrecords"
	"github.com/unstoppabledomains/resolution-go/v3/namingservice"
	"github.com/unstoppabledomains/resolution-go/v3/utils"
)

// Ens is a naming service handles Ethereum naming service resolution.
type Ens struct {
	service EnsService
}

func (e *Ens) IsSupportedDomain(domainName string) (bool, error) {
	if domainName == "" {
		return false, nil
	}

	_, extension := utils.SplitDomain(domainName)

	return e.service.domainExists(e.service.namehash(extension))
}

func (e *Ens) Namehash(domainName string) (string, error) {
	normalizedName := utils.NormalizeName(domainName)
	return e.service.namehash(normalizedName).String(), nil
}

func (e *Ens) Resolver(domainName string) (string, error) {
	normalizedName := utils.NormalizeName(domainName)
	resolverAddress, err := e.service.resolver(e.service.namehash(normalizedName))

	if err != nil {
		return "", err
	}

	if resolverAddress == NullAddress {
		return "", &DomainNotConfiguredError{DomainName: normalizedName}
	}

	return resolverAddress, nil
}

func (e *Ens) ReverseOf(addr string) (string, error) {
	return e.service.reverseOf(addr)
}

func (e *Ens) Owner(domainName string) (string, error) {
	normalizedName := utils.NormalizeName(domainName)
	return e.service.ownerOf(e.service.namehash(normalizedName))
}

func (e *Ens) Addr(domainName, ticker string) (string, error) {
	bip44, err := newBip44Mapping()

	if err != nil {
		return "", err
	}

	coinType, ok := bip44[ticker]

	if !ok {
		return "", &EnsInvalidCoinType{CoinType: ticker}
	}

	normalizedName := utils.NormalizeName(domainName)

	resolverAddress, err := e.Resolver(domainName)

	if err != nil {
		return "", err
	}

	return e.service.addrCoinRecord(resolverAddress, e.service.namehash(normalizedName), big.NewInt(coinType))
}

func (e *Ens) CoinAddress(domainName string, coin string) (string, error) {
	normalizedName := utils.NormalizeName(domainName)

	var coinNum *big.Int
	if strings.HasPrefix(coin, "0x8") { // hexadecimal representation
		num, err := strconv.ParseInt(coin[3:], 16, 64)

		if err != nil {
			return "", err
		}

		coinNum = big.NewInt(num)
	} else {
		coinNum = new(big.Int)
		_, ok := coinNum.SetString(coin, 10)

		if !ok {
			return "", &EnsInvalidCoinType{CoinType: coin}
		}
	}

	resolverAddress, err := e.Resolver(normalizedName)

	if err != nil {
		return "", err
	}

	return e.service.addrCoinRecord(resolverAddress, e.service.namehash(normalizedName), coinNum)
}

func (e *Ens) ContentHash(domainName string) (string, error) {
	normalizedName := utils.NormalizeName(domainName)
	resolverAddress, err := e.Resolver(normalizedName)
	if err != nil {
		return "", err
	}

	return e.service.contenthashRecord(resolverAddress, e.service.namehash(normalizedName))
}

func (e *Ens) TextRecord(domainName, key string) (string, error) {
	normalizedName := utils.NormalizeName(domainName)
	resolverAddress, err := e.Resolver(normalizedName)

	if err != nil {
		return "", err
	}

	return e.service.textRecord(resolverAddress, e.service.namehash(normalizedName), key)
}

func (e *Ens) Records(domainName string, keys []string) (map[string]string, error) {
	return nil, nil
}

func (e *Ens) Record(domainName string, key string) (string, error) {
	return "", nil
}

func (e *Ens) AddrVersion(domainName string, ticker string, version string) (string, error) {
	return "", nil
}

func (e *Ens) Email(domainName string) (string, error) {
	namehash := e.service.namehash(domainName)
	resolverAddress, err := e.service.resolver(namehash)

	if err != nil {
		return "", err
	}

	if resolverAddress == NullAddress {
		return "", &DomainNotConfiguredError{DomainName: domainName}
	}

	return e.service.textRecord(resolverAddress, namehash, "email")
}

// IpfsHash Retrieve hash of IPFS website attached to domain.
func (e *Ens) IpfsHash(domainName string) (string, error) {
	namehash := e.service.namehash(domainName)
	resolverAddress, err := e.service.resolver(namehash)

	if err != nil {
		return "", err
	}

	if resolverAddress == NullAddress {
		return "", &DomainNotConfiguredError{DomainName: domainName}
	}

	return e.service.textRecord(resolverAddress, namehash, "avatar")
}

// HTTPUrl Retrieve the http redirect url of a domain.
func (e *Ens) HTTPUrl(domainName string) (string, error) {
	namehash := e.service.namehash(domainName)
	resolverAddress, err := e.service.resolver(namehash)

	if err != nil {
		return "", err
	}

	if resolverAddress == NullAddress {
		return "", &DomainNotConfiguredError{DomainName: domainName}
	}

	return e.service.textRecord(resolverAddress, namehash, "url")
}

// AllRecords Retrieve all records of a domain.
// Returns result in string or empty string record is not found.
// Deprecated: This method will be removed in future releases
func (e *Ens) AllRecords(domainName string) (map[string]string, error) {
	return nil, nil
}

// Locations Retrieve locations of domains
// Returns key-value map of domain names to location
func (e *Ens) Locations(domainNames []string) (map[string]namingservice.Location, error) {
	return make(map[string]namingservice.Location), nil
}

// DNS Retrieve the DNS records of a domain.
// Returns a set of valid and filtered non-empty DNS records attached to domain.
func (e *Ens) DNS(domainName string, types []dnsrecords.Type) ([]dnsrecords.Record, error) {
	return nil, nil
}

// TokenURI returns ERC721 metadata token URI
func (e *Ens) TokenURI(domainName string) (string, error) {
	namehash := e.service.namehash(domainName)

	isWrapped, err := e.service.isWrapped(namehash)

	if err != nil {
		return "", err
	}

	networkId := e.service.networkId
	ensContracts, err := newEnsContracts()

	if err != nil {
		return "", err
	}

	var networkName = Mainnet

	if networkId == 5 {
		networkName = Goerli
	}

	nameWrapContract := ensContracts[networkName]["NameWrapper"].Address

	if isWrapped {
		return fmt.Sprintf("https://metadata.ens.domains/%s/%s/%s", networkName, nameWrapContract, namehash), nil
	}

	erc721Hash := e.service.labelNamehash(domainName)

	registrarContract := ensContracts[networkName]["BaseRegistrarImplementation"].Address

	return fmt.Sprintf("https://metadata.ens.domains/%s/%s/%s", networkName, registrarContract, erc721Hash), nil
}

// TokenURIMetadata returns ERC721 metadata
func (e *Ens) TokenURIMetadata(domainName string) (TokenMetadata, error) {
	return TokenMetadata{}, nil
}

func (e *Ens) Unhash(domainHash string) (string, error) {
	return "", nil
}
