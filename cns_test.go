package resolution

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/unstoppabledomains/resolution-go/dnsrecords"
)

type MockedMetadataClient struct {
	Response *http.Response
	Err      error
}

func (m *MockedMetadataClient) SetResponse(resp *http.Response) *MockedMetadataClient {
	m.Response = resp
	return m
}

func (m *MockedMetadataClient) SetError(err error) *MockedMetadataClient {
	m.Err = err
	return m
}

func (m *MockedMetadataClient) Get(_ string) (resp *http.Response, err error) {
	return m.Response, m.Err
}

var cns, _ = NewCnsBuilder().Build()

func TestCnsBuilder(t *testing.T) {
	t.Parallel()
	builder := NewCnsBuilder()
	_, err := builder.Build()
	assert.Nil(t, err)
	assert.NotNil(t, cns.contractBackend)
	assert.NotNil(t, cns.metadataClient)
	assert.NotNil(t, cns.supportedKeys)
	assert.NotNil(t, cns.proxyReader)
}

func TestCnsBuilderSetBackend(t *testing.T) {
	t.Parallel()
	backend, _ := ethclient.Dial("https://mainnet.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e")
	builder := NewCnsBuilder()
	builder.SetContractBackend(backend)
	cns, err := builder.Build()
	assert.Nil(t, err)
	assert.Equal(t, backend, cns.contractBackend)
}

func TestCnsBuilderSetMetadataClient(t *testing.T) {
	t.Parallel()
	client := &http.Client{}
	builder := NewCnsBuilder()
	builder.SetMetadataClient(client)
	cns, err := builder.Build()
	assert.Nil(t, err)
	assert.Equal(t, client, cns.metadataClient)
}

func TestNewCnsWithSupportedKeys(t *testing.T) {
	t.Parallel()
	cnsService, _ := NewCnsBuilder().Build()
	deprecatedKeyName := cnsService.supportedKeys["crypto.ETH.address"]
	assert.Equal(t, "ETH", deprecatedKeyName.DeprecatedKeyName)
}

func TestCnsDataValue(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	expectedRecord := "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
	data, err := cns.Data(testDomain, []string{"crypto.ETH.address"})
	assert.Nil(t, err)
	assert.Equal(t, data.Values[0], expectedRecord)
}

func TestCnsData(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	expectedRecord := "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
	expectedOwner := common.HexToAddress("0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2")
	expectedResolver := common.HexToAddress("0xb66DcE2DA6afAAa98F2013446dBCB0f4B0ab2842")
	data, err := cns.Data(testDomain, []string{"crypto.ETH.address"})
	assert.Nil(t, err)
	assert.Equal(t, data.Values[0], expectedRecord)
	assert.Equal(t, expectedOwner, data.Owner)
	assert.Equal(t, expectedResolver, data.Resolver)
}

func TestCnsEmptyDataValues(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	data, _ := cns.Data(testDomain, []string{"empty record"})
	assert.Equal(t, data.Values[0], "")
	assert.Len(t, data.Values, 1)
}

func TestCnsDomainNotRegistered(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegisteredError
	testDomain := "not-registered-long-domain-name.crypto"
	_, err := cns.Data(testDomain, []string{"crypto.ETH.address"})
	assert.ErrorAs(t, err, &expectedError)
}

func TestCnsDomainNotConfigured(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotConfiguredError
	testDomain := "reseller-test-paul2.crypto"
	_, err := cns.Data(testDomain, []string{"crypto.ETH.address"})
	assert.ErrorAs(t, err, &expectedError)
}

func TestCnsRecords(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	expectedRecords := map[string]string{"crypto.ETH.address": "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8", "crypto.BTC.address": "bc1q359khn0phg58xgezyqsuuaha28zkwx047c0c3y"}
	records, err := cns.Records(testDomain, []string{"crypto.ETH.address", "crypto.BTC.address"})
	assert.Nil(t, err)
	assert.Equal(t, records, expectedRecords)
}

func TestCnsNoRecords(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	records, err := cns.Records(testDomain, []string{})
	assert.Nil(t, err)
	assert.Empty(t, records)
}

func TestCnsEmptyRecords(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	expectedRecords := map[string]string{"record-not-exist": "", "crypto.ETH.address": "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8", "crypto.BTC.address": "bc1q359khn0phg58xgezyqsuuaha28zkwx047c0c3y"}
	records, err := cns.Records(testDomain, []string{"record-not-exist", "crypto.ETH.address", "crypto.BTC.address"})
	assert.Nil(t, err)
	assert.Equal(t, records, expectedRecords)
}

