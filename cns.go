package resolution

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	kns "github.com/jgimeno/go-namehash"
	"github.com/unstoppabledomains/resolution-go/cns/contracts/proxyreader"
	"github.com/unstoppabledomains/resolution-go/cns/contracts/resolver"
	"github.com/unstoppabledomains/resolution-go/dnsrecords"
	"math/big"
	"strings"
)

type Cns struct {
	proxyReader     *proxyreader.Contract
	supportedKeys   supportedKeys
	contractBackend bind.ContractBackend
}

const cnsProvider = "https://mainnet.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e"
const cnsEventsStartingBlock uint64 = 9923764

var cnsZeroAddress = common.HexToAddress("0x0")
var cnsMainnetProxyReader = common.HexToAddress("0xa6E7cEf2EDDEA66352Fd68E5915b60BDbb7309f5")
var cnsMainnetDefaultResolver = common.HexToAddress("0xb66DcE2DA6afAAa98F2013446dBCB0f4B0ab2842")

// NewCns Creates instance of Cns with specific provider
func NewCns(backend bind.ContractBackend) (*Cns, error) {
	contract, err := proxyreader.NewContract(cnsMainnetProxyReader, backend)
	if err != nil {
		return nil, err
	}
	supportedKeys, err := newSupportedKeys()
	if err != nil {
		return nil, err
	}

	return &Cns{proxyReader: contract, supportedKeys: supportedKeys, contractBackend: backend}, nil
}

// NewCnsWithDefaultBackend Creates instance of Cns with default provider
func NewCnsWithDefaultBackend() (*Cns, error) {
	backend, err := ethclient.Dial(cnsProvider)
	if err != nil {
		return nil, err
	}
	cns, err := NewCns(backend)
	if err != nil {
		return nil, err
	}

	return cns, nil
}

// Data Retrieve data of domain
func (c *Cns) Data(domainName string, keys []string) (*struct {
	Resolver common.Address
	Owner    common.Address
	Values   []string
}, error) {
	normalizedName := normalizeName(domainName)
	if !c.IsSupportedDomain(normalizedName) {
		return nil, &DomainNotSupported{DomainName: normalizedName}
	}
	namehash := kns.NameHash(normalizedName)
	tokenID := namehash.Big()
	data, err := c.proxyReader.GetData(&bind.CallOpts{Pending: false}, keys, tokenID)
	if err != nil {
		return nil, err
	}
	if data.Owner == cnsZeroAddress {
		return nil, &DomainNotRegistered{DomainName: normalizedName}
	}
	if data.Resolver == cnsZeroAddress {
		return nil, &DomainNotConfigured{DomainName: normalizedName}
	}

	return &data, nil
}

// Records retrieve records of domain
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

// Record Retrieve single record of domain
func (c *Cns) Record(domainName string, key string) (string, error) {
	records, err := c.Records(domainName, []string{key})
	if err != nil {
		return "", err
	}
	return records[key], nil
}

// Addr Retrieve the value of domain's currency ticker
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

// AddrVersion Retrieve the version value of domain's currency ticker - useful for multichain currencies
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

// Email Retrieve the email of domain
func (c *Cns) Email(domainName string) (string, error) {
	value, err := c.Record(domainName, emailKey)
	if err != nil {
		return "", err
	}

	return value, nil
}

// Resolver Retrieve the resolver set for a domain
func (c *Cns) Resolver(domainName string) (string, error) {
	data, err := c.Data(domainName, []string{})
	if err != nil {
		return "", err
	}

	return data.Resolver.String(), nil
}

// Owner Retrieve the owner of a domain
func (c *Cns) Owner(domainName string) (string, error) {
	data, err := c.Data(domainName, []string{})
	if err != nil {
		return "", err
	}

	return data.Owner.String(), nil
}

// IpfsHash Retrieve the ipfs hash of a domain
func (c *Cns) IpfsHash(domainName string) (string, error) {
	records, err := c.Records(domainName, ipfsKeys)
	if err != nil {
		return "", err
	}
	return returnFirstNonEmpty(records, ipfsKeys), nil
}

// HTTPUrl Retrieve the http redirect url of a domain
func (c *Cns) HTTPUrl(domainName string) (string, error) {
	records, err := c.Records(domainName, redirectUrlKeys)
	if err != nil {
		return "", err
	}
	return returnFirstNonEmpty(records, redirectUrlKeys), nil
}

// AllRecords Retrieve all records of a domain
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

// DNS Retrieve the DNS records of a domain
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

// IsSupportedDomain checks whether domain name is supported by the naming service
func (c *Cns) IsSupportedDomain(domainName string) bool {
	return strings.HasSuffix(domainName, ".crypto")
}
