package resolution

import (
	"encoding/json"
)

// contracts struct of contracts
type contracts map[string]struct {
	Address         string
	Implementation  string
	LegacyAddresses []string
	DeploymentBlock string
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

func parseContracts(data []byte) (contracts, error) {
	var contractsObject struct {
		Contracts contracts
	}
	err := json.Unmarshal(data, &contractsObject)
	if err != nil {
		return nil, err
	}
	return contractsObject.Contracts, nil
}

func newContracts() (NetworkContracts, error) {
	networks := make(NetworkContracts)
	var err error
	networks[Mainnet], err = parseContracts(unsMainnetConfigJSON)
	if err != nil {
		return nil, err
	}
	networks[Polygon], err = parseContracts(unsPolygonConfigJSON)
	if err != nil {
		return nil, err
	}
	networks[Mumbai], err = parseContracts(unsMumbaiConfigJSON)
	if err != nil {
		return nil, err
	}
	networks[Goerli], err = parseContracts(unsGoerliConfigJSON)
	if err != nil {
		return nil, err
	}
	return networks, nil
}

//go:embed uns/uns-config.json
var f embed.FS
var unsMainnetConfigJSON, _ := f.ReadFile("uns/uns-config.json")
