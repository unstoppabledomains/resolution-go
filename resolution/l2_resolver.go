package resolution

import (
	"github.com/unstoppabledomains/resolution-go/v3/resolution/namingservice"
)

const zeroAddress = "0x0000000000000000000000000000000000000000"

type genericFunctions struct {
	L1Function func() (interface{}, error)
	L2Function func() (interface{}, error)
	ZFunction  func() (interface{}, error)
}

func resolveGeneric(functions genericFunctions) (interface{}, error) {
	type chanStruct struct {
		result interface{}
		err    error
	}

	c1 := make(chan chanStruct)
	c2 := make(chan chanStruct)
	cz := make(chan chanStruct)

	returnToChannel := func(f func() (interface{}, error), c chan chanStruct) {
		r, e := f()
		c <- chanStruct{r, e}
	}

	go returnToChannel(functions.L1Function, c1)
	go returnToChannel(functions.L2Function, c2)
	go returnToChannel(functions.ZFunction, cz)

	result := <-c2
	if result.err != nil {
		_, notRegistered := result.err.(*DomainNotRegisteredError)
		if notRegistered {
			result = <-c1
			if result.err != nil {
				_, notRegistered := result.err.(*DomainNotRegisteredError)
				if notRegistered {
					result = <-cz
					if result.err != nil {
						return nil, result.err
					}
				}
			}
		}
	}

	return result.result, result.err
}

type stringMapFunction func() (map[string]string, error)
type stringMapResolverParams struct {
	L1Function stringMapFunction
	L2Function stringMapFunction
	ZFunction  stringMapFunction
}

func resolveStringMap(functions stringMapResolverParams) (map[string]string, error) {
	convertToGenericFunction := func(f stringMapFunction) func() (interface{}, error) {
		return func() (interface{}, error) {
			res, err := f()
			return res, err
		}
	}

	res, err := resolveGeneric(genericFunctions{
		L1Function: convertToGenericFunction(functions.L1Function),
		L2Function: convertToGenericFunction(functions.L2Function),
		ZFunction:  convertToGenericFunction(functions.ZFunction),
	})

	strmap, ok := res.(map[string]string)
	if ok {
		return strmap, err
	}
	return nil, err
}

type stringFunction func() (string, error)
type stringResolverParams struct {
	L1Function stringFunction
	L2Function stringFunction
	ZFunction  stringFunction
}

func resolveString(functions stringResolverParams) (string, error) {
	convertToGenericFunction := func(f stringFunction) func() (interface{}, error) {
		return func() (interface{}, error) {
			res, err := f()
			return res, err
		}
	}

	res, err := resolveGeneric(genericFunctions{
		L1Function: convertToGenericFunction(functions.L1Function),
		L2Function: convertToGenericFunction(functions.L2Function),
		ZFunction:  convertToGenericFunction(functions.ZFunction),
	})

	str, ok := res.(string)
	if ok {
		return str, err
	}
	return "", err
}

type stringMapLocationFuction func() (map[string]namingservice.Location, error)
type stringMapLocationParams struct {
	L1Function stringMapLocationFuction
	L2Function stringMapLocationFuction
	ZFunction  stringMapLocationFuction
}

func resolveLocations(functions stringMapLocationParams) (map[string]namingservice.Location, error) {
	type chanStruct struct {
		result map[string]namingservice.Location
		err    error
	}

	c1 := make(chan chanStruct)
	c2 := make(chan chanStruct)
	cz := make(chan chanStruct)

	returnToChannel := func(f func() (map[string]namingservice.Location, error), c chan chanStruct) {
		r, e := f()
		c <- chanStruct{r, e}
	}

	go returnToChannel(functions.L1Function, c1)
	go returnToChannel(functions.L2Function, c2)
	go returnToChannel(functions.ZFunction, cz)

	resultL1 := <-c1
	resultL2 := <-c2
	resultZ := <-cz

	locations := map[string]namingservice.Location{}

	for domainName, location := range resultL1.result {
		if location.OwnerAddress != zeroAddress {
			locations[domainName] = namingservice.Location{
				RegistryAddress:       location.RegistryAddress,
				ResolverAddress:       location.ResolverAddress,
				OwnerAddress:          location.OwnerAddress,
				BlockchainProviderUrl: location.BlockchainProviderUrl,
				NetworkId:             location.NetworkId,
				Blockchain:            "ETH",
			}
			return locations, nil
		} else {
			locations[domainName] = namingservice.Location{
				RegistryAddress:       "",
				ResolverAddress:       "",
				NetworkId:             0,
				Blockchain:            "",
				OwnerAddress:          "",
				BlockchainProviderUrl: "",
			}
		}
	}

	for domainName, location := range resultL2.result {
		if location.OwnerAddress != zeroAddress {
			locations[domainName] = namingservice.Location{
				RegistryAddress:       location.RegistryAddress,
				ResolverAddress:       location.ResolverAddress,
				OwnerAddress:          location.OwnerAddress,
				BlockchainProviderUrl: location.BlockchainProviderUrl,
				NetworkId:             location.NetworkId,
				Blockchain:            "MATIC",
			}
			return locations, nil
		} else {
			locations[domainName] = namingservice.Location{
				RegistryAddress:       "",
				ResolverAddress:       "",
				NetworkId:             0,
				Blockchain:            "",
				OwnerAddress:          "",
				BlockchainProviderUrl: "",
			}
		}
	}

	for domainName, location := range resultZ.result {
		if location.OwnerAddress != zeroAddress {
			locations[domainName] = namingservice.Location{
				RegistryAddress:       location.RegistryAddress,
				ResolverAddress:       location.ResolverAddress,
				OwnerAddress:          location.OwnerAddress,
				BlockchainProviderUrl: location.BlockchainProviderUrl,
				NetworkId:             location.NetworkId,
				Blockchain:            "ZIL",
			}
		} else {
			locations[domainName] = namingservice.Location{
				RegistryAddress:       "",
				ResolverAddress:       "",
				NetworkId:             0,
				Blockchain:            "",
				OwnerAddress:          "",
				BlockchainProviderUrl: "",
			}
			return locations, nil
		}
	}
	return locations, nil
}
