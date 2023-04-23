package resolution

import (
	"testing"

	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/stretchr/testify/assert"
	"github.com/unstoppabledomains/resolution-go/v3/resolution/dnsrecords"
)

var zns, _ = NewZnsBuilder().SetProvider(provider.NewProvider("https://dev-api.zilliqa.com")).SetNetwork("testnet").Build()

func TestZnsBuilder(t *testing.T) {
	t.Parallel()
	_, err := NewZnsBuilder().Build()
	assert.Nil(t, err)
}

func TestZnsBuilderSetProvider(t *testing.T) {
	t.Parallel()
	znsProvider := provider.NewProvider("https://dev-api.zilliqa.com")
	builder := NewZnsBuilder()
	builder.SetProvider(znsProvider)
	znsService, err := builder.Build()
	assert.Nil(t, err)
	assert.Equal(t, znsProvider, znsService.provider)
}

func TestZnsStateDomainNotRegistered(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegisteredError
	_, err := zns.State("long-not-registered-name.zil")
	assert.ErrorAs(t, err, &expectedError)
}

func TestZnsStateDomainNotConfigured(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotConfiguredError
	_, err := zns.State("unconfigured-domain.zil")
	assert.ErrorAs(t, err, &expectedError)
}

func TestZnsRecords(t *testing.T) {
	t.Parallel()
	expectedRecords := map[string]string{
		"ipfs.html.value":    "QmVaAtQbi3EtsfpKoLzALm6vXphdi2KjMgxEDKeGg6wHuK",
		"crypto.BCH.address": "qrq4sk49ayvepqz7j7ep8x4km2qp8lauvcnzhveyu6",
	}
	records, err := zns.Records("testing.zil", []string{"ipfs.html.value", "crypto.BCH.address"})
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, records)
}

func TestZnsNoRecords(t *testing.T) {
	t.Parallel()
	records, err := zns.Records("testing.zil", []string{})
	assert.Nil(t, err)
	assert.Empty(t, records)
}

func TestZnsEmptyRecords(t *testing.T) {
	t.Parallel()
	expectedRecords := map[string]string{
		"ipfs.html.value":    "QmVaAtQbi3EtsfpKoLzALm6vXphdi2KjMgxEDKeGg6wHuK",
		"crypto.BCH.address": "qrq4sk49ayvepqz7j7ep8x4km2qp8lauvcnzhveyu6",
		"key-not-exist":      "",
	}
	records, err := zns.Records("testing.zil", []string{"ipfs.html.value", "crypto.BCH.address", "key-not-exist"})
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, records)
}

func TestZnsRecord(t *testing.T) {
	t.Parallel()
	expectedRecord := "0x45b31e01AA6f42F0549aD482BE81635ED3149abb"
	record, err := zns.Record("testing.zil", "crypto.ETH.address")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestZnsEmptyRecord(t *testing.T) {
	t.Parallel()
	record, err := zns.Record("testing.zil", "non-existent-key")
	assert.Nil(t, err)
	assert.Empty(t, record)
}

func TestZnsOwner(t *testing.T) {
	t.Parallel()
	expectedOwner := "0x003e3cdfeceae96efe007f8196a1b1b1df547eee"
	owner, err := zns.Owner("testing.zil")
	assert.Nil(t, err)
	assert.Equal(t, expectedOwner, owner)
}

func TestZnsResolver(t *testing.T) {
	t.Parallel()
	expectedResolver := "0x02621c64a57e1424adfe122569f2356145f05d4f"
	resolver, err := zns.Resolver("testing.zil")
	assert.Nil(t, err)
	assert.Equal(t, expectedResolver, resolver)
}

func TestZnsAddr(t *testing.T) {
	t.Parallel()
	expectedRecord := "0x45b31e01AA6f42F0549aD482BE81635ED3149abb"
	record, err := zns.Addr("testing.zil", "ETH")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestZnsAddrVersion(t *testing.T) {
	t.Parallel()
	expectedRecord := "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
	record, err := zns.AddrVersion("testing.zil", "USDT", "ERC20")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestZnsEmail(t *testing.T) {
	t.Parallel()
	expectedEmail := "derainberk@gmail.com"
	email, err := zns.Email("testing.zil")
	assert.Nil(t, err)
	assert.Equal(t, expectedEmail, email)
}

func TestZnsIpfs(t *testing.T) {
	t.Parallel()
	testDomain := "testing.zil"
	expectedRecord := "QmVaAtQbi3EtsfpKoLzALm6vXphdi2KjMgxEDKeGg6wHuK"
	record, err := zns.IpfsHash(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestZnsHTTPUrl(t *testing.T) {
	t.Parallel()
	testDomain := "testing.zil"
	expectedRecord := "www.unstoppabledomains.com"
	record, err := zns.HTTPUrl(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestZnsDns(t *testing.T) {
	t.Parallel()
	testDomain := "testing.zil"
	dnsRecords, err := zns.DNS(testDomain, []dnsrecords.Type{"A"})
	assert.Nil(t, err)
	assert.Empty(t, dnsRecords)
}

func TestZnsEmptyDns(t *testing.T) {
	t.Parallel()
	testDomain := "testing.zil"
	dnsRecords, err := zns.DNS(testDomain, []dnsrecords.Type{})
	assert.Nil(t, err)
	assert.Empty(t, dnsRecords)
}

func TestZnsIsSupportedDomain(t *testing.T) {
	t.Parallel()

	isSupportedDomain := func(domain string) bool {
		isSupported, _ := zns.IsSupportedDomain(domain)
		return isSupported
	}

	assert.True(t, isSupportedDomain("valid.zil"))
	assert.False(t, isSupportedDomain("valid.crypto"))
	assert.False(t, isSupportedDomain("invalid.com"))
}

func TestZnsTokenUriIsNotSupported(t *testing.T) {
	t.Parallel()
	var expectedError *MethodIsNotSupportedError
	_, err := zns.TokenURI("testing.zil")
	assert.ErrorAs(t, err, &expectedError)
}

func TestZnsTokenUriMetadataIsNotSupported(t *testing.T) {
	t.Parallel()
	var expectedError *MethodIsNotSupportedError
	_, err := zns.TokenURIMetadata("testing.zil")
	assert.ErrorAs(t, err, &expectedError)
}

func TestZnsUnhashIsNotSupported(t *testing.T) {
	t.Parallel()
	var expectedError *MethodIsNotSupportedError
	_, err := zns.Unhash("testing.zil")
	assert.ErrorAs(t, err, &expectedError)
}
