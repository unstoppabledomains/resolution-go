package resolution

import (
	"embed"
	"encoding/json"
)

//go:embed ens/ens-config.json
var ensConfigEmbed embed.FS
var ensConfigJSON, _ = ensConfigEmbed.ReadFile("ens/ens-config.json")

func parseAllEnsContracts(data []byte) (networks, error) {
	var networksObject struct {
		Networks networks
	}
	err := json.Unmarshal(data, &networksObject)
	if err != nil {
		return nil, err
	}
	return networksObject.Networks, nil
}

func newEnsContracts() (NetworkContracts, error) {
	networks := make(NetworkContracts)
	var err error
	net, err := parseAllEnsContracts(ensConfigJSON)
	if err != nil {
		return nil, err
	}

	networks[Mainnet] = net["1"].Contracts
	networks[Goerli] = net["5"].Contracts
	return networks, nil
}
