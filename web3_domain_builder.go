package resolution

import (
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/unstoppabledomains/resolution-go/v3/udclient"
)

// Web3Builder is a builder to setup and build instance of Web3Domain Resolution
type Web3DomainBuilder interface {
	// SetContractBackend set Ethereum backend for communication with Web3Domain registry
	SetEthContractBackend(backend bind.ContractBackend) Web3DomainBuilder

	// SetContractBackendProviderUrl set Ethereum backend Rpc URL
	SetEthContractBackendProviderUrl(url string) Web3DomainBuilder

	// SetL2ContractBackend set Ethereum backend for communication with Web3Domain L2registry
	SetMaticContractBackend(backend bind.ContractBackend) Web3DomainBuilder

	// SetL2ContractBackendProviderUrl set Polygon backend Rpc URL
	SetMaticContractBackendProviderUrl(url string) Web3DomainBuilder

	// SetMetadataClient set http backend for communication with ERC721 metadata server
	SetMetadataClient(backend MetadataClient) Web3DomainBuilder

	// SetUdClient set http proxy backends for communication with Web3Domain registry
	SetUdClient(apiKey string) Web3DomainBuilder

	// SetEthereumNetwork set Ethereum network for communication with Web3Domain registry
	SetEthereumNetwork(network string) Web3DomainBuilder

	// SetL2EthereumNetwork set Ethereum network for communication with Web3Domain L2 registry
	SetMaticNetwork(network string) Web3DomainBuilder

	Build() (*Web3Domain, error)
}

type web3ClientBuilder struct {
	l1ContractBackend bind.ContractBackend
	l2ContractBackend bind.ContractBackend
	metadataClient    MetadataClient
	l1Network         string
	l2Network         string
	l1ProviderUrl     string
	l2ProviderUrl     string
}

// Web3DomainBuilder Creates builder to setup new instance of Web3Domain
func NewWeb3DomainBuilder() Web3DomainBuilder {
	return &web3ClientBuilder{
		l1Network: "mainnet",
		l2Network: "polygon",
	}
}

// SetContractBackend set Ethereum backend for communication with Web3Domain registry
func (w3b *web3ClientBuilder) SetEthContractBackend(backend bind.ContractBackend) Web3DomainBuilder {
	w3b.l1ContractBackend = backend
	return w3b
}

// SetContractBackendProviderUrl set Ethereum backend Rpc URL
func (w3b *web3ClientBuilder) SetEthContractBackendProviderUrl(url string) Web3DomainBuilder {
	w3b.l1ProviderUrl = url
	return w3b
}

// SetL2ContractBackend set Polygon backend for communication with Web3Domain registry
func (w3b *web3ClientBuilder) SetMaticContractBackend(backend bind.ContractBackend) Web3DomainBuilder {
	w3b.l2ContractBackend = backend
	return w3b
}

// SetL2ContractBackendProviderUrl set Polygon backend Rpc URL
func (w3b *web3ClientBuilder) SetMaticContractBackendProviderUrl(url string) Web3DomainBuilder {
	w3b.l2ProviderUrl = url
	return w3b
}

func (w3b *web3ClientBuilder) SetMetadataClient(client MetadataClient) Web3DomainBuilder {
	w3b.metadataClient = client
	return w3b
}

func (w3b *web3ClientBuilder) SetUdClient(apiKey string) Web3DomainBuilder {
	client, err := udclient.Dial(apiKey)

	if err != nil {
		panic(err)
	}

	w3b.l1ContractBackend = client.L1ContractBackend
	w3b.l2ContractBackend = client.L2ContractBackend
	return w3b
}

func (w3b *web3ClientBuilder) SetEthereumNetwork(network string) Web3DomainBuilder {
	w3b.l1Network = network
	return w3b
}

func (w3b *web3ClientBuilder) SetMaticNetwork(network string) Web3DomainBuilder {
	w3b.l2Network = network
	return w3b
}

func (w3b *web3ClientBuilder) Build() (*Web3Domain, error) {

	if w3b.metadataClient == nil {
		w3b.metadataClient = &http.Client{}
	}

	if w3b.l1Network == "" {
		return nil, &UnsConfigurationError{Layer: Layer1, InvalidField: "network"}
	}
	if w3b.l2Network == "" {
		return nil, &UnsConfigurationError{Layer: Layer2, InvalidField: "network"}
	}

	if w3b.l1ContractBackend == nil && w3b.l1ProviderUrl == "" {
		return nil, &UnsConfigurationError{Layer: Layer1, InvalidField: "contractBackend"}
	}

	if w3b.l2ContractBackend == nil && w3b.l2ProviderUrl == "" {
		return nil, &UnsConfigurationError{Layer: Layer2, InvalidField: "contractBackend"}
	}

	uns, err := NewUnsBuilder().
		SetContractBackendProviderUrl(w3b.l1ProviderUrl).
		SetContractBackend(w3b.l1ContractBackend).
		SetL2ContractBackendProviderUrl(w3b.l2ProviderUrl).
		SetL2ContractBackend(w3b.l2ContractBackend).
		SetMetadataClient(w3b.metadataClient).
		SetEthereumNetwork(w3b.l1Network).
		Build()

	if err != nil {
		return nil, err
	}

	ens, err := NewEnsBuilder().SetContractBackendProviderUrl(w3b.l1ProviderUrl).SetContractBackend(w3b.l1ContractBackend).SetEthereumNetwork(w3b.l1Network).SetMetadataClient(w3b.metadataClient).Build()

	if err != nil {
		return nil, err
	}

	return &Web3Domain{
		uns,
		ens,
	}, nil
}
