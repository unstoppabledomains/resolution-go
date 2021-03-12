package resolution

import (
	"encoding/json"
	"strings"

	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/unstoppabledomains/resolution-go/dnsrecords"
)

// Zns Zns
type Zns struct {
	provider ZnsProvider
}

// ZnsProvider ZnsProvider
type ZnsProvider interface {
	GetSmartContractSubState(contractAddress string, params ...interface{}) (string, error)
}

// ZnsDomainState State of ZNS domain
type ZnsDomainState struct {
	Resolver string
	Owner    string
	Records  map[string]string
}

type registrySubState struct {
	Result map[string]map[string]struct {
		Arguments   []string
		Argtypes    []string
		Constructor string
	}
}

type resolverSubState struct {
	Result map[string]map[string]string
}

const znsDefaultProvider = "https://api.zilliqa.com"
const znsMainnetRegistry = "9611c53BE6d1b32058b2747bdeCECed7e1216793"
const znsContractField = "records"
const znsZeroAddress = "0x0000000000000000000000000000000000000000"

// NewZns Creates Zns instance
func NewZns(provider ZnsProvider) *Zns {
	return &Zns{provider: provider}
}

// NewZnsWithDefaultProvider Creates instance of Zns with default provider
func NewZnsWithDefaultProvider() *Zns {
	return &Zns{provider: provider.NewProvider(znsDefaultProvider)}
}

// State Retrieve the ZnsDomainState of a domain
func (z *Zns) State(domainName string) (*ZnsDomainState, error) {
	normalizedName := normalizeName(domainName)
	if !z.IsSupportedDomain(normalizedName) {
		return nil, &DomainNotSupported{DomainName: normalizedName}
	}
	namehash, err := ZnsNameHash(domainName)
	if err != nil {
		return nil, err
	}
	response, err := z.provider.GetSmartContractSubState(znsMainnetRegistry, znsContractField, []string{namehash})
	if err != nil {
		return nil, err
	}

	var registryState registrySubState
	err = json.Unmarshal([]byte(response), &registryState)
	if err != nil {
		return nil, err
	}
	registryValues := registryState.Result[znsContractField][namehash].Arguments
	if len(registryValues) == 0 {
		return nil, &DomainNotRegistered{DomainName: domainName}
	}
	owner, resolver := registryValues[0], registryValues[1]
	if owner == znsZeroAddress {
		return nil, &DomainNotRegistered{DomainName: domainName}
	}
	if resolver == znsZeroAddress {
		return nil, &DomainNotConfigured{DomainName: domainName}
	}

	response, err = z.provider.GetSmartContractSubState(strings.TrimPrefix(resolver, "0x"), znsContractField, []string{})
	if err != nil {
		return nil, err
	}
	var resolverState resolverSubState
	err = json.Unmarshal([]byte(response), &resolverState)
	if err != nil {
		return nil, err
	}
	records := resolverState.Result[znsContractField]

	return &ZnsDomainState{Owner: owner, Resolver: resolver, Records: records}, nil
}

// Records Retrieve the records of a domain
func (z *Zns) Records(domainName string, keys []string) (map[string]string, error) {
	state, err := z.State(domainName)
	if err != nil {
		return nil, err
	}
	records := make(map[string]string, len(keys))
	for _, recordKey := range keys {
		records[recordKey] = state.Records[recordKey]
	}

	return records, err
}

// Record Retrieve a single record of a domain
func (z *Zns) Record(domainName string, key string) (string, error) {
	records, err := z.Records(domainName, []string{key})
	if err != nil {
		return "", nil
	}
	return records[key], nil
}

// Owner Retrieve the owner of a domain
func (z *Zns) Owner(domainName string) (string, error) {
	state, err := z.State(domainName)
	if err != nil {
		return "", err
	}

	return state.Owner, err
}

// Resolver Retrieve the resolver set for a domain
func (z *Zns) Resolver(domainName string) (string, error) {
	state, err := z.State(domainName)
	if err != nil {
		return "", err
	}

	return state.Resolver, err
}

// Addr Retrieve the value of domain's currency ticker
func (z *Zns) Addr(domainName string, ticker string) (string, error) {
	key, err := buildCryptoKey(ticker)
	if err != nil {
		return "", err
	}
	value, err := z.Record(domainName, key)
	if err != nil {
		return "", err
	}
	return value, nil
}

// AddrVersion Retrieve the version value of domain's currency ticker - useful for multichain currencies
func (z *Zns) AddrVersion(domainName string, ticker string, version string) (string, error) {
	key, err := buildCryptoKeyVersion(ticker, version)
	if err != nil {
		return "", err
	}
	value, err := z.Record(domainName, key)
	if err != nil {
		return "", err
	}
	return value, nil
}

// Email Retrieve the email of a domain
func (z *Zns) Email(domainName string) (string, error) {
	value, err := z.Record(domainName, emailKey)
	if err != nil {
		return "", err
	}

	return value, nil
}

// AllRecords Retrieve the all records of a domain
func (z *Zns) AllRecords(domainName string) (map[string]string, error) {
	state, err := z.State(domainName)
	if err != nil {
		return nil, err
	}

	return state.Records, err
}

// IpfsHash Retrieve the ipfs hash of a domain
func (z *Zns) IpfsHash(domainName string) (string, error) {
	records, err := z.Records(domainName, ipfsKeys)
	if err != nil {
		return "", err
	}
	return returnFirstNonEmpty(records, ipfsKeys), nil
}

// HTTPUrl Retrieve the http redirect url of a domain
func (z *Zns) HTTPUrl(domainName string) (string, error) {
	records, err := z.Records(domainName, redirectUrlKeys)
	if err != nil {
		return "", err
	}
	return returnFirstNonEmpty(records, redirectUrlKeys), nil
}

// DNS Retrieve DNS records of domain
func (z *Zns) DNS(domainName string, types []dnsrecords.Type) ([]dnsrecords.Record, error) {
	keys, err := dnsTypesToCryptoRecordKeys(types)
	if err != nil {
		return nil, err
	}
	records, err := z.Records(domainName, keys)
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
func (z *Zns) IsSupportedDomain(domainName string) bool {
	return strings.HasSuffix(domainName, ".zil")
}
