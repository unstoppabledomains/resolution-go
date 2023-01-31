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

var unsMainnetConfigJSON = []byte(`
{
	"contracts": {
		"UNSRegistry": {
			"address": "0x049aba7510f45BA5b64ea9E658E342F904DB358D",
			"implementation": "0xa715562307AA8AEDCba976b3793b3337F371c14a",
			"legacyAddresses": [],
			"deploymentBlock": "0xd62e9d",
			"forwarder": "0x049aba7510f45BA5b64ea9E658E342F904DB358D"
		},
		"CNSRegistry": {
			"address": "0xD1E5b0FF1287aA9f9A268759062E4Ab08b9Dacbe",
			"legacyAddresses": [],
			"deploymentBlock": "0x8a958b",
			"forwarder": "0x97B0E89fC1B7eD4A8B237D9d8Fcce9b234f25A37"
		},
		"MintingManager": {
			"address": "0x2a7084870bB724175a3C96Da8FaA55128fa3E19D",
			"implementation": "0x8caAeaD19aab5f54C94BB9F4be32e200E54AC8D7",
			"legacyAddresses": [],
			"deploymentBlock": "0xc2fee0",
			"forwarder": "0xb970fbCF52cd8111c76c379D4f2FE12E7f8AE7fb"
		},
		"ProxyAdmin": {
			"address": "0xAA16DA78110D9A9742c760a1a064F28654Ab93de",
			"legacyAddresses": [],
			"deploymentBlock": "0xc2fedc"
		},
		"SignatureController": {
			"address": "0x82EF94294C95aD0930055f31e53A34509227c5f7",
			"legacyAddresses": [],
			"deploymentBlock": "0x8a95a6"
		},
		"MintingController": {
			"address": "0xb0EE56339C3253361730F50c08d3d7817ecD60Ca",
			"legacyAddresses": [],
			"deploymentBlock": "0x8a95aa",
			"deprecated": true
		},
		"WhitelistedMinter": {
			"address": "0xd3fF3377b0ceade1303dAF9Db04068ef8a650757",
			"legacyAddresses": [],
			"deploymentBlock": "0xa76ad3",
			"deprecated": true
		},
		"URIPrefixController": {
			"address": "0x09B091492759737C03da9dB7eDF1CD6BCC3A9d91",
			"legacyAddresses": [],
			"deploymentBlock": "0x8a95ae",
			"deprecated": true
		},
		"DomainZoneController": {
			"address": "0xeA70777e28E00E81f58b8921fC47F78B8a72eFE7",
			"legacyAddresses": [],
			"deploymentBlock": "0x98ca20",
			"deprecated": true
		},
		"Resolver": {
			"address": "0xb66DcE2DA6afAAa98F2013446dBCB0f4B0ab2842",
			"legacyAddresses": [
				"0xa1cac442be6673c49f8e74ffc7c4fd746f3cbd0d",
				"0x878bc2f3f717766ab69c0a5f9a6144931e61aed3"
			],
			"deploymentBlock": "0x960844",
			"forwarder": "0x486eb10E4F48C038513ECAf11585Ca2779768CF2"
		},
		"ProxyReader": {
			"address": "0x578853aa776Eef10CeE6c4dd2B5862bdcE767A8B",
			"implementation": "0xfE97D99558BDe54FB9Cb20F0C45f9199bB8df0a0",
			"legacyAddresses": [
				"0x6E68f3EaAD2CC946C6CC7f4859251d8D70Dd3EDB",
				"0x1BDc0fD4fbABeed3E611fd6195fCd5d41dcEF393",
				"0x58034A288D2E56B661c9056A0C27273E5460B63c",
				"0xc3C2BAB5e3e52DBF311b2aAcEf2e40344f19494E",
				"0xfEe4D4F0aDFF8D84c12170306507554bC7045878",
				"0xa6E7cEf2EDDEA66352Fd68E5915b60BDbb7309f5",
				"0x7ea9Ee21077F84339eDa9C80048ec6db678642B1"
			],
			"deploymentBlock": "0xf2f03c"
		},
		"TwitterValidationOperator": {
			"address": "0x2F659766E3D08561CA3408FbAba7C0749ab2c402",
			"legacyAddresses": ["0xbb486C6E9cF1faA86a6E3eAAFE2e5665C0507855"],
			"deploymentBlock": "0xc300b5"
		},
		"FreeMinter": {
			"address": "0x1fC985cAc641ED5846b631f96F35d9b48Bc3b834",
			"legacyAddresses": [],
			"deploymentBlock": "0xacc390",
			"deprecated": true
		},
		"MintableERC721Predicate": {
			"address": "0x932532aA4c0174b8453839A6E44eE09Cc615F2b7",
			"legacyAddresses": [],
			"deploymentBlock": "0xa3cf69"
		},
		"RootChainManager": {
			"address": "0xA0c68C638235ee32657e8f720a23ceC1bFc77C77",
			"legacyAddresses": [],
			"deploymentBlock": "0xa3cf4d"
		}
	}
}`)

