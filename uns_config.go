package resolution

import (
	"encoding/json"
	"embed"
)

// contracts struct of contracts
type contracts map[string]struct {
	Address         string
	Implementation  string
	LegacyAddresses []string
	DeploymentBlock string
}

// networks struct of contracts
type networks map[string]struct {
	Contracts contracts
}

const (
	Mainnet string = "mainnet"
	Polygon string = "polygon"
	Mumbai  string = "mumbai"
	Goerli  string = "goerli"
)

const (
	Layer1 string = "Layer 1"
	Layer2 string = "Layer 2"
)

type NetworkContracts map[string]contracts

var NetworkProviders = map[string]string{
	Mainnet: "https://mainnet.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e",
	Goerli:  "https://goerli.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e",
	Polygon: "https://polygon-mainnet.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e",
	Mumbai:  "https://polygon-mumbai.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e",
}

var NetworkNameToId = map[string]int{
	Mainnet: 1,
	Polygon: 137,
	Mumbai:  80001,
	Goerli:  5,
}

//go:embed uns/uns-config.json
var unsConfigEmbed embed.FS
var unsConfigJSON, _= unsConfigEmbed.ReadFile("uns/uns-config.json")

func parseAllContracts(data []byte) (networks, error) {
	var networksObject struct {
		Networks networks
	}
	err := json.Unmarshal(data, &networksObject)
	if err != nil {
		return nil, err
	}
	return networksObject.Networks, nil
}

func newContracts() (NetworkContracts, error) {
	networks := make(NetworkContracts)
	var err error
	net, err := parseAllContracts(unsConfigJSON)
	if err != nil {
		return nil, err
	}

	networks[Mainnet] = net["1"].Contracts
	networks[Polygon] = net["137"].Contracts
	networks[Mumbai] = net["80001"].Contracts
	networks[Goerli] = net["5"].Contracts
	return networks, nil
}
