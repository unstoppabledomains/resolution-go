package resolution

import (
	"net/http"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func getL1TestProviderUrl() string {
	if os.Getenv("L1_TEST_NET_RPC_URL") != "" {
		return os.Getenv("L1_TEST_NET_RPC_URL")
	}

	panic("L1_TEST_NET_RPC_URL is not set!")
}

func getL2TestProviderUrl() string {
	if os.Getenv("L2_TEST_NET_RPC_URL") != "" {
		return os.Getenv("L2_TEST_NET_RPC_URL")
	}

	panic("L2_TEST_NET_RPC_URL is not set!")
}

func TestUnsBuilderSetBackend(t *testing.T) {
	t.Parallel()

	backendl1, _ := ethclient.Dial(getL1TestProviderUrl())
	backendl2, _ := ethclient.Dial(getL2TestProviderUrl())
	builder := NewUnsBuilder().SetEthereumNetwork("goerli").SetL2EthereumNetwork("mumbai")
	builder.SetContractBackend(backendl1)
	builder.SetL2ContractBackend(backendl2)
	uns, err := builder.Build()
	assert.Nil(t, err)
	assert.Equal(t, backendl1, uns.l1Service.contractBackend)
	assert.Equal(t, backendl2, uns.l2Service.contractBackend)
}

func TestUnsBuilderSetProviderUrl(t *testing.T) {
	t.Parallel()

	builder := NewUnsBuilder()
	builder = builder.SetEthereumNetwork("goerli").SetContractBackendProviderUrl(getL1TestProviderUrl())
	builder = builder.SetL2EthereumNetwork("mumbai").SetL2ContractBackendProviderUrl(getL2TestProviderUrl())
	uns, err := builder.Build()
	assert.Nil(t, err)
	assert.NotNil(t, uns)
}

func TestUnsBuilderSetProxyBackend(t *testing.T) {
	t.Parallel()

	builder := NewUnsBuilder().SetUdClient("test")
	uns, err := builder.Build()

	assert.Nil(t, err)
	assert.Equal(t, uns.l1Service.networkId, 1)
	assert.Equal(t, uns.l2Service.networkId, 137)
}

func TestUnsBuilderSetMetadataClient(t *testing.T) {
	t.Parallel()
	client := &http.Client{}

	backendl1, _ := ethclient.Dial(getL1TestProviderUrl())
	backendl2, _ := ethclient.Dial(getL2TestProviderUrl())
	builder := NewUnsBuilder().SetEthereumNetwork("goerli").SetL2EthereumNetwork("mumbai")
	builder.SetContractBackend(backendl1)
	builder.SetL2ContractBackend(backendl2)

	builder.SetMetadataClient(client)
	uns, err := builder.Build()
	assert.Nil(t, err)
	assert.Equal(t, client, uns.l1Service.metadataClient)
	assert.Equal(t, client, uns.l2Service.metadataClient)
}

func TestUnsBuilderChecksContractBackend(t *testing.T) {
	t.Parallel()

	var expectedError *UnsConfigurationError

	builder := NewUnsBuilder().SetEthereumNetwork("goerli").SetL2EthereumNetwork("mumbai")
	uns, err := builder.Build()
	assert.Nil(t, uns)
	assert.NotNil(t, err)
	assert.ErrorAs(t, err, &expectedError)
}

func TestUnsBuilderChecksL2ContractBackend(t *testing.T) {
	t.Parallel()
	var expectedError *UnsConfigurationError

	backendl1, _ := ethclient.Dial(getL1TestProviderUrl())
	builder := NewUnsBuilder().SetEthereumNetwork("goerli").SetL2EthereumNetwork("mumbai")
	builder.SetContractBackend(backendl1)
	_, err := builder.Build()
	assert.ErrorAs(t, err, &expectedError)
	assert.Equal(t, "Invalid UNS configuration value of contractBackend for Layer 2", err.Error())
}

func TestUnsBuilderChecksL1ContractBackend(t *testing.T) {
	t.Parallel()
	var expectedError *UnsConfigurationError

	backendl2, _ := ethclient.Dial(getL2TestProviderUrl())
	builder := NewUnsBuilder().SetEthereumNetwork("goerli").SetL2EthereumNetwork("mumbai")
	builder.SetL2ContractBackend(backendl2)
	_, err := builder.Build()
	assert.ErrorAs(t, err, &expectedError)
	assert.Equal(t, "Invalid UNS configuration value of contractBackend for Layer 1", err.Error())
}