var unsPolygonConfigJSON = []byte(`
{
	"contracts": {
		"UNSRegistry": {
			"address": "0xa9a6A3626993D487d2Dbda3173cf58cA1a9D9e9f",
			"legacyAddresses": [],
			"deploymentBlock": "0x019d6188",
			"implementation": "0x5442953b0BFFf69FC945f5f1387cbFD2e2673447",
			"forwarder": "0xa9a6A3626993D487d2Dbda3173cf58cA1a9D9e9f"
		},
		"CNSRegistry": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"forwarder": "0x0000000000000000000000000000000000000000"
		},
		"MintingManager": {
			"address": "0x7be83293BeeDc9Eba1bd76c66A65F10F3efaeC26",
			"legacyAddresses": [],
			"deploymentBlock": "0x01272f41",
			"implementation": "0xBb45a6E10224Aa36EAcd812205F3763D353e9783",
			"forwarder": "0xC37d3c4326ab0E1D2b9D8b916bBdf5715f780fcF"
		},
		"ProxyAdmin": {
			"address": "0xe1D668052D52388F52b90f4d1798DB2b04bC3b88",
			"legacyAddresses": [],
			"deploymentBlock": "0x01272d15"
		},
		"SignatureController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		},
		"MintingController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"WhitelistedMinter": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"URIPrefixController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"DomainZoneController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"Resolver": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"forwarder": "0x0000000000000000000000000000000000000000"
		},
		"ProxyReader": {
			"address": "0x91EDd8708062bd4233f4Dd0FCE15A7cb4d500091",
			"implementation": "0x9D0F27232b5c364488083e3B10F6963F635Ae521",
			"legacyAddresses": [
				"0x68Af8fFFCdC6218836C62Bc2Fd2D35dA544471dD",
				"0x3E67b8c702a1292d1CEb025494C84367fcb12b45",
				"0x423F2531bd5d3C3D4EF7C318c2D1d9BEDE67c680",
				"0xA3f32c8cd786dc089Bd1fC175F2707223aeE5d00"
			],
			"deploymentBlock": "0x021b1c05"
		},
		"TwitterValidationOperator": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		},
		"FreeMinter": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"MintableERC721Predicate": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		},
		"RootChainManager": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		}
	}
}`)

var unsMumbaiConfigJSON = []byte(`
{
	"contracts": {
		"UNSRegistry": {
			"address": "0x2a93C52E7B6E7054870758e15A1446E769EdfB93",
			"legacyAddresses": [],
			"deploymentBlock": "0x0189f713",
			"implementation": "0xAc1a1F2136BfDe3a353a95C0676Cd0d55f311ee3",
			"forwarder": "0x2a93C52E7B6E7054870758e15A1446E769EdfB93"
		},
		"CNSRegistry": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"forwarder": "0x0000000000000000000000000000000000000000"
		},
		"MintingManager": {
			"address": "0x428189346bb3CC52f031A1092fd47C919AC30A9f",
			"legacyAddresses": [],
			"deploymentBlock": "0x01213f4a",
			"implementation": "0xCC17E698bA21bae4277579F22cA51135AaF00777",
			"forwarder": "0xEf3a491A8750BEC2Dff5339CF6Df94436d432C4d"
		},
		"ProxyAdmin": {
			"address": "0x460d63117c7Ab1624b7474C45BF46eC6702f57ce",
			"legacyAddresses": [],
			"deploymentBlock": "0x01213b22"
		},
		"SignatureController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		},
		"MintingController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"WhitelistedMinter": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"URIPrefixController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"DomainZoneController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"Resolver": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"forwarder": "0x0000000000000000000000000000000000000000"
		},
		"ProxyReader": {
			"address": "0xBD4674F11d512120dFc8BAe5f84963d7419A5db2",
			"implementation": "0xacE4C348E1703657201082Ba449aA45ADf8F936a",
			"legacyAddresses": [
				"0x71f7C0A978A541aB13Bd5783470f38b0dd71Cf78",
				"0x6fe7c857C1B0E54492C8762f27e0a45CA7ff264B",
				"0xbd9e01F6513E7C05f71Bf21d419a3bDF1EA9104b",
				"0x332A8191905fA8E6eeA7350B5799F225B8ed30a9"
			],
			"deploymentBlock": "0x01bb07d3"
		},
		"TwitterValidationOperator": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		},
		"FreeMinter": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"MintableERC721Predicate": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		},
		"RootChainManager": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		}
	}
}`)

