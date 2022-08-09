package resolution

import (
	"encoding/json"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	kns "github.com/jgimeno/go-namehash"
	"github.com/unstoppabledomains/resolution-go/v2/dnsrecords"
	"github.com/unstoppabledomains/resolution-go/v2/namingservice"
	"github.com/unstoppabledomains/resolution-go/v2/uns/contracts/proxyreader"
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
	networkId              int
	blockchainProviderUrl  string
}

type MetadataClient interface {
	Get(url string) (resp *http.Response, err error)
}

var unsZeroAddress = common.HexToAddress("0x0")

func domainNameToTokenId(domainName string) *big.Int {
	normalizedName := normalizeName(domainName)
	namehash := kns.NameHash(normalizedName)
	return namehash.Big()
}

// Data Get raw data attached to domain
func (c *UnsService) data(domainName string, keys []string) (*struct {
	Resolver common.Address
	Owner    common.Address
	Values   []string
}, error) {
	normalizedDomain, err := normalizeAndVerifyName(domainName)
	if err != nil {
		return nil, err
	}
	tokenID := domainNameToTokenId(normalizedDomain)
	data, err := c.proxyReader.GetData(&bind.CallOpts{Pending: false}, keys, tokenID)
	if err != nil {
		return nil, err
	}
	if data.Owner == unsZeroAddress {
		return nil, &DomainNotRegisteredError{DomainName: domainName}
	}
	if data.Resolver == unsZeroAddress {
		return nil, &DomainNotConfiguredError{DomainName: domainName, Layer: c.Layer}
	}

	return &data, nil
}

func (c *UnsService) records(domainName string, keys []string) (map[string]string, error) {
	normalizedDomain, err := normalizeAndVerifyName(domainName)
	if err != nil {
		return nil, err
	}
	data, err := c.data(normalizedDomain, keys)
	if err != nil {
		return nil, err
	}
	allRecords := make(map[string]string, len(keys))
	for index, key := range keys {
		allRecords[key] = data.Values[index]
	}
	return allRecords, nil
}

func (c *UnsService) record(domainName string, key string) (string, error) {
	normalizedDomain, err := normalizeAndVerifyName(domainName)
	if err != nil {
		return "", err
	}
	records, err := c.records(normalizedDomain, []string{key})
	if err != nil {
		return "", err
	}
	return records[key], nil
}

func (c *UnsService) addr(domainName string, ticker string) (string, error) {
	normalizedDomain, err := normalizeAndVerifyName(domainName)
	if err != nil {
		return "", err
	}
	key, err := buildCryptoKey(ticker)
	if err != nil {
		return "", err
	}
	value, err := c.record(normalizedDomain, key)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *UnsService) addrVersion(domainName string, ticker string, version string) (string, error) {
	normalizedDomain, err := normalizeAndVerifyName(domainName)
	if err != nil {
		return "", err
	}
	key, err := buildCryptoKeyVersion(ticker, version)
	if err != nil {
		return "", err
	}
	value, err := c.record(normalizedDomain, key)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *UnsService) email(domainName string) (string, error) {
	normalizedDomain, err := normalizeAndVerifyName(domainName)
	if err != nil {
		return "", err
	}
	value, err := c.record(normalizedDomain, emailKey)
	if err != nil {
		return "", err
	}

	return value, nil
}

func (c *UnsService) resolver(domainName string) (string, error) {
	normalizedDomain, err := normalizeAndVerifyName(domainName)
	if err != nil {
		return "", err
	}
	data, err := c.data(normalizedDomain, []string{})
	if err != nil {
		return "", err
	}

	return data.Resolver.String(), nil
}

func (c *UnsService) owner(domainName string) (string, error) {
	normalizedDomain, err := normalizeAndVerifyName(domainName)
	if err != nil {
		return "", err
	}
	tokenID := domainNameToTokenId(normalizedDomain)
	data, err := c.proxyReader.GetData(&bind.CallOpts{Pending: false}, []string{}, tokenID)
	if err != nil {
		return "", err
	}
	if data.Owner == unsZeroAddress {
		return "", &DomainNotRegisteredError{DomainName: normalizedDomain}
	}
	return data.Owner.String(), nil
}

func (c *UnsService) ipfsHash(domainName string) (string, error) {
	normalizedDomain, err := normalizeAndVerifyName(domainName)
	if err != nil {
		return "", err
	}
	records, err := c.records(normalizedDomain, ipfsKeys)
	if err != nil {
		return "", err
	}
	return returnFirstNonEmpty(records, ipfsKeys), nil
}

func (c *UnsService) httpUrl(domainName string) (string, error) {
	normalizedDomain, err := normalizeAndVerifyName(domainName)
	if err != nil {
		return "", err
	}
	records, err := c.records(normalizedDomain, redirectUrlKeys)
	if err != nil {
		return "", err
	}
	return returnFirstNonEmpty(records, redirectUrlKeys), nil
}

