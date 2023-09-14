package resolution

import (
	"reflect"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/unstoppabledomains/resolution-go/v3/namingservice"
)

func TestEnforceImplementInterface(t *testing.T) {
	t.Parallel()
	assert.Implements(t, (*NamingService)(nil), &Zns{provider: nil})
	assert.Implements(t, (*NamingService)(nil), &Uns{
		l1Service: UnsService{},
		l2Service: UnsService{},
	})
}

func TestDetectNamingServiceType(t *testing.T) {
	t.Parallel()
	var serviceType string
	serviceType, err := DetectNamingService("test.zil")
	assert.Nil(t, err)
	assert.Equal(t, namingservice.ZNS, serviceType)

	serviceType, err = DetectNamingService("test.crypto")
	assert.Nil(t, err)
	assert.Equal(t, namingservice.UNS, serviceType)

	serviceType, err = DetectNamingService("test.asdasdas")
	assert.Nil(t, err)
	assert.Equal(t, namingservice.UNS, serviceType)

	serviceType, err = DetectNamingService("test.wallet")
	assert.Nil(t, err)
	assert.Equal(t, namingservice.UNS, serviceType)
}

func TestDetectNamingServiceTypeInvalidDomain(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotSupportedError
	_, err := DetectNamingService("aaaazzsd..")
	assert.ErrorAs(t, err, &expectedError)

	_, err = DetectNamingService("aaaazzsd")
	assert.ErrorAs(t, err, &expectedError)
}

var testWeb3Domains = map[string]DomainData{
	"UnsDomain": {Name: "reseller-test-udtesting-459239285.crypto"},
	"EnsDomain": {Name: "test.eth"},
}

func getWeb3Domain() *Web3Domain {
	builder := NewWeb3DomainBuilder()
	builder = builder.SetEthereumNetwork("goerli").SetEthContractBackendProviderUrl(getL1TestProviderUrl())
	builder = builder.SetMaticNetwork("mumbai").SetMaticContractBackendProviderUrl(getL2TestProviderUrl())
	web3Domain, _ := builder.Build()
	return web3Domain
}

func TestWeb3DomainBuilderWithUrl(t *testing.T) {
	t.Parallel()
	builder := NewWeb3DomainBuilder()
	builder = builder.SetEthereumNetwork("goerli").SetEthContractBackendProviderUrl(getL1TestProviderUrl())
	builder = builder.SetMaticNetwork("mumbai").SetMaticContractBackendProviderUrl(getL2TestProviderUrl())
	web3Domain, _ := builder.Build()
	assert.NotNil(t, web3Domain)
}

func TestWeb3DomainBuilderWithContractBackend(t *testing.T) {
	builder := NewWeb3DomainBuilder()
	ethBackend, _ := ethclient.Dial(getL1TestProviderUrl())
	backendL2, _ := ethclient.Dial(getL2TestProviderUrl())
	builder = builder.SetEthContractBackend(ethBackend)
	builder = builder.SetMaticContractBackend(backendL2)
	web3Domain, _ := builder.Build()
	assert.NotNil(t, web3Domain)
}

func TestWeb3DomainBuilderWithApiKey(t *testing.T) {
	builder := NewWeb3DomainBuilder().SetUdClient("some key")
	web3Domain, _ := builder.Build()
	assert.NotNil(t, web3Domain)
}

func TestWeb3DomainGetNamingServiceForUnsDomain(t *testing.T) {
	t.Parallel()

	namingservice := getWeb3Domain().getNamingServiceForDomain(testWeb3Domains["UnsDomain"].Name)
	assert.NotNil(t, namingservice)
	assert.Equal(t, reflect.TypeOf(namingservice) == reflect.TypeOf(&Uns{}), true)
}

func TestWeb3DomainGetNamingServiceForEnsDomain(t *testing.T) {
	t.Parallel()
	namingservice := getWeb3Domain().getNamingServiceForDomain(testWeb3Domains["EnsDomain"].Name)
	assert.NotNil(t, namingservice)
	assert.Equal(t, reflect.TypeOf(namingservice) == reflect.TypeOf(&Ens{}), true)
}

func TestWeb3DomainDomainExpiryForUnsDomain(t *testing.T) {
	t.Parallel()
	expiry, _ := getWeb3Domain().DomainExpiry(testWeb3Domains["UnsDomain"].Name)
	assert.Equal(t, expiry.After(time.Now().Add(99)), true)
}

func TestWeb3DomainDomainExpiryForEnsDomain(t *testing.T) {
	t.Parallel()
	expiry, err := getWeb3Domain().DomainExpiry(testWeb3Domains["EnsDomain"].Name)

	assert.Nil(t, err)
	assert.Equal(t, expiry.After(time.Now().Add(1)), true) // assuming this  domain last for years
}

func TestWeb3DomainNameHashForEthDomain(t *testing.T) {
	t.Parallel()

	nameHash, _ := getWeb3Domain().Namehash(testWeb3Domains["EnsDomain"].Name)
	assert.Equal(t, nameHash, "0xeb4f647bea6caa36333c816d7b46fdcb05f9466ecacc140ea8c66faf15b3d9f1")
}

func TestWeb3DomainNameHashForEthSubDomain(t *testing.T) {
	t.Parallel()

	nameHash, _ := getWeb3Domain().Namehash("test1.test.eth")
	assert.Equal(t, nameHash, "0x2fae416d6dd9ac2a3d1035f5de2c4ee6905fd34bf1918c7a497cafbc696b20e7")
}

func TestWeb3DomainGetEnsResolver(t *testing.T) {
	t.Parallel()

	resolver, err := getWeb3Domain().Resolver(testWeb3Domains["EnsDomain"].Name)

	assert.Nil(t, err)
	assert.Equal(t, resolver, "0x4B1488B7a6B320d2D721406204aBc3eeAa9AD329")
}

func TestWeb3DomainGetEnsOwner(t *testing.T) {
	t.Parallel()

	owner, err := getWeb3Domain().Owner(testWeb3Domains["EnsDomain"].Name)

	assert.Nil(t, err)
	assert.Equal(t, owner, "0x145E8aa4ECff3Bdea8d98739105AE038f1f0E352")
}

func TestWeb3DomainGetEnsReverse(t *testing.T) {
	t.Parallel()

	domain, err := getWeb3Domain().ReverseOf("0x2ba39217d82d0d4b1ff90125909973a8572fb2a4")

	assert.Nil(t, err)
	assert.Equal(t, domain, "ppi.eth")
}

func TestWeb3DomainGetNoReverseResult(t *testing.T) {
	t.Parallel()

	domain, err := getWeb3Domain().ReverseOf("0x2ba39217d82d0d4b1ff90125909973a8572fb2a1")

	assert.Nil(t, err)
	assert.Equal(t, domain, "")
}

func TestWeb3DomainGetEnsTokenUri(t *testing.T) {
	t.Parallel()

	tokenUri, err := getWeb3Domain().TokenURI(testWeb3Domains["EnsDomain"].Name)

	assert.Nil(t, err)
	assert.Equal(t, tokenUri, "https://metadata.ens.domains/mainnet/0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85/0x9c22ff5f21f0b81b113e63f7db6da94fedef11b2119b4088b89664fb9a3cb658")
}
