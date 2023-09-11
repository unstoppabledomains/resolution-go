package resolution

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
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
	return e.service.namehash(domainName).String(), nil
}

func (e *Ens) Resolver(domainName string) (string, error) {
	resolverAddress, err := e.service.resolver(e.service.namehash(domainName))

	if err != nil {
		return "", err
	}

	if resolverAddress == NullAddress {
		return "", &DomainNotConfiguredError{DomainName: domainName}
	}

	return resolverAddress, nil
}

func (e *Ens) ReverseOf(addr string) (string, error) {
	return e.service.reverseOf(addr)
}

func (e *Ens) Owner(domainName string) (string, error) {
	return e.service.ownerOf(e.service.namehash(domainName))
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

	resolverAddress, err := e.Resolver(domainName)

	if err != nil {
		return "", err
	}

	return e.service.addrCoinRecord(resolverAddress, e.service.namehash(domainName), big.NewInt(coinType))
}

func (e *Ens) CoinAddress(domainName string, coin string) (string, error) {

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

	resolverAddress, err := e.Resolver(domainName)

	if err != nil {
		return "", err
	}

	return e.service.addrCoinRecord(resolverAddress, e.service.namehash(domainName), coinNum)
}

func (e *Ens) ContentHash(domainName string) (string, error) {
	resolverAddress, err := e.Resolver(domainName)
	if err != nil {
		return "", err
	}

	return e.service.contenthashRecord(resolverAddress, e.service.namehash(domainName))
}

func (e *Ens) TextRecord(domainName, key string) (string, error) {
	resolverAddress, err := e.Resolver(domainName)

	if err != nil {
		return "", err
	}

	return e.service.textRecord(resolverAddress, e.service.namehash(domainName), key)
}

func (e *Ens) Records(domainName string, keys []string) (map[string]string, error) {
	return nil, &DomainNotSupportedError{DomainName: domainName}
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
	networkId := e.service.networkId
	result := make(map[string]namingservice.Location)

	for _, domainName := range domainNames {
		namehash := e.service.namehash(domainName)

		result[domainName] = namingservice.Location{
			NetworkId: networkId,
		}

		resolverAddress, err := e.service.resolver(namehash)

		if err != nil || resolverAddress == NullAddress {
			continue
		}

		owner, err := e.service.ownerOf(namehash)

		if err != nil || owner == NullAddress {
			continue
		}

		result[domainName] = namingservice.Location{
			NetworkId:             networkId,
			ResolverAddress:       resolverAddress,
			BlockchainProviderUrl: e.service.blockchainProviderUrl,
			Blockchain:            "ETH",
			OwnerAddress:          owner,
			RegistryAddress:       "0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85",
		}
	}

	return result, nil
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

	tokenUri, err := e.TokenURI(domainName)

	if err != nil {
		return TokenMetadata{}, err
	}

	metadataResponse, err := e.service.metadataClient.Get(tokenUri)

	if err != nil {
		return TokenMetadata{}, err
	}

	defer metadataResponse.Body.Close()

	var result TokenMetadata

	err = json.NewDecoder(metadataResponse.Body).Decode(&result)

	if err != nil {
		return TokenMetadata{}, err
	}

	for _, attr := range result.Attributes {
		switch v := attr.Value.(type) {
		case float64:
			fmt.Printf("Number value: %f\n", v)
		case string:
			fmt.Printf("String value: %s\n", v)
		case map[string]interface{}:
			fmt.Println("Object value:", v)
		}
	}

	if result.ExternalUrl == "" {
		result.ExternalUrl = fmt.Sprintf("https://unstoppabledomains.com/search?searchTerm=%s", domainName)
	}

	return result, nil
}

func (e *Ens) Unhash(domainHash string) (string, error) {
	client := e.service.metadataClient
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
	registrarContract := ensContracts[networkName]["BaseRegistrarImplementation"].Address

	ch := make(chan *http.Response, 2)

	go func() {
		resp, err := client.Get(fmt.Sprintf("https://metadata.ens.domains/%s/%s/%s", networkName, nameWrapContract, domainHash))
		if err != nil {
			ch <- nil
			return
		}
		ch <- resp
	}()

	go func() {
		fmt.Printf("https://metadata.ens.domains/%s/%s/%s\n", networkName, registrarContract, domainHash)
		resp, err := client.Get(fmt.Sprintf("https://metadata.ens.domains/%s/%s/%s", networkName, registrarContract, domainHash))
		if err != nil {
			ch <- nil
			return
		}
		ch <- resp
	}()

	for i := 0; i < 2; i++ {
		resp := <-ch
		if resp != nil && resp.StatusCode == 200 {
			defer resp.Body.Close()

			var result TokenMetadata

			err = json.NewDecoder(resp.Body).Decode(&result)

			if err != nil {
				return "", err
			}

			return result.Name, nil
		}
	}

	return "", nil
}
