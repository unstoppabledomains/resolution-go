package resolution

import (
	"encoding/json"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethclient"
	kns "github.com/jgimeno/go-namehash"
	"github.com/unstoppabledomains/resolution-go/cns/contracts/proxyreader"
	"github.com/unstoppabledomains/resolution-go/cns/contracts/resolver"
	"github.com/unstoppabledomains/resolution-go/dnsrecords"
)

// Cns is a naming service handles .crypto domains resolution.
type Cns struct {
	proxyReader     *proxyreader.Contract
	supportedKeys   supportedKeys
	contractBackend bind.ContractBackend
	metadataClient  MetadataClient
}

// CnsBuilder is a builder to setup and build instance of Cns service.
type CnsBuilder interface {
	// SetContractBackend set Ethereum backend for communication with CNS registry
	SetContractBackend(backend bind.ContractBackend) CnsBuilder

	// SetMetadataClient set http backend for communication with ERC721 metadata server
	SetMetadataClient(backend MetadataClient) CnsBuilder

	// Build Cns instance
	Build() (*Cns, error)
}

type cnsBuilder struct {
	contractBackend bind.ContractBackend
	metadataClient  MetadataClient
}

type MetadataClient interface {
	Get(url string) (resp *http.Response, err error)
}

const cnsProvider = "https://mainnet.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e"
const cnsEventsStartingBlock uint64 = 9923764

var cnsZeroAddress = common.HexToAddress("0x0")
var cnsMainnetProxyReader = common.HexToAddress("0xa6E7cEf2EDDEA66352Fd68E5915b60BDbb7309f5")
var cnsMainnetDefaultResolver = common.HexToAddress("0xb66DcE2DA6afAAa98F2013446dBCB0f4B0ab2842")

// NewCnsBuilder Creates builder to setup new instance os Cns service.
func NewCnsBuilder() CnsBuilder {
	return &cnsBuilder{}
}

// SetContractBackend set Ethereum backend for communication with CNS registry
func (cb *cnsBuilder) SetContractBackend(backend bind.ContractBackend) CnsBuilder {
	cb.contractBackend = backend
	return cb
}

func (cb *cnsBuilder) SetMetadataClient(client MetadataClient) CnsBuilder {
	cb.metadataClient = client
	return cb
}

// Build Cns instance
func (cb *cnsBuilder) Build() (*Cns, error) {
	if cb.contractBackend == nil {
		backend, err := ethclient.Dial(cnsProvider)
		if err != nil {
			return nil, err
		}
		cb.contractBackend = backend
	}
	if cb.metadataClient == nil {
		cb.metadataClient = &http.Client{}
	}
	proxyReaderContract, err := proxyreader.NewContract(cnsMainnetProxyReader, cb.contractBackend)
	if err != nil {
		return nil, err
	}
	supportedKeys, err := newSupportedKeys()
	if err != nil {
		return nil, err
	}
	return &Cns{proxyReader: proxyReaderContract, supportedKeys: supportedKeys, contractBackend: cb.contractBackend, metadataClient: cb.metadataClient}, nil
}

