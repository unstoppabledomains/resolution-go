package resolution

import (
	"github.com/DeRain/resolution-go/cns/contracts/proxyreader"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	kns "github.com/jgimeno/go-namehash"
	s "strings"
)

const DefaultProvider = "https://mainnet.infura.io/v3/f3c9708a98674a9fb0ce475354d1e711"

var zeroAddress = common.HexToAddress("0x0")
var proxyReaderMainnetAddress = common.HexToAddress("0xa6E7cEf2EDDEA66352Fd68E5915b60BDbb7309f5")

type Cns struct {
	ProxyReader *proxyreader.Contract
}

func NewCns(backend bind.ContractBackend) (*Cns, error) {
	contract, err := proxyreader.NewContract(proxyReaderMainnetAddress, backend)
	if err != nil {
		return nil, err
	}

	return &Cns{ProxyReader: contract}, nil
}

func NewCnsWithDefaultBackend() (*Cns, error) {
	backend, err := ethclient.Dial(DefaultProvider)
	if err != nil {
		return nil, err
	}
	cns, err := NewCns(backend)
	if err != nil {
		return nil, err
	}

	return cns, nil
}

func (c *Cns) Data(domainName string, keys []string) (*struct {
	Resolver common.Address
	Owner    common.Address
	Values   []string
}, error) {
	normalizedName := NormalizeName(domainName)
	namehash := kns.NameHash(normalizedName)
	tokenId := namehash.Big()
	data, err := c.ProxyReader.GetData(&bind.CallOpts{Pending: false}, keys, tokenId)
	if err != nil {
		return nil, err
	}
	if data.Owner == zeroAddress {
		return nil, &DomainNotRegistered{DomainName: normalizedName}
	}
	if data.Resolver == zeroAddress {
		return nil, &DomainNotConfigured{DomainName: normalizedName}
	}

	return &data, nil
}

func (c *Cns) Records(domainName string, keys []string) ([]string, error) {
	data, err := c.Data(domainName, keys)
	if err != nil {
		return nil, err
	}

	return data.Values, nil
}

func (c *Cns) Record(domainName string, key string) (string, error) {
	data, err := c.Data(domainName, []string{key})
	if err != nil {
		return "", err
	}
	if len(data.Values) == 0 {
		return "", nil
	}

	return data.Values[0], nil
}

func (c *Cns) Addr(domainName string, ticker string) (string, error) {
	key := "crypto." + s.ToUpper(ticker) + ".address"
	value, err := c.Record(domainName, key)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *Cns) AddrVersion(domainName string, ticker string, version string) (string, error) {
	key := "crypto." + s.ToUpper(ticker) + ".version." + s.ToUpper(version) + ".address"
	value, err := c.Record(domainName, key)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *Cns) Email(domainName string) (string, error) {
	key := "whois.email.value"
	value, err := c.Record(domainName, key)
	if err != nil {
		return "", err
	}

	return value, nil
}

func (c *Cns) Resolver(domainName string) (common.Address, error) {
	data, err := c.Data(domainName, []string{})
	if err != nil {
		return common.Address{}, err
	}

	return data.Resolver, nil
}

// todo ipfs hash
// todo http url (redirect)
// todo dns records
// todo all records