func TestCnsRecord(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	expectedRecord := "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
	record, err := cns.Record(testDomain, "crypto.ETH.address")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestCnsEmptyRecord(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	record, err := cns.Record(testDomain, "record-not-exist")
	assert.Nil(t, err)
	assert.Empty(t, record)
}

func TestCnsAddr(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	expectedRecord := "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
	record, err := cns.Addr(testDomain, "ETH")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestCnsAddrLowerCaseTicker(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	expectedRecord := "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
	record, err := cns.Addr(testDomain, "eth")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestCnsEmail(t *testing.T) {
	t.Parallel()
	testDomain := "reseller-test-paul019.crypto"
	expectedRecord := "paul@unstoppabledomains.com"
	record, err := cns.Email(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestCnsResolver(t *testing.T) {
	t.Parallel()
	testDomain := "reseller-test-mago017.crypto"
	expectedRecord := "0x878bC2f3f717766ab69C0A5f9A6144931E61AEd3"
	record, err := cns.Resolver(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestCnsOwner(t *testing.T) {
	t.Parallel()
	testDomain := "reseller-test-paul019.crypto"
	expectedRecord := "0xA1cAc442Be6673C49f8E74FFC7c4fD746f3cBD0D"
	record, err := cns.Resolver(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestCnsAddrVersion(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-usdt.crypto"
	expectedRecord := "0xe7474D07fD2FA286e7e0aa23cd107F8379085037"
	record, err := cns.AddrVersion(testDomain, "USDT", "ERC20")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestCnsIpfs(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-ipfs.crypto"
	expectedRecord := "QmVJ26hBrwwNAPVmLavEFXDUunNDXeFSeMPmHuPxKe6dJv"
	record, err := cns.IpfsHash(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestCnsIpfsLegacy(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-ipfs-legacy.crypto"
	expectedRecord := "QmT9qk3CRYbFDWpDFYeAv8T8H1gnongwKhh5J68NLkLir6"
	record, err := cns.IpfsHash(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestCnsHTTPUrl(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-redirect.crypto"
	expectedRecord := "https://example.com/home.html"
	record, err := cns.HTTPUrl(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestCnsHttpUrlLegacy(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-redirect-legacy.crypto"
	expectedRecord := "https://legacy-example.com/home.html"
	record, err := cns.HTTPUrl(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestCnsAllRecords(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	expectedRecords := map[string]string{
		"crypto.BTC.address":         "bc1q359khn0phg58xgezyqsuuaha28zkwx047c0c3y",
		"crypto.ETH.address":         "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8",
		"ipfs.html.value":            "Qme54oEzRkgooJbCDr78vzKAWcv6DDEZqRhhDyDtzgrZP6",
		"ipfs.redirect_domain.value": "https://abbfe6z95qov3d40hf6j30g7auo7afhp.mypinata.cloud/ipfs/Qme54oEzRkgooJbCDr78vzKAWcv6DDEZqRhhDyDtzgrZP6",
	}
	allRecords, err := cns.AllRecords(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, allRecords)
}

func TestCnsAllRecordsStandardKeys(t *testing.T) {
	t.Parallel()
	testDomain := "monmouthcounty.crypto"
	expectedRecords := map[string]string{
		"crypto.BTC.address":         "3NwuV8nVT2VKbtCs8evChdiW6kHTHcVpdn",
		"crypto.ETH.address":         "0x1C42088b82f6Fa5fB883A14240C4E066dDFf1517",
		"crypto.LTC.address":         "MTnTNwKikiMi97Teq8XQRabL9SZ4HjnKNB",
		"crypto.ADA.address":         "DdzFFzCqrhsfc3MQvjsLr9BHkaFYeE7BotyTATdETRoSPj6QPiotK4xpcFZk66KVmtr87tvUFTcbTHZRkcdbMR5Ss6jCfzCVtFRMB7WE",
		"ipfs.html.value":            "QmYqX8D8SkaF5YcpaWMyi5xM43UEteFiSNKYsjLcdvCWud",
		"ipfs.redirect_domain.value": "https://abbfe6z95qov3d40hf6j30g7auo7afhp.mypinata.cloud/ipfs/QmYqX8D8SkaF5YcpaWMyi5xM43UEteFiSNKYsjLcdvCWud",
	}
	allRecords, err := cns.AllRecords(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, allRecords)
}

func TestCnsDnsA(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-dns.crypto"
	expectedRecords := []dnsrecords.Record{
		{Type: "A", TTL: 1800, Value: "10.0.0.1"},
		{Type: "A", TTL: 1800, Value: "10.0.0.2"},
	}
	dnsRecords, err := cns.DNS(testDomain, []dnsrecords.Type{"A"})
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedRecords, dnsRecords)
}

func TestCnsDnsCname(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-dns-cname.crypto"
	expectedRecords := []dnsrecords.Record{
		{Type: "CNAME", TTL: 1111, Value: "example.com."},
	}
	dnsRecords, err := cns.DNS(testDomain, []dnsrecords.Type{"CNAME"})
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedRecords, dnsRecords)
}

func TestCnsDnsGlobalTtl(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-dns-global-ttl.crypto"
	expectedRecords := []dnsrecords.Record{
		{Type: "A", TTL: 1000, Value: "10.0.0.1"},
		{Type: "A", TTL: 1000, Value: "10.0.0.2"},
	}
	dnsRecords, err := cns.DNS(testDomain, []dnsrecords.Type{"A"})
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedRecords, dnsRecords)
}

func TestCnsIsSupportedDomain(t *testing.T) {
	t.Parallel()

	isSupportedDomain := func(domain string) bool {

		isSupported, _ := cns.IsSupportedDomain(domain)
		return isSupported
	}

	assert.True(t, isSupportedDomain("valid.crypto"))
	assert.False(t, isSupportedDomain("invalid.zil"))
	assert.True(t, isSupportedDomain("invalid.com"))
	assert.True(t, isSupportedDomain("radomin-domain.com"))
	assert.True(t, isSupportedDomain("some-domain.net"))
	assert.True(t, isSupportedDomain("some-domain.wiowejfo.qwefwef"))
	assert.True(t, isSupportedDomain("some-domain.wiowejfo.qwd"))
	assert.False(t, isSupportedDomain("some-domain.wiowejfo.zil"))
}

func TestCnsUnsupportedDomainError(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotSupportedError
	_, err := cns.Data("invalid.zil", []string{"crypto.ETH.address"})
	assert.ErrorAs(t, err, &expectedError)
}

func TestCnsTokenURI(t *testing.T) {
	t.Parallel()
	tokenURI, err := cns.TokenURI("udtestdev-test.crypto")
	expectedTokenURI := "https://metadata.unstoppabledomains.com/metadata/udtestdev-test.crypto"
	assert.Nil(t, err)
	assert.Equal(t, expectedTokenURI, tokenURI)
}

func TestCnsTokenURIDomainIsNotRegistered(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegisteredError
	_, err := cns.TokenURI("unregistered-domain-name.crypto")
	assert.ErrorAs(t, err, &expectedError)
}

func TestCnsTokenUriNotSupportedDomain(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotSupportedError
	_, err := cns.TokenURI("invalid.zil")
	assert.ErrorAs(t, err, &expectedError)
}

func TestCnsTokenURIMetadata(t *testing.T) {
	t.Parallel()
	expectedMetadata := TokenMetadata{
		Name:        "udtestdev-test.crypto",
		Description: "A .crypto blockchain domain. Use it to resolve your cryptocurrency addresses and decentralized websites.\nhttps://gateway.pinata.cloud/ipfs/Qme54oEzRkgooJbCDr78vzKAWcv6DDEZqRhhDyDtzgrZP6",
		ExternalUrl: "https://unstoppabledomains.com/search?searchTerm=udtestdev-test.crypto",
		Image:       "https://storage.googleapis.com/dot-crypto-metadata-api/unstoppabledomains_crypto.png",
		Attributes: []TokenMetadataAttribute{
			{
				TraitType: "domain",
				Value:     "udtestdev-test.crypto",
			},
		},
	}
	metadata, err := cns.TokenURIMetadata("udtestdev-test.crypto")
	assert.Nil(t, err)
	assert.Equal(t, expectedMetadata.Name, metadata.Name)
	assert.Equal(t, expectedMetadata.Description, metadata.Description)
	assert.Equal(t, expectedMetadata.ExternalUrl, metadata.ExternalUrl)
	assert.Equal(t, expectedMetadata.Image, metadata.Image)
	assert.Contains(t, metadata.Attributes, expectedMetadata.Attributes[0])
}

func TestCnsTokenURIMetadataNotSupportedDomain(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegisteredError
	_, err := cns.TokenURIMetadata("unregistered-domain-name.crypto")
	assert.ErrorAs(t, err, &expectedError)
}

func TestCnsUnhash(t *testing.T) {
	t.Parallel()
	expectedDomainName := "ryan.crypto"
	domainName, err := cns.Unhash("0x691f36df38168d9297e784f45a87257a70c58c4040d469c6d0b91d253a837e32")
	assert.Nil(t, err)
	assert.Equal(t, expectedDomainName, domainName)
}

func TestCnsUnhashWithout0xPrefix(t *testing.T) {
	t.Parallel()
	expectedDomainName := "ryan.crypto"
	domainName, err := cns.Unhash("691f36df38168d9297e784f45a87257a70c58c4040d469c6d0b91d253a837e32")
	assert.Nil(t, err)
	assert.Equal(t, expectedDomainName, domainName)
}

func TestCnsUnhashInvalidDomain(t *testing.T) {
	t.Parallel()
	var expectedError *InvalidDomainNameReturnedError
	body, _ := json.Marshal(TokenMetadata{
		Name:            "brad.crypto",
		Description:     "",
		Image:           "",
		ExternalUrl:     "",
		ExternalLink:    "",
		ImageData:       "",
		BackgroundColor: "",
		AnimationUrl:    "",
		YoutubeUrl:      "",
		Attributes:      nil,
	})
	var mockedClient MockedMetadataClient
	mockedClient.SetResponse(&http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer(body)),
	})
	mockedClient.SetError(nil)
	cnsWithMockedMetadataClient, _ := NewCnsBuilder().SetMetadataClient(&mockedClient).Build()
	domainName, err := cnsWithMockedMetadataClient.Unhash("691f36df38168d9297e784f45a87257a70c58c4040d469c6d0b91d253a837e32")
	assert.Empty(t, domainName)
	assert.ErrorAs(t, err, &expectedError)
}
