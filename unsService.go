package resolution

import (
	"encoding/json"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	kns "github.com/jgimeno/go-namehash"
	"github.com/unstoppabledomains/resolution-go/cns/contracts/resolver"
	"github.com/unstoppabledomains/resolution-go/dnsrecords"
	"github.com/unstoppabledomains/resolution-go/uns/contracts/proxyreader"
	"github.com/unstoppabledomains/resolution-go/uns/contracts/registry"
)

// Uns is a naming service handles Unstoppable domains resolution.
type UnsService struct {
	cnsDefaultResolver     common.Address
	unsRegistry            common.Address
	cnsStartingEventsBlock uint64
	unsStartingEventsBlock uint64
	proxyReader            *proxyreader.Contract
	supportedKeys          supportedKeys
	contractBackend        bind.ContractBackend
	metadataClient         MetadataClient
	Layer                  string
}

type MetadataClient interface {
	Get(url string) (resp *http.Response, err error)
}

var unsZeroAddress = common.HexToAddress("0x0")

// Data Get raw data attached to domain
func (c *UnsService) Data(domainName string, keys []string) (*struct {
	Resolver common.Address
	Owner    common.Address
	Values   []string
}, error) {
	normalizedName := normalizeName(domainName)
	namehash := kns.NameHash(normalizedName)
	tokenID := namehash.Big()
	data, err := c.proxyReader.GetData(&bind.CallOpts{Pending: false}, keys, tokenID)
	if err != nil {
		return nil, err
	}
	if data.Owner == unsZeroAddress {
		return nil, &DomainNotRegisteredError{DomainName: normalizedName}
	}
	if data.Resolver == unsZeroAddress {
		return nil, &DomainNotConfiguredError{DomainName: normalizedName, Layer: c.Layer}
	}

	return &data, nil
}

func (c *UnsService) Records(domainName string, keys []string) (map[string]string, error) {
	data, err := c.Data(domainName, keys)
	if err != nil {
		return nil, err
	}
	allRecords := make(map[string]string, len(keys))
	for index, key := range keys {
		allRecords[key] = data.Values[index]
	}
	return allRecords, nil
}

func (c *UnsService) Record(domainName string, key string) (string, error) {
	records, err := c.Records(domainName, []string{key})
	if err != nil {
		return "", err
	}
	return records[key], nil
}

func (c *UnsService) Addr(domainName string, ticker string) (string, error) {
	key, err := buildCryptoKey(ticker)
	if err != nil {
		return "", err
	}
	value, err := c.Record(domainName, key)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *UnsService) AddrVersion(domainName string, ticker string, version string) (string, error) {
	key, err := buildCryptoKeyVersion(ticker, version)
	if err != nil {
		return "", err
	}
	value, err := c.Record(domainName, key)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *UnsService) Email(domainName string) (string, error) {
	value, err := c.Record(domainName, emailKey)
	if err != nil {
		return "", err
	}

	return value, nil
}

func (c *UnsService) Resolver(domainName string) (string, error) {
	data, err := c.Data(domainName, []string{})
	if err != nil {
		return "", err
	}

	return data.Resolver.String(), nil
}

func (c *UnsService) Owner(domainName string) (string, error) {
	data, err := c.Data(domainName, []string{})
	if err != nil {
		return "", err
	}

	return data.Owner.String(), nil
}

func (c *UnsService) IpfsHash(domainName string) (string, error) {
	records, err := c.Records(domainName, ipfsKeys)
	if err != nil {
		return "", err
	}
	return returnFirstNonEmpty(records, ipfsKeys), nil
}

func (c *UnsService) HTTPUrl(domainName string) (string, error) {
	records, err := c.Records(domainName, redirectUrlKeys)
	if err != nil {
		return "", err
	}
	return returnFirstNonEmpty(records, redirectUrlKeys), nil
}

func (c *UnsService) getAllKeysFromContractEvents(contract *resolver.Contract, eventsStartingBlock uint64, domainName string) ([]string, error) {
	var allKeys []string
	normalizedName := normalizeName(domainName)
	namehash := kns.NameHash(normalizedName)
	resetRecordsIterator, err := contract.FilterResetRecords(&bind.FilterOpts{Start: eventsStartingBlock}, []*big.Int{namehash.Big()})
	if err != nil {
		return nil, err
	}
	newKeyEventsStartingBlock := eventsStartingBlock
	for resetRecordsIterator.Next() {
		if resetRecordsIterator.Error() != nil {
			return nil, err
		}
		newKeyEventsStartingBlock = resetRecordsIterator.Event.Raw.BlockNumber
	}
	newKeyIterator, err := contract.FilterNewKey(&bind.FilterOpts{Start: newKeyEventsStartingBlock}, []*big.Int{namehash.Big()}, []string{})
	if err != nil {
		return nil, err
	}
	for newKeyIterator.Next() {
		if newKeyIterator.Error() != nil {
			return nil, err
		}
		allKeys = append(allKeys, newKeyIterator.Event.Key)
	}
	if len(allKeys) == 0 {
		for key := range c.supportedKeys {
			allKeys = append(allKeys, key)
		}
	}
	return allKeys, err
}

