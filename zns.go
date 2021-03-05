package resolution

import (
	"encoding/json"
	s "strings"

	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/unstoppabledomains/resolution-go/dnsrecords"
)

const znsDefaultProvider = "https://api.zilliqa.com"
const znsMainnetRegistry = "9611c53BE6d1b32058b2747bdeCECed7e1216793"
const znsContractField = "records"
const znsZeroAddress = "0x0000000000000000000000000000000000000000"

// Zns Zns
type Zns struct {
	Provider ZnsProvider
}

// ZnsProvider ZnsProvider
type ZnsProvider interface {
	GetSmartContractSubState(contractAddress string, params ...interface{}) (string, error)
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

// ZnsDomainState State of ZNS domain
type ZnsDomainState struct {
	Resolver string
	Owner    string
	Records  map[string]string
}

// NewZns Creates Zns instance
func NewZns(provider ZnsProvider) *Zns {
	return &Zns{Provider: provider}
}

// NewZnsWithDefaultProvider Creates instance of Zns with default provider
func NewZnsWithDefaultProvider() *Zns {
	return &Zns{Provider: provider.NewProvider(znsDefaultProvider)}
}

// State Retrieve the ZnsDomainState of a domain
func (z *Zns) State(domainName string) (*ZnsDomainState, error) {
	// todo validate domain name
	namehash, err := ZnsNameHash(domainName)
	if err != nil {
		return nil, err
	}
	response, err := z.Provider.GetSmartContractSubState(znsMainnetRegistry, znsContractField, []string{namehash})
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

	response, err = z.Provider.GetSmartContractSubState(s.TrimPrefix(resolver, "0x"), znsContractField, []string{})
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
	key := "crypto." + s.ToUpper(ticker) + ".address"
	value, err := z.Record(domainName, key)
	if err != nil {
		return "", err
	}
	return value, nil
}

// AddrVersion Retrieve the version value of domain's currency ticker - useful for multichain currencies
func (z *Zns) AddrVersion(domainName string, ticker string, version string) (string, error) {
	// todo replace concat by string builder
	key := "crypto." + s.ToUpper(ticker) + ".version." + s.ToUpper(version) + ".address"
	value, err := z.Record(domainName, key)
	if err != nil {
		return "", err
	}
	return value, nil
}

// Email Retrieve the email of a domain
func (z *Zns) Email(domainName string) (string, error) {
	key := "whois.email.value"
	value, err := z.Record(domainName, key)
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
	records, err := z.Records(domainName, []string{"dweb.ipfs.hash", "ipfs.html.value"})
	if err != nil {
		return "", err
	}
	if records["dweb.ipfs.hash"] != "" {
		return records["dweb.ipfs.hash"], nil
	}
	if records["ipfs.html.value"] != "" {
		return records["ipfs.html.value"], nil
	}

	return "", nil
}

// HTTPUrl Retrieve the http redirect url of a domain
func (z *Zns) HTTPUrl(domainName string) (string, error) {
	records, err := z.Records(domainName, []string{"browser.redirect_url", "ipfs.redirect_domain.value"})
	if err != nil {
		return "", err
	}
	if records["browser.redirect_url"] != "" {
		return records["browser.redirect_url"], nil
	}
	if records["ipfs.redirect_domain.value"] != "" {
		return records["ipfs.redirect_domain.value"], nil
	}

	return "", nil
}

// DNS Retrieve DNS records of domain
func (z *Zns) DNS(domainName string, types []dnsrecords.Type) ([]dnsrecords.Record, error) {
	keys, err := DNSTypesToCryptoRecordKeys(types)
	if err != nil {
		return nil, err
	}
	records, err := z.Records(domainName, keys)
	if err != nil {
		return nil, err
	}
	dnsRecords, err := CryptoRecordsToDNS(records)
	if err != nil {
		return nil, err
	}

	return dnsRecords, nil
}
