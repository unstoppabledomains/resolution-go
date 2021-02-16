package resolution

import (
	"encoding/json"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"strings"
)

const znsDefaultProvider = "https://api.zilliqa.com"
const znsMainnetRegistry = "9611c53BE6d1b32058b2747bdeCECed7e1216793"
const znsContractField = "records"
const znsZeroAddress = "0x0000000000000000000000000000000000000000"

type ZnsProvider interface {
	GetSmartContractSubState(contractAddress string, params ...interface{}) (string, error)
}

type Zns struct {
	Provider ZnsProvider
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

type znsDomainState struct {
	Resolver string
	Owner    string
	Records  map[string]string
}

func NewZns(provider ZnsProvider) *Zns {
	return &Zns{Provider: provider}
}

func NewZnsWithDefaultProvider() *Zns {
	return &Zns{Provider: provider.NewProvider(znsDefaultProvider)}
}

func (z *Zns) State(domainName string, keys []string) (*znsDomainState, error) {
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

	response, err = z.Provider.GetSmartContractSubState(strings.TrimPrefix(resolver, "0x"), znsContractField, keys)
	if err != nil {
		return nil, err
	}
	var resolverState resolverSubState
	err = json.Unmarshal([]byte(response), &resolverState)
	if err != nil {
		return nil, err
	}
	records := resolverState.Result[znsContractField]

	return &znsDomainState{Owner: owner, Resolver: resolver, Records: records}, nil
}
