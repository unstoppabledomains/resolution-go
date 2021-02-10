package resolution

import (
	"github.com/DeRain/resolution-go/cns/contracts/proxyreader"
	"github.com/DeRain/resolution-go/cns/supportedkeys"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	kns "github.com/jgimeno/go-namehash"
	"github.com/spf13/viper"
	s "strings"
)

const DefaultProvider = "https://mainnet.infura.io/v3/f3c9708a98674a9fb0ce475354d1e711"

var zeroAddress = common.HexToAddress("0x0")
var proxyReaderMainnetAddress = common.HexToAddress("0xa6E7cEf2EDDEA66352Fd68E5915b60BDbb7309f5")

type Cns struct {
	ProxyReader   *proxyreader.Contract
	SupportedKeys *viper.Viper
}

func NewCns(backend bind.ContractBackend) (*Cns, error) {
	contract, err := proxyreader.NewContract(proxyReaderMainnetAddress, backend)
	if err != nil {
		return nil, err
	}
	supportedKeys, err := supportedkeys.NewConfig()
	if err != nil {
		return nil, err
	}

	return &Cns{ProxyReader: contract, SupportedKeys: supportedKeys}, nil
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

func (c *Cns) Resolver(domainName string) (string, error) {
	data, err := c.Data(domainName, []string{})
	if err != nil {
		return "", err
	}

	return data.Resolver.String(), nil
}

func (c *Cns) Owner(domainName string) (string, error) {
	data, err := c.Data(domainName, []string{})
	if err != nil {
		return "", err
	}

	return data.Owner.String(), nil
}

func (c *Cns) IpfsHash(domainName string) (string, error) {
	data, err := c.Data(domainName, []string{"dweb.ipfs.hash", "ipfs.html.value"})
	if err != nil {
		return "", err
	}
	if data.Values[0] != "" {
		return data.Values[0], nil
	}
	if data.Values[1] != "" {
		return data.Values[1], nil
	}

	return "", nil
}

func (c *Cns) HttpUrl(domainName string) (string, error) {
	data, err := c.Data(domainName, []string{"browser.redirect_url", "ipfs.redirect_domain.value"})
	if err != nil {
		return "", err
	}
	if data.Values[0] != "" {
		return data.Values[0], nil
	}
	if data.Values[1] != "" {
		return data.Values[1], nil
	}

	return "", nil
}

// todo dns records
// todo all records
