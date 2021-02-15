package resolution

import (
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
)

const znsDefaultProvider = "https://api.zilliqa.com"
const znsMainnetRegistry = "9611c53BE6d1b32058b2747bdeCECed7e1216793"
const znsContractField = "records"

type ZnsProvider interface {
	GetSmartContractSubState(contractAddress string, params ...interface{}) (string, error)
}

type Zns struct {
	Provider ZnsProvider
}

func NewZns(provider ZnsProvider) *Zns {
	return &Zns{Provider: provider}
}

func NewZnsWithDefaultProvider() *Zns {
	return &Zns{Provider: provider.NewProvider(znsDefaultProvider)}
}

func (z *Zns) State(domainName string, keys []string) {
	namehash, err := ZnsNameHash(domainName)
	if err != nil {
		return
	}
	state, err := z.Provider.GetSmartContractSubState(znsMainnetRegistry, znsContractField, []string{namehash})
	if err != nil {
		// todo continue with resolver query
		fmt.Println(err)
	} else {
		fmt.Println(state)
	}
}
