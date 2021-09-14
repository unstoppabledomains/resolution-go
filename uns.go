package resolution

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/unstoppabledomains/resolution-go/dnsrecords"
)

// Uns is a naming service handles Unstoppable domains resolution.
type Uns struct {
	l1Service UnsService
	l2Service UnsService
}

// Data Get raw data attached to domain
func (c *Uns) Data(domainName string, keys []string) (*struct {
	Resolver common.Address
	Owner    common.Address
	Values   []string
}, error) {
	prepGenericFunction := func(s *UnsService) func() (interface{}, error) {
		return func() (interface{}, error) {
			res, err := s.Data(domainName, keys)
			return res, err
		}
	}

	res, err := ResolveGeneric(GenericFunctions{
		L1Function: prepGenericFunction(&c.l1Service),
		L2Function: prepGenericFunction(&c.l2Service),
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
	return ResolveStringMap(StringMapResolverParams{
		L1Function: func() (map[string]string, error) { return c.l1Service.Records(domainName, keys) },
		L2Function: func() (map[string]string, error) { return c.l2Service.Records(domainName, keys) }})
}

func (c *Uns) Record(domainName string, key string) (string, error) {
	return ResolveString(StringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.Record(domainName, key) },
		L2Function: func() (string, error) { return c.l2Service.Record(domainName, key) }})
}

func (c *Uns) Addr(domainName string, ticker string) (string, error) {
	return ResolveString(StringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.Addr(domainName, ticker) },
		L2Function: func() (string, error) { return c.l2Service.Addr(domainName, ticker) }})
}

func (c *Uns) AddrVersion(domainName string, ticker string, version string) (string, error) {
	return ResolveString(StringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.AddrVersion(domainName, ticker, version) },
		L2Function: func() (string, error) { return c.l2Service.AddrVersion(domainName, ticker, version) }})
}

func (c *Uns) Email(domainName string) (string, error) {
	return ResolveString(StringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.Email(domainName) },
		L2Function: func() (string, error) { return c.l2Service.Email(domainName) }})
}

func (c *Uns) Resolver(domainName string) (string, error) {
	return ResolveString(StringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.Resolver(domainName) },
		L2Function: func() (string, error) { return c.l2Service.Resolver(domainName) }})
}

func (c *Uns) Owner(domainName string) (string, error) {
	return ResolveString(StringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.Owner(domainName) },
		L2Function: func() (string, error) { return c.l2Service.Owner(domainName) }})
}

func (c *Uns) IpfsHash(domainName string) (string, error) {
	return ResolveString(StringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.IpfsHash(domainName) },
		L2Function: func() (string, error) { return c.l2Service.IpfsHash(domainName) }})
}

func (c *Uns) HTTPUrl(domainName string) (string, error) {
	return ResolveString(StringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.HTTPUrl(domainName) },
		L2Function: func() (string, error) { return c.l2Service.HTTPUrl(domainName) }})
}

func (c *Uns) AllRecords(domainName string) (map[string]string, error) {
	return ResolveStringMap(StringMapResolverParams{
		L1Function: func() (map[string]string, error) { return c.l1Service.AllRecords(domainName) },
		L2Function: func() (map[string]string, error) { return c.l2Service.AllRecords(domainName) }})
}

func (c *Uns) DNS(domainName string, types []dnsrecords.Type) ([]dnsrecords.Record, error) {
	prepGenericFunction := func(s *UnsService) func() (interface{}, error) {
		return func() (interface{}, error) {
			res, err := s.DNS(domainName, types)
			return res, err
		}
	}

	res, err := ResolveGeneric(GenericFunctions{
		L1Function: prepGenericFunction(&c.l1Service),
		L2Function: prepGenericFunction(&c.l2Service),
	})

	data, ok := res.([]dnsrecords.Record)
	if ok {
		return data, err
	}
	return nil, err
}

func (c *Uns) IsSupportedDomain(domainName string) (bool, error) {
	return c.l1Service.IsSupportedDomain(domainName)
}

func (c *Uns) TokenURI(domainName string) (string, error) {
	return ResolveString(StringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.TokenURI(domainName) },
		L2Function: func() (string, error) { return c.l2Service.TokenURI(domainName) }})
}

func (c *Uns) TokenURIMetadata(domainName string) (TokenMetadata, error) {
	prepGenericFunction := func(s *UnsService) func() (interface{}, error) {
		return func() (interface{}, error) {
			res, err := s.TokenURIMetadata(domainName)
			return res, err
		}
	}

	res, err := ResolveGeneric(GenericFunctions{
		L1Function: prepGenericFunction(&c.l1Service),
		L2Function: prepGenericFunction(&c.l2Service),
	})

	data, ok := res.(TokenMetadata)
	if ok {
		return data, err
	}
	return TokenMetadata{}, err
}

func (c *Uns) Unhash(domainHash string) (string, error) {
	return ResolveString(StringResolverParams{
		L1Function: func() (string, error) { return c.l1Service.Unhash(domainHash) },
		L2Function: func() (string, error) { return c.l2Service.Unhash(domainHash) }})
}

func (c *Uns) Namehash(domainName string) (string, error) {
	return c.l1Service.Namehash(domainName)
}
