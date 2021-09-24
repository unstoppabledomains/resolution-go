package resolution

type GenericFunctions struct {
	L1Function func() (interface{}, error)
	L2Function func() (interface{}, error)
}

func ResolveGeneric(functions GenericFunctions) (interface{}, error) {
	type chanStruct struct {
		result interface{}
		err    error
	}

	c1 := make(chan chanStruct)
	c2 := make(chan chanStruct)

	returnToChannel := func(f func() (interface{}, error), c chan chanStruct) {
		r, e := f()
		c <- chanStruct{r, e}
	}

	go returnToChannel(functions.L1Function, c1)
	go returnToChannel(functions.L2Function, c2)

	result := <-c2
	if result.err != nil {
		_, notRegistered := result.err.(*DomainNotRegisteredError)
		if notRegistered {
			result = <-c1
		} else {
			return nil, result.err
		}
	}

	return result.result, result.err
}

type stringMapFunction func() (map[string]string, error)
type StringMapResolverParams struct {
	L1Function stringMapFunction
	L2Function stringMapFunction
}

func ResolveStringMap(functions StringMapResolverParams) (map[string]string, error) {
	convertToGenericFunction := func(f stringMapFunction) func() (interface{}, error) {
		return func() (interface{}, error) {
			res, err := f()
			return res, err
		}
	}

	res, err := ResolveGeneric(GenericFunctions{
		L1Function: convertToGenericFunction(functions.L1Function),
		L2Function: convertToGenericFunction(functions.L2Function),
	})

	strmap, ok := res.(map[string]string)
	if ok {
		return strmap, err
	}
	return nil, err
}

type stringFunction func() (string, error)
type StringResolverParams struct {
	L1Function stringFunction
	L2Function stringFunction
}

func ResolveString(functions StringResolverParams) (string, error) {
	convertToGenericFunction := func(f stringFunction) func() (interface{}, error) {
		return func() (interface{}, error) {
			res, err := f()
			return res, err
		}
	}

	res, err := ResolveGeneric(GenericFunctions{
		L1Function: convertToGenericFunction(functions.L1Function),
		L2Function: convertToGenericFunction(functions.L2Function),
	})

	str, ok := res.(string)
	if ok {
		return str, err
	}
	return "", err
}