func (c *UnsService) locations(domainNames []string) (map[string]namingservice.Location, error) {
	tokenIDs := make([]*big.Int, 0, len(domainNames))
	for _, domainName := range domainNames {
		normalizedDomain, err := normalizeAndVerifyName(domainName)
		if err != nil {
			return nil, err
		}
		namehash := kns.NameHash(normalizedDomain)
		tokenID := namehash.Big()
		tokenIDs = append(tokenIDs, tokenID)
	}
	var multicallInput [][]byte

	proxyReaderMethodsAbi, err := abi.JSON(strings.NewReader(proxyreader.ContractABI))
	if err != nil {
		return nil, err
	}

	for _, tokenId := range tokenIDs {
		packedRegistryOfFn, err := proxyReaderMethodsAbi.Pack("registryOf", tokenId)
		if err != nil {
			return nil, err
		}
		multicallInput = append(multicallInput, packedRegistryOfFn)
	}

	packedGetDataForManyFn, err := proxyReaderMethodsAbi.Pack("getDataForMany", []string{}, tokenIDs)
	if err != nil {
		return nil, err
	}

	multicallInput = append(multicallInput, packedGetDataForManyFn)
	ret, err := c.proxyReader.ContractCaller.Multicall(&bind.CallOpts{}, multicallInput)

	if err != nil {
		return nil, err
	}

	registryAddresses := []common.Address{}
	encodedRegistryAddresses := ret[:len(ret)-1]
	for _, encodedRegistryAddress := range encodedRegistryAddresses {
		decodedRegstryAddress := common.Address{}
		err := proxyReaderMethodsAbi.UnpackIntoInterface(&decodedRegstryAddress, "registryOf", encodedRegistryAddress)
		if err != nil {
			return nil, err
		}
		registryAddresses = append(registryAddresses, decodedRegstryAddress)
	}

	type GetDataForManyOutstruct struct {
		Resolvers []common.Address
		Owners    []common.Address
		Values    [][]string
	}
	dataOutstruct := GetDataForManyOutstruct{
		Resolvers: []common.Address{},
		Owners:    []common.Address{},
		Values:    [][]string{},
	}
	data := ret[len(ret)-1]
	err = proxyReaderMethodsAbi.UnpackIntoInterface(&dataOutstruct, "getDataForMany", data)
	if err != nil {
		return nil, err
	}

	locations := map[string]namingservice.Location{}
	for i, domainName := range domainNames {
		locations[domainName] = namingservice.Location{
			RegistryAddress:       (registryAddresses[i]).String(),
			ResolverAddress:       (dataOutstruct.Resolvers[i]).String(),
			OwnerAddress:          (dataOutstruct.Owners[i]).String(),
			NetworkId:             c.networkId,
			BlockchainProviderUrl: c.blockchainProviderUrl,
		}
	}

	return locations, nil
}

func (c *UnsService) DNS(domainName string, types []dnsrecords.Type) ([]dnsrecords.Record, error) {
	normalizedDomain, err := normalizeAndVerifyName(domainName)
	if err != nil {
		return nil, err
	}
	keys, err := dnsTypesToCryptoRecordKeys(types)
	if err != nil {
		return nil, err
	}
	records, err := c.records(normalizedDomain, keys)
	if err != nil {
		return nil, err
	}
	dnsRecords, err := cryptoRecordsToDNS(records)
	if err != nil {
		return nil, err
	}

	return dnsRecords, nil
}

func (c *UnsService) isSupportedDomain(domainName string) (bool, error) {
	chunks := strings.Split(domainName, ".")
	if len(chunks) < 2 {
		return false, nil
	}
	extension := chunks[len(chunks)-1]
	if extension == "zil" || extension == "coin" {
		return false, &DomainNotSupportedError{DomainName: domainName}
	}
	namehash := kns.NameHash(extension)
	tokenID := namehash.Big()
	data, err := c.proxyReader.Exists(&bind.CallOpts{Pending: false}, tokenID)
	if err != nil {
		return false, err
	}
	return data, nil
}

func (c *UnsService) tokenURI(domainName string) (string, error) {
	normalizedDomain, err := normalizeAndVerifyName(domainName)
	if err != nil {
		return "", err
	}
	namehash := kns.NameHash(normalizedDomain)
	return c.tokenUriByNamehash(namehash)
}

func (c *UnsService) tokenURIMetadata(domainName string) (TokenMetadata, error) {
	normalizedDomain, err := normalizeAndVerifyName(domainName)
	if err != nil {
		return TokenMetadata{}, err
	}
	namehash := kns.NameHash(normalizedDomain)
	return c.tokenURIMetadataByNamehash(namehash)
}

func (c *UnsService) tokenURIMetadataByNamehash(namehash common.Hash) (TokenMetadata, error) {
	tokenUri, err := c.tokenUriByNamehash(namehash)
	if err != nil {
		return TokenMetadata{}, err
	}
	metadata, err := c.tokenMetadataByUri(tokenUri)
	if err != nil {
		return TokenMetadata{}, err
	}
	return metadata, nil
}

func (c *UnsService) unhash(domainHash string) (string, error) {
	namehash := common.HexToHash(domainHash)

	metadata, err := c.tokenURIMetadataByNamehash(namehash)

	if err != nil {
		return "", err
	}

	domainName := metadata.Name
	if domainName == "" {
		return "", &DomainNotRegisteredError{Namehash: namehash.String()}
	}

	check, err := c.namehash(domainName)
	if err != nil {
		return "", err
	}

	if common.HexToHash(check) != namehash {
		return "", &InvalidDomainNameReturnedError{Namehash: domainHash, DomainName: domainName}
	}

	normalizedDomain, err := normalizeAndVerifyName(domainName)
	if err != nil {
		return "", err
	}

	return normalizedDomain, nil
}

func (c *UnsService) namehash(domainName string) (string, error) {
	normalizedDomain, err := normalizeAndVerifyName(domainName)
	if err != nil {
		return "", err
	}
	namehash := kns.NameHash(normalizedDomain)
	return namehash.String(), nil
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
