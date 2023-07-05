package resolution

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/unstoppabledomains/resolution-go/v3/ens/contracts/ensreader"
)

type EnsBuilder interface {
	// SetContractBackend set Ethereum backend for communication with Ens registry
	SetContractBackend(backend bind.ContractBackend) EnsBuilder

	// SetContractBackendProviderUrl set Ethereum backend Rpc URL
	SetContractBackendProviderUrl(url string) EnsBuilder

	// Build Uns instance
	Build() (*Ens, error)
}

type ensBuilder struct {
	contractBackend bind.ContractBackend
	network         string
	providerUrl     string
}

func NewEnsBuilder() EnsBuilder {
	return &ensBuilder{
		network: "mainnet",
	}
}

// SetContractBackend set Ethereum backend for communication with ENS registry
func (eb *ensBuilder) SetContractBackend(backend bind.ContractBackend) EnsBuilder {
	eb.contractBackend = backend
	return eb
}

// SetContractBackendProviderUrl set Ethereum backend Rpc URL
func (eb *ensBuilder) SetContractBackendProviderUrl(url string) EnsBuilder {
	eb.providerUrl = url
	return eb
}

func (ens *ensBuilder) BuildService(netContracts contracts, contractBackend bind.ContractBackend, network, provider string) (*EnsService, error) {
	ensRegistryReader := common.HexToAddress(netContracts["Registry"].Address)

	if contractBackend == nil {
		backend, err := ethclient.Dial(provider)
		if err != nil {
			return nil, err
		}
		contractBackend = backend
	}

	ensReaderContract, err := ensreader.NewContract(ensRegistryReader, contractBackend)

	if err != nil {
		return nil, err
	}

	return &EnsService{
		ensReaderContract:     ensReaderContract,
		contractBackend:       contractBackend,
		networkId:             1,
		blockchainProviderUrl: provider,
	}, nil
}

func (eb *ensBuilder) Build() (*Ens, error) {
	contracts, err := newEnsContracts()

	if err != nil {
		return nil, err
	}

	if eb.network == "" {
		return nil, &EnsConfigurationError{InvalidField: "network"}
	}

	service, err := eb.BuildService(contracts[eb.network], eb.contractBackend, eb.network, eb.providerUrl)

	if err != nil {
		return nil, err
	}

	return &Ens{
		*service,
	}, nil
}
