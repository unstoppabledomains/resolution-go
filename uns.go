package resolution

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/unstoppabledomains/resolution-go/v2/dnsrecords"
	"github.com/unstoppabledomains/resolution-go/v2/namingservice"
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
		ZFunction:  func() (map[string]string, error) { return c.zService.Records(domainName, keys) },
	})
}

func (c *Uns) Record(domainName string, key string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.record(domainName, key) },
		L2Function: func() (string, error) { return c.l2Service.record(domainName, key) },
		ZFunction:  func() (string, error) { return c.zService.Record(domainName, key) }})
}

func (c *Uns) Addr(domainName string, ticker string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.addr(domainName, ticker) },
		L2Function: func() (string, error) { return c.l2Service.addr(domainName, ticker) },
		ZFunction:  func() (string, error) { return c.zService.Addr(domainName, ticker) }})
}

func (c *Uns) AddrVersion(domainName string, ticker string, version string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.addrVersion(domainName, ticker, version) },
		L2Function: func() (string, error) { return c.l2Service.addrVersion(domainName, ticker, version) },
		ZFunction:  func() (string, error) { return c.zService.AddrVersion(domainName, ticker, version) }})
}

func (c *Uns) Email(domainName string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.email(domainName) },
		L2Function: func() (string, error) { return c.l2Service.email(domainName) },
		ZFunction:  func() (string, error) { return c.zService.Email(domainName) }})
}

func (c *Uns) Resolver(domainName string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.resolver(domainName) },
		L2Function: func() (string, error) { return c.l2Service.resolver(domainName) },
		ZFunction:  func() (string, error) { return c.zService.Resolver(domainName) }})
}

func (c *Uns) Owner(domainName string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.owner(domainName) },
		L2Function: func() (string, error) { return c.l2Service.owner(domainName) },
		ZFunction:  func() (string, error) { return c.zService.Owner(domainName) },
	})
}

func (c *Uns) IpfsHash(domainName string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.ipfsHash(domainName) },
		L2Function: func() (string, error) { return c.l2Service.ipfsHash(domainName) },
		ZFunction:  func() (string, error) { return c.zService.IpfsHash(domainName) }})
}

func (c *Uns) HTTPUrl(domainName string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.httpUrl(domainName) },
		L2Function: func() (string, error) { return c.l2Service.httpUrl(domainName) },
		ZFunction:  func() (string, error) { return c.zService.HTTPUrl(domainName) }})
}

func (c *Uns) AllRecords(domainName string) (map[string]string, error) {
	standardKeys, err := newSupportedKeys()
	if err != nil {
		return make(map[string]string), err
	}
	metadata, err := c.TokenURIMetadata(domainName)
	if err != nil {
		return make(map[string]string), err
	}
	recordKeys := make([]string, 0, len(metadata.Properties.Records)+len(standardKeys))
	for k := range standardKeys {
		recordKeys = append(recordKeys, k)
	}
	for k := range metadata.Properties.Records {
		recordKeys = append(recordKeys, k)
	}
	recordsMap, err := resolveStringMap(stringMapResolverParams{
		L1Function: func() (map[string]string, error) { return c.l1Service.records(domainName, recordKeys) },
		L2Function: func() (map[string]string, error) { return c.l2Service.records(domainName, recordKeys) },
		ZFunction:  func() (map[string]string, error) { return c.zService.Records(domainName, recordKeys) },
	})
	if err != nil {
		return make(map[string]string), err
	}
	return removeEmptyRecords(recordsMap), nil
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
	return c.l1Service.isSupportedDomain(domainName)
}

func (c *Uns) TokenURI(domainName string) (string, error) {
	return resolveString(stringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.tokenURI(domainName) },
		L2Function: func() (string, error) { return c.l2Service.tokenURI(domainName) },
		ZFunction:  func() (string, error) { return c.zService.TokenURI(domainName) }})
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

func removeEmptyRecords(records map[string]string) map[string]string {
	for k := range records {
		if records[k] == "" {
			delete(records, k)
		}
	}
	return records
}
