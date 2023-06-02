package resolution

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/unstoppabledomains/resolution-go/v4/resolutions/dnsrecords"
	"github.com/unstoppabledomains/resolution-go/v4/resolutions/namingservice"
)

// Uns is a naming service handles Unstoppable domains resolution.
type Uns struct {
	l1Service UnsService
	l2Service UnsService
	zService  Zns
}

// Data Get raw data attached to domain
func (c *Uns) Data(domainName string, keys []string) (*struct {
	Resolver common.Address
	Owner    common.Address
	Values   []string
}, error) {
	convertToGenericFunction := func(s *UnsService) func() (interface{}, error) {
		return func() (interface{}, error) {
			res, err := s.data(domainName, keys)
			return res, err
		}
	}

	convertToGenericZFunction := func(s *Zns) func() (interface{}, error) {
		return func() (interface{}, error) {
			res, err := s.State(domainName)
			return res, err
		}
	}

	res, err := resolveGeneric(genericFunctions{
		L1Function: convertToGenericFunction(&c.l1Service),
		L2Function: convertToGenericFunction(&c.l2Service),
		ZFunction:  convertToGenericZFunction(&c.zService),
	})

	data, ok := res.(*struct {
		Resolver common.Address
		Owner    common.Address
		Values   []string
	})
	if ok {
		return data, err
	}
	return nil, err
}

func (c *Uns) Records(domainName string, keys []string) (map[string]string, error) {
	return resolveStringMap(stringMapResolverParams{
		L1Function: func() (map[string]string, error) { return c.l1Service.records(domainName, keys) },
		L2Function: func() (map[string]string, error) { return c.l2Service.records(domainName, keys) },
		ZFunction: func() (map[string]string, error) {
			isSupported, _ := c.zService.IsSupportedDomain(domainName)
			if isSupported {
				return c.zService.Records(domainName, keys)
			} else {
				return nil, &DomainNotSupportedError{DomainName: domainName}
			}
		}})
}

func (c *Uns) Record(domainName string, key string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.record(domainName, key) },
		L2Function: func() (string, error) { return c.l2Service.record(domainName, key) },
		ZFunction: func() (string, error) {
			isSupported, _ := c.zService.IsSupportedDomain(domainName)
			if isSupported {
				return c.zService.Record(domainName, key)
			} else {
				return "", &DomainNotSupportedError{DomainName: domainName}
			}
		}})
}

func (c *Uns) Addr(domainName string, ticker string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.addr(domainName, ticker) },
		L2Function: func() (string, error) { return c.l2Service.addr(domainName, ticker) },
		ZFunction: func() (string, error) {
			isSupported, _ := c.zService.IsSupportedDomain(domainName)
			if isSupported {
				return c.zService.Addr(domainName, ticker)
			} else {
				return "", &DomainNotSupportedError{DomainName: domainName}
			}
		}})
}

func (c *Uns) AddrVersion(domainName string, ticker string, version string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.addrVersion(domainName, ticker, version) },
		L2Function: func() (string, error) { return c.l2Service.addrVersion(domainName, ticker, version) },
		ZFunction: func() (string, error) {
			isSupported, _ := c.zService.IsSupportedDomain(domainName)
			if isSupported {
				return c.zService.AddrVersion(domainName, ticker, version)
			} else {
				return "", &DomainNotSupportedError{DomainName: domainName}
			}
		}})
}

func (c *Uns) GetAddr(domainName, family, token string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.getAddress(domainName, family, token) },
		L2Function: func() (string, error) { return c.l2Service.getAddress(domainName, family, token) },
		ZFunction: func() (string, error) {
			return "", &DomainNotSupportedError{DomainName: domainName}
		}})
}

func (c *Uns) ReverseOf(addr string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.reverseOf(addr) },
		L2Function: func() (string, error) { return c.l2Service.reverseOf(addr) },
		ZFunction: func() (string, error) {
			return "", &AddressNotSupportedError{Address: addr}
		}})
}

func (c *Uns) Email(domainName string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.email(domainName) },
		L2Function: func() (string, error) { return c.l2Service.email(domainName) },
		ZFunction: func() (string, error) {
			isSupported, _ := c.zService.IsSupportedDomain(domainName)
			if isSupported {
				return c.zService.Email(domainName)
			} else {
				return "", &DomainNotSupportedError{DomainName: domainName}
			}
		}})
}

func (c *Uns) Resolver(domainName string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.resolver(domainName) },
		L2Function: func() (string, error) { return c.l2Service.resolver(domainName) },
		ZFunction: func() (string, error) {
			isSupported, _ := c.zService.IsSupportedDomain(domainName)
			if isSupported {
				return c.zService.Resolver(domainName)
			} else {
				return "", &DomainNotSupportedError{DomainName: domainName}
			}
		}})
}

func (c *Uns) Owner(domainName string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.owner(domainName) },
		L2Function: func() (string, error) { return c.l2Service.owner(domainName) },
		ZFunction: func() (string, error) {
			isSupported, _ := c.zService.IsSupportedDomain(domainName)
			if isSupported {
				return c.zService.Owner(domainName)
			} else {
				return "", &DomainNotRegisteredError{DomainName: domainName}
			}
		},
	})
}

func (c *Uns) IpfsHash(domainName string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.ipfsHash(domainName) },
		L2Function: func() (string, error) { return c.l2Service.ipfsHash(domainName) },
		ZFunction: func() (string, error) {
			isSupported, _ := c.zService.IsSupportedDomain(domainName)
			if isSupported {
				return c.zService.IpfsHash(domainName)
			} else {
				return "", &DomainNotRegisteredError{DomainName: domainName}
			}
		}})
}

func (c *Uns) HTTPUrl(domainName string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.httpUrl(domainName) },
		L2Function: func() (string, error) { return c.l2Service.httpUrl(domainName) },
		ZFunction: func() (string, error) {
			isSupported, _ := c.zService.IsSupportedDomain(domainName)
			if isSupported {
				return c.zService.HTTPUrl(domainName)
			} else {
				return "", &DomainNotRegisteredError{DomainName: domainName}
			}
		}})
}

func (c *Uns) DNS(domainName string, types []dnsrecords.Type) ([]dnsrecords.Record, error) {
	convertToGenericFunction := func(s *UnsService) func() (interface{}, error) {
		return func() (interface{}, error) {
			res, err := s.DNS(domainName, types)
			return res, err
		}
	}

	convertToGenericZFunction := func(s *Zns) func() (interface{}, error) {
		return func() (interface{}, error) {
			res, err := s.DNS(domainName, types)
			return res, err
		}
	}

	res, err := resolveGeneric(genericFunctions{
		L1Function: convertToGenericFunction(&c.l1Service),
		L2Function: convertToGenericFunction(&c.l2Service),
		ZFunction:  convertToGenericZFunction(&c.zService),
	})

	data, ok := res.([]dnsrecords.Record)
	if ok {
		return data, err
	}
	return nil, err
}

func (c *Uns) Locations(domainNames []string) (map[string]namingservice.Location, error) {
	locations, err := resolveLocations(stringMapLocationParams{
		L1Function: func() (map[string]namingservice.Location, error) { return c.l1Service.locations(domainNames) },
		L2Function: func() (map[string]namingservice.Location, error) { return c.l2Service.locations(domainNames) },
		ZFunction:  func() (map[string]namingservice.Location, error) { return c.zService.Locations(domainNames) },
	})

	if err != nil {
		return map[string]namingservice.Location{}, err
	}
	return locations, nil
}

func (c *Uns) IsSupportedDomain(domainName string) (bool, error) {
	l1Supports, err := c.l1Service.isSupportedDomain(domainName)
	if err != nil {
		return false, err
	}
	l2Supports, err := c.l2Service.isSupportedDomain(domainName)
	if err != nil {
		return false, err
	}
	return l1Supports || l2Supports, nil
}

func (c *Uns) TokenURI(domainName string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.tokenURI(domainName) },
		L2Function: func() (string, error) { return c.l2Service.tokenURI(domainName) },
		ZFunction: func() (string, error) {
			isSupported, _ := c.zService.IsSupportedDomain(domainName)
			if isSupported {
				return c.zService.TokenURI(domainName)
			} else {
				return "", &DomainNotSupportedError{DomainName: domainName}
			}
		}})
}

func (c *Uns) TokenURIMetadata(domainName string) (TokenMetadata, error) {
	convertToGenericFunction := func(s *UnsService) func() (interface{}, error) {
		return func() (interface{}, error) {
			res, err := s.tokenURIMetadata(domainName)
			return res, err
		}
	}

	convertToGenericZFunction := func(s *Zns) func() (interface{}, error) {
		return func() (interface{}, error) {
			res, err := s.TokenURIMetadata(domainName)
			return res, err
		}
	}

	res, err := resolveGeneric(genericFunctions{
		L1Function: convertToGenericFunction(&c.l1Service),
		L2Function: convertToGenericFunction(&c.l2Service),
		ZFunction:  convertToGenericZFunction(&c.zService),
	})

	data, ok := res.(TokenMetadata)
	if ok {
		return data, err
	}
	return TokenMetadata{}, err
}

func (c *Uns) Unhash(domainHash string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.unhash(domainHash) },
		L2Function: func() (string, error) { return c.l2Service.unhash(domainHash) },
		ZFunction:  func() (string, error) { return c.zService.Unhash(domainHash) }})
}

func (c *Uns) Namehash(domainName string) (string, error) {
	return c.l1Service.namehash(domainName)
}