func (c *UnsService) AllRecords(domainName string) (map[string]string, error) {
	data, err := c.Data(domainName, []string{})
	if err != nil {
		return nil, err
	}
	var allKeys []string
	if data.Resolver == c.cnsDefaultResolver || data.Resolver == c.unsRegistry {
		contract, err := resolver.NewContract(data.Resolver, c.contractBackend)
		if err != nil {
			return nil, err
		}
		eventsStartingBlock := c.cnsStartingEventsBlock
		if data.Resolver == c.unsRegistry {
			eventsStartingBlock = c.unsStartingEventsBlock
		}
		allKeys, err = c.getAllKeysFromContractEvents(contract, eventsStartingBlock, domainName)
		if err != nil {
			return nil, err
		}
	} else {
		for key := range c.supportedKeys {
			allKeys = append(allKeys, key)
		}
	}
	recordsData, err := c.Data(domainName, allKeys)
	if err != nil {
		return nil, err
	}
	allRecords := make(map[string]string)
	for index, key := range allKeys {
		if len(recordsData.Values[index]) > 0 {
			allRecords[key] = recordsData.Values[index]
		}
	}

	return allRecords, nil
}

func (c *UnsService) DNS(domainName string, types []dnsrecords.Type) ([]dnsrecords.Record, error) {
	keys, err := dnsTypesToCryptoRecordKeys(types)
	if err != nil {
		return nil, err
	}
	records, err := c.Records(domainName, keys)
	if err != nil {
		return nil, err
	}
	dnsRecords, err := cryptoRecordsToDNS(records)
	if err != nil {
		return nil, err
	}

	return dnsRecords, nil
}

func (c *UnsService) IsSupportedDomain(domainName string) (bool, error) {
	chunks := strings.Split(domainName, ".")
	if len(chunks) < 2 {
		return false, nil
	}
	extension := chunks[len(chunks)-1]
	if extension == "zil" {
		return false, nil
	}
	namehash := kns.NameHash(extension)
	tokenID := namehash.Big()
	data, err := c.proxyReader.Exists(&bind.CallOpts{Pending: false}, tokenID)
	if err != nil {
		return false, err
	}
	return data, nil
}

func (c *UnsService) TokenURI(domainName string) (string, error) {
	normalizedName := normalizeName(domainName)
	namehash := kns.NameHash(normalizedName)
	tokenUri, err := c.tokenUriByNamehash(namehash)
	if err != nil {
		return "", err
	}

	return tokenUri, nil
}

func (c *UnsService) TokenURIMetadata(domainName string) (TokenMetadata, error) {
	tokenUri, err := c.TokenURI(domainName)
	if err != nil {
		return TokenMetadata{}, err
	}
	metadata, err := c.tokenMetadataByUri(tokenUri)
	if err != nil {
		return TokenMetadata{}, err
	}
	return metadata, nil
}

func (c *UnsService) Unhash(domainHash string) (string, error) {
	namehash := common.HexToHash(domainHash)

	registryAddress, err := c.proxyReader.RegistryOf(&bind.CallOpts{}, namehash.Big())
	if err != nil {
		return "", err
	}
	eventsStartingBlock := c.cnsStartingEventsBlock
	if registryAddress == c.unsRegistry {
		eventsStartingBlock = c.unsStartingEventsBlock
	}
	domainName, _ := c.hashToNameFromNewURIEvents(namehash, registryAddress, eventsStartingBlock)

	if domainName == "" {
		return "", &DomainNotRegisteredError{Namehash: namehash.String()}
	}

	check, err := c.Namehash(domainName)
	if err != nil {
		return "", err
	}

	if common.HexToHash(check) != namehash {
		return "", &InvalidDomainNameReturnedError{Namehash: domainHash, DomainName: domainName}
	}

	return domainName, nil
}

func (c *UnsService) Namehash(domainName string) (string, error) {
	namehash := kns.NameHash(domainName)
	return namehash.String(), nil
}

func (c *UnsService) hashToNameFromNewURIEvents(namehash common.Hash, registryAddress common.Address, eventsStartingBlock uint64) (string, error) {
	registryContract, err := registry.NewContract(registryAddress, c.contractBackend)
	if err != nil {
		return "", err
	}
	newUriIterator, err := registryContract.FilterNewURI(&bind.FilterOpts{Start: eventsStartingBlock}, []*big.Int{namehash.Big()})
	if err != nil {
		return "", err
	}

	domainName := ""
	for newUriIterator.Next() {
		if newUriIterator.Error() != nil {
			return "", err
		}
		domainName = newUriIterator.Event.Uri
	}

	return domainName, nil
}

func (c *UnsService) tokenUriByNamehash(namehash common.Hash) (string, error) {
	tokenId := namehash.Big()
	tokenUri, err := c.proxyReader.TokenURI(&bind.CallOpts{Pending: false}, tokenId)
	if err != nil {
		if strings.HasPrefix(err.Error(), vm.ErrExecutionReverted.Error()) {
			return "", &DomainNotRegisteredError{Namehash: namehash.String()}
		}
		return "", err
	}

	return tokenUri, nil
}

func (c *UnsService) tokenMetadataByUri(tokenUri string) (TokenMetadata, error) {
	metadataResponse, err := c.metadataClient.Get(tokenUri)
	if err != nil {
		return TokenMetadata{}, err
	}
	defer metadataResponse.Body.Close()
	var returnedMetadata TokenMetadata
	err = json.NewDecoder(metadataResponse.Body).Decode(&returnedMetadata)
	if err != nil {
		return TokenMetadata{}, err
	}

	return returnedMetadata, nil
}