var unsGoerliConfigJSON = []byte(`
{
	"contracts": {
		"UNSRegistry": {
			"address": "0x070e83FCed225184E67c86302493ffFCDB953f71",
			"implementation": "0x4473e84898E3F58feEFb7529dfF9E83Ff26CCae9",
			"legacyAddresses": [],
			"deploymentBlock": "0x5b57ea",
			"forwarder": "0x070e83FCed225184E67c86302493ffFCDB953f71"
		},
		"CNSRegistry": {
			"address": "0x801452cFAC27e79a11c6b185986fdE09e8637589",
			"legacyAddresses": [],
			"deploymentBlock": "0x5b57d7",
			"forwarder": "0x00443017FFaa4C840Caf5Dc7d3CB59147f363080"
		},
		"MintingManager": {
			"address": "0x9ee42D3EB042e06F8Cd241890C4fA0d51e4DA345",
			"implementation": "0xFB11410f3067BB6Db61bC335f0de23bE87A1767e",
			"legacyAddresses": [],
			"deploymentBlock": "0x5b57ec",
			"forwarder": "0x7F9F48cF94C69ce91D4b442DA186F31118ac0185"
		},
		"ProxyAdmin": {
			"address": "0xf4906E210523F9dA79E33811A44EE000441F4E04",
			"legacyAddresses": [],
			"deploymentBlock": "0x5b57e8"
		},
		"SignatureController": {
			"address": "0x5199dAE4B24B987ba18FcE1b64664D1B798d372B",
			"legacyAddresses": [],
			"deploymentBlock": "0x5b57d8"
		},
		"MintingController": {
			"address": "0xCEC41677be322049cC885c0DAe2fE0D52CA195ca",
			"legacyAddresses": [],
			"deploymentBlock": "0x5b57d9",
			"deprecated": true
		},
		"WhitelistedMinter": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"URIPrefixController": {
			"address": "0x29465e3d2daA588E62375977bCe9b3f51406a794",
			"legacyAddresses": [],
			"deploymentBlock": "0x5b57da",
			"deprecated": true
		},
		"DomainZoneController": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"Resolver": {
			"address": "0x0555344A5F440Bd1d8cb6B42db46c5e5D4070437",
			"legacyAddresses": [],
			"deploymentBlock": "0x5b57dc",
			"forwarder": "0xFCc1A95B7287Ae7a8B7cA813F12991dF5714d4C7"
		},
		"ProxyReader": {
			"address": "0x76007c52C73972A441aFA1A0E1016B140ffdE689",
			"implementation": "0x0B0A42A7FeA63e75396D0dcD77626F706AB9Fdfb",
			"legacyAddresses": [
				"0x77cb0e7503Ea82315421BcF0eE9603451cd285F6",
				"0xE3b961856C417d081a02cBa0161a051268F52677",
				"0x9A70ff906D422C2FD0F7B94244D6b36DB62Ee982",
				"0xFc5f608149f4D9e2Ed0733efFe9DD57ee24BCF68"
			],
			"deploymentBlock": "0x78b972"
		},
		"TwitterValidationOperator": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0"
		},
		"FreeMinter": {
			"address": "0x0000000000000000000000000000000000000000",
			"legacyAddresses": [],
			"deploymentBlock": "0x0",
			"deprecated": true
		},
		"MintableERC721Predicate": {
			"address": "0x56E14C4C1748a818a5564D33cF774c59EB3eDF59",
			"legacyAddresses": [],
			"deploymentBlock": "0x2fc240"
		},
		"RootChainManager": {
			"address": "0xBbD7cBFA79faee899Eaf900F13C9065bF03B1A74",
			"legacyAddresses": [],
			"deploymentBlock": "0x2dc9b9"
		}
	}
}
`)
