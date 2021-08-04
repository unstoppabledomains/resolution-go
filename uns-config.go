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

func newContracts() (contracts, contracts, error) {
	var mainnetContractsObject struct {
		Contracts contracts
	}
	var rinkebyContractsObject struct {
		Contracts contracts
	}
	err := json.Unmarshal(unsMainnetConfigJSON, &mainnetContractsObject)
	if err != nil {
		return nil, nil, err
	}
	err = json.Unmarshal(unsRinkebyConfigJSON, &rinkebyContractsObject)
	if err != nil {
		return nil, nil, err
	}
	return mainnetContractsObject.Contracts, rinkebyContractsObject.Contracts, nil
}

var unsMainnetConfigJSON = []byte(`
{
    "version": "0.1.0",
    "contracts": {
        "UNSRegistry": {
            "address": "0x049aba7510f45BA5b64ea9E658E342F904DB358D",
            "implementation": "0x6d4FaFcD5B4A8b360c80eAca8D91e3Bc814ecc39",
            "legacyAddresses": [],
            "deploymentBlock": "0xc2fede"
        },
        "CNSRegistry": {
            "address": "0xD1E5b0FF1287aA9f9A268759062E4Ab08b9Dacbe",
            "legacyAddresses": [],
            "deploymentBlock": "0x8a958b"
        },
        "MintingManager": {
            "address": "0x2a7084870bB724175a3C96Da8FaA55128fa3E19D",
            "implementation": "0x2f133a06fe4fc845E41261aCFF6831a727ea9062",
            "legacyAddresses": [],
            "deploymentBlock": "0xc2fee0"
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
            "deploymentBlock": "0x8a95aa"
        },
        "WhitelistedMinter": {
            "address": "0xd3fF3377b0ceade1303dAF9Db04068ef8a650757",
            "legacyAddresses": [],
            "deploymentBlock": "0xa76ad3"
        },
        "URIPrefixController": {
            "address": "0x09B091492759737C03da9dB7eDF1CD6BCC3A9d91",
            "legacyAddresses": [],
            "deploymentBlock": "0x8a95ae"
        },
        "DomainZoneController": {
            "address": "0xeA70777e28E00E81f58b8921fC47F78B8a72eFE7",
            "legacyAddresses": [],
            "deploymentBlock": "0x98ca20"
        },
        "Resolver": {
            "address": "0xb66DcE2DA6afAAa98F2013446dBCB0f4B0ab2842",
            "legacyAddresses": [
                "0xa1cac442be6673c49f8e74ffc7c4fd746f3cbd0d",
                "0x878bc2f3f717766ab69c0a5f9a6144931e61aed3"
            ],
            "deploymentBlock": "0x960844"
        },
        "ProxyReader": {
            "address": "0xfEe4D4F0aDFF8D84c12170306507554bC7045878",
            "legacyAddresses": [
                "0xa6E7cEf2EDDEA66352Fd68E5915b60BDbb7309f5",
                "0x7ea9Ee21077F84339eDa9C80048ec6db678642B1"
            ],
            "deploymentBlock": "0xc300b4"
        },
        "TwitterValidationOperator": {
            "address": "0x2F659766E3D08561CA3408FbAba7C0749ab2c402",
            "legacyAddresses": [
                "0xbb486C6E9cF1faA86a6E3eAAFE2e5665C0507855"
            ],
            "deploymentBlock": "0xc300b5"
        },
        "FreeMinter": {
            "address": "0x1fC985cAc641ED5846b631f96F35d9b48Bc3b834",
            "legacyAddresses": [],
            "deploymentBlock": "0xacc390"
        }
    }
}`)

var unsRinkebyConfigJSON = []byte(`
{
    "version": "0.1.0",
    "contracts": {
        "UNSRegistry": {
            "address": "0x7fb83000B8eD59D3eAD22f0D584Df3a85fBC0086",
            "implementation": "0x0c97caecb791620c61e26c5a9754150a1cae2159",
            "legacyAddresses": [],
            "deploymentBlock": "0x85e628"
        },
        "CNSRegistry": {
            "address": "0xAad76bea7CFEc82927239415BB18D2e93518ecBB",
            "legacyAddresses": [],
            "deploymentBlock": "0x7232bc"
        },
        "MintingManager": {
            "address": "0xdAAf99A920D31F4f5720e4667b12b24e54A03070",
            "implementation": "0xb8efb23b97408147ca63ae0494b2242c5029ffc4",
            "legacyAddresses": [],
            "deploymentBlock": "0x85e629"
        },
        "ProxyReader": {
            "address": "0x299974AeD8911bcbd2C61262605b89F591a53E83",
            "legacyAddresses": [
                "0x9F19473F6a98a715176291c930558E1954fd3D1e",
                "0x3A2e74CF832cbA3d77E72708d55370119E4323a6"
            ],
            "deploymentBlock": "0x8671b4"
        },
        "ProxyAdmin": {
            "address": "0xaf9815005A208d1460b6fC60B4f90B9f2185E88c",
            "legacyAddresses": [],
            "deploymentBlock": "0x85e627"
        },
        "SignatureController": {
            "address": "0x66a5e3e2C27B4ce4F46BBd975270BE154748D164",
            "legacyAddresses": [],
            "deploymentBlock": "0x7232be"
        },
        "MintingController": {
            "address": "0x51765307AeB3Df2E647014a2C501d5324212467c",
            "legacyAddresses": [],
            "deploymentBlock": "0x7232bf",
            "deprecated": true
        },
        "WhitelistedMinter": {
            "address": "0xbcB32f13f90978a9e059E8Cb40FaA9e6619d98e7",
            "legacyAddresses": [],
            "deploymentBlock": "0x7232c6",
            "deprecated": true
        },
        "URIPrefixController": {
            "address": "0xe1d2e4B9f0518CA5c803073C3dFa886470627237",
            "legacyAddresses": [],
            "deploymentBlock": "0x7232c0",
            "deprecated": true
        },
        "DomainZoneController": {
            "address": "0x6f8F96A566663C1d4fEe70edD37E9b62Fe39dE5D",
            "legacyAddresses": [],
            "deploymentBlock": "0x7232c2",
            "deprecated": true
        },
        "Resolver": {
            "address": "0x95AE1515367aa64C462c71e87157771165B1287A",
            "legacyAddresses": [],
            "deploymentBlock": "0x7232cf"
        },
        "TwitterValidationOperator": {
            "address": "0x9ea4A63184ebE9CBA55CD1af473D98075Aa02b4C",
            "legacyAddresses": [
                "0x1CB337b3b208dc29a6AcE8d11Bb591b66c5Dd83d"
            ],
            "deploymentBlock": "0x86935e"
        },
        "FreeMinter": {
            "address": "0x84214215904cDEbA9044ECf95F3eBF009185AAf4",
            "legacyAddresses": [],
            "deploymentBlock": "0x740d93",
            "deprecated": true
        }
    }
}`)