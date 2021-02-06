package resolution

import (
	"github.com/DeRain/resolution-go/cns/contracts/proxyreader"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	kns "github.com/jgimeno/go-namehash"
)

const DefaultProvider = "https://mainnet.infura.io/v3/f3c9708a98674a9fb0ce475354d1e711"
const ProxyReaderMainnetAddress = "0xa6E7cEf2EDDEA66352Fd68E5915b60BDbb7309f5"

type Cns struct {
	ProxyReader *proxyreader.Contract
}

func NewCnsWithDefaultProvider() (*Cns, error) {
	address := common.HexToAddress(ProxyReaderMainnetAddress)
	backend, err := ethclient.Dial(DefaultProvider)
	if err != nil {
		return nil, err
	}
	contract, err := proxyreader.NewContract(address, backend)
	if err != nil {
		return nil, err
	}

	return &Cns{ProxyReader: contract}, nil
}

func (c *Cns) Record(domain string, key string) (record string, err error) {
	name := NormalizeName(domain)
	namehash := kns.NameHash(name)
	record, err = c.ProxyReader.Get(&bind.CallOpts{Pending: false}, key, namehash.Big())
	return
}
