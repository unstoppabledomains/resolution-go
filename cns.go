package resolution

import (
	"github.com/DeRain/resolution-go/proxyreader"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

const ProxyReaderMainnetAddress = "0xa6E7cEf2EDDEA66352Fd68E5915b60BDbb7309f5"

// Registry is the structure for the registry contract
type Cns struct {
	backend      bind.ContractBackend
	Contract     *proxyreader.Contract
	ContractAddr common.Address
}

func NewCns(backend bind.ContractBackend) *Cns {
	// todo
	return &Cns{backend: backend}
}

func (c *Cns) Record(domain string, key string) (record string, err error) {
	// todo

	return
}
