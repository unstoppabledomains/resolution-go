package resolution

import (
	"embed"
	"encoding/json"
)

//go:embed ens/ens-config.json
var ensConfigEmbed embed.FS
var ensConfigJSON, _ = ensConfigEmbed.ReadFile("ens/ens-config.json")

//go:embed ens/bip44.json
var bip44Embed embed.FS
var bip44Json, _ = bip44Embed.ReadFile("ens/bip44.json")

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

func newBip44Mapping() (map[string]int64, error) {
	var bip44Mapping map[string]int64

	err := json.Unmarshal(bip44Json, &bip44Mapping)
	if err != nil {
		return nil, err
	}

	return bip44Mapping, nil
}