// Data Get raw data attached to domain
func (c *Cns) Data(domainName string, keys []string) (*struct {
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
	if data.Owner == cnsZeroAddress {
		return nil, &DomainNotRegisteredError{DomainName: normalizedName}
	}
	if data.Resolver == cnsZeroAddress {
		return nil, &DomainNotConfiguredError{DomainName: normalizedName}
	}

	return &data, nil
}

func (c *Cns) Records(domainName string, keys []string) (map[string]string, error) {
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

func (c *Cns) Record(domainName string, key string) (string, error) {
	records, err := c.Records(domainName, []string{key})
	if err != nil {
		return "", err
	}
	return records[key], nil
}

func (c *Cns) Addr(domainName string, ticker string) (string, error) {
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

func (c *Cns) AddrVersion(domainName string, ticker string, version string) (string, error) {
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

func (c *Cns) Email(domainName string) (string, error) {
	value, err := c.Record(domainName, emailKey)
	if err != nil {
		return "", err
	}

	return value, nil
}

func (c *Cns) Resolver(domainName string) (string, error) {
	data, err := c.Data(domainName, []string{})
	if err != nil {
		return "", err
	}

	return data.Resolver.String(), nil
}

func (c *Cns) Owner(domainName string) (string, error) {
	data, err := c.Data(domainName, []string{})
	if err != nil {
		return "", err
	}

	return data.Owner.String(), nil
}

func (c *Cns) IpfsHash(domainName string) (string, error) {
	records, err := c.Records(domainName, ipfsKeys)
	if err != nil {
		return "", err
	}
	return returnFirstNonEmpty(records, ipfsKeys), nil
}

func (c *Cns) HTTPUrl(domainName string) (string, error) {
	records, err := c.Records(domainName, redirectUrlKeys)
	if err != nil {
		return "", err
	}
	return returnFirstNonEmpty(records, redirectUrlKeys), nil
}

func (c *Cns) AllRecords(domainName string) (map[string]string, error) {
	data, err := c.Data(domainName, []string{})
	if err != nil {
		return nil, err
	}
	var allKeys []string
	if data.Resolver == cnsMainnetDefaultResolver {
		resolverContract, err := resolver.NewContract(data.Resolver, c.contractBackend)
		if err != nil {
			return nil, err
		}
		normalizedName := normalizeName(domainName)
		namehash := kns.NameHash(normalizedName)
		resetRecordsIterator, err := resolverContract.FilterResetRecords(&bind.FilterOpts{Start: cnsEventsStartingBlock}, []*big.Int{namehash.Big()})
		if err != nil {
			return nil, err
		}
		newKeyEventsStartingBlock := cnsEventsStartingBlock
		for resetRecordsIterator.Next() {
			if resetRecordsIterator.Error() != nil {
				return nil, err
			}
			newKeyEventsStartingBlock = resetRecordsIterator.Event.Raw.BlockNumber
		}
		newKeyIterator, err := resolverContract.FilterNewKey(&bind.FilterOpts{Start: newKeyEventsStartingBlock}, []*big.Int{namehash.Big()}, []string{})
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

func (c *Cns) DNS(domainName string, types []dnsrecords.Type) ([]dnsrecords.Record, error) {
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

func (c *Cns) IsSupportedDomain(domainName string) (bool, error) {
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

func (c *Cns) TokenURI(domainName string) (string, error) {
	normalizedName := normalizeName(domainName)
	namehash := kns.NameHash(normalizedName)
	tokenUri, err := c.tokenUriByNamehash(namehash)
	if err != nil {
		return "", err
	}

	return tokenUri, nil
}

func (c *Cns) TokenURIMetadata(domainName string) (TokenMetadata, error) {
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

func (c *Cns) Unhash(domainHash string) (string, error) {
	namehash := common.HexToHash(domainHash)
	tokenUri, err := c.tokenUriByNamehash(namehash)
	if err != nil {
		return "", err
	}
	metadata, err := c.tokenMetadataByUri(tokenUri)
	if err != nil {
		return "", err
	}
	domainName := normalizeName(metadata.Name)
	expectedNamehash := kns.NameHash(domainName)
	if namehash != expectedNamehash {
		return "", &InvalidDomainNameReturnedError{
			DomainName: domainName,
			Namehash:   domainHash,
		}
	}

	return domainName, nil
}

func (c *Cns) tokenUriByNamehash(namehash common.Hash) (string, error) {
	tokenId := namehash.Big()
	tokenUri, err := c.proxyReader.TokenURI(&bind.CallOpts{Pending: false}, tokenId)
	if err != nil {
		if err.Error() == vm.ErrExecutionReverted.Error() {
			return "", &DomainNotRegisteredError{Namehash: namehash.String()}
		}
		return "", err
	}

	return tokenUri, nil
}

func (c *Cns) tokenMetadataByUri(tokenUri string) (TokenMetadata, error) {
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
