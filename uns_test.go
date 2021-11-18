package resolution

import (
	"net/http"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/unstoppabledomains/resolution-go/dnsrecords"
	"github.com/unstoppabledomains/resolution-go/namingservice"
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

var uns, _ = NewUnsBuilder().SetEthereumNetwork("rinkeby").SetL2EthereumNetwork("mumbai").Build()

func TestNewUnsWithSupportedKeys(t *testing.T) {
	t.Parallel()
	unsService, _ := NewUnsBuilder().SetEthereumNetwork("rinkeby").SetL2EthereumNetwork("mumbai").Build()
	deprecatedKeyName := unsService.l1Service.supportedKeys["crypto.ETH.address"]
	assert.Equal(t, "ETH", deprecatedKeyName.DeprecatedKeyName)
}

func TestUnsDataValue(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecord := "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2"
	data, err := uns.Data(testDomain, []string{"crypto.ETH.address"})
	assert.Nil(t, err)
	assert.Equal(t, data.Values[0], expectedRecord)
}

func TestUnsL2DataValue(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test-l2-domain-784391.wallet"
	expectedRecord := "0x6A1fd9a073256f14659fe59613bbf169Ed27CdcC"
	data, err := uns.Data(testDomain, []string{"crypto.LINK.address"})
	assert.Nil(t, err)
	assert.Equal(t, data.Values[0], expectedRecord)
}

func TestUnsData(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecord := "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2"
	expectedOwner := common.HexToAddress("0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2")
	expectedResolver := common.HexToAddress("0x95AE1515367aa64C462c71e87157771165B1287A")
	data, err := uns.Data(testDomain, []string{"crypto.ETH.address"})
	assert.Nil(t, err)
	assert.Equal(t, data.Values[0], expectedRecord)
	assert.Equal(t, expectedOwner, data.Owner)
	assert.Equal(t, expectedResolver, data.Resolver)
}

func TestUnsL2Data(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test-l2-domain-784391.wallet"
	expectedRecord := "0x6A1fd9a073256f14659fe59613bbf169Ed27CdcC"
	expectedOwner := common.HexToAddress("0x499dd6d875787869670900a2130223d85d4f6aa7")
	expectedResolver := common.HexToAddress("0x2a93C52E7B6E7054870758e15A1446E769EdfB93")
	data, err := uns.Data(testDomain, []string{"crypto.LINK.address"})
	assert.Nil(t, err)
	assert.Equal(t, data.Values[0], expectedRecord)
	assert.Equal(t, expectedOwner, data.Owner)
	assert.Equal(t, expectedResolver, data.Resolver)
}

func TestUnsEmptyDataValues(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	data, _ := uns.Data(testDomain, []string{"empty record"})
	assert.Equal(t, data.Values[0], "")
	assert.Len(t, data.Values, 1)
}

func TestUnsL2EmptyDataValues(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test-l2-domain-784391.wallet"
	data, _ := uns.Data(testDomain, []string{"empty record"})
	assert.Equal(t, data.Values[0], "")
	assert.Len(t, data.Values, 1)
}

func TestUnsDomainNotRegistered(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegisteredError
	testDomain := "not-registered-long-domain-name.crypto"
	_, err := uns.Data(testDomain, []string{"crypto.ETH.address"})
	assert.ErrorAs(t, err, &expectedError)
}

func TestUnsDomainNotConfigured(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotConfiguredError
	testDomain := "udtestdev-d0137c.crypto"
	_, err := uns.Data(testDomain, []string{"crypto.XLM.address"})
	assert.ErrorAs(t, err, &expectedError)
}

func TestUnsRecords(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	expectedRecords := map[string]string{"crypto.ETH.address": "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2", "crypto.BTC.address": ""}
	records, err := uns.Records(testDomain, []string{"crypto.ETH.address", "crypto.BTC.address"})
	assert.Nil(t, err)
	assert.Equal(t, records, expectedRecords)
}

func TestUnsL2Records(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test-l2-domain-784391.wallet"
	expectedRecords := map[string]string{"crypto.LINK.address": "0x6A1fd9a073256f14659fe59613bbf169Ed27CdcC", "crypto.BTC.address": ""}
	records, err := uns.Records(testDomain, []string{"crypto.LINK.address", "crypto.BTC.address"})
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, records)
}

func TestUnsEmptyRecords(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	expectedRecords := map[string]string{"crypto.BTC.address": "", "crypto.ETH.address": "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2", "record-not-exist": ""}
	records, err := uns.Records(testDomain, []string{"record-not-exist", "crypto.ETH.address", "crypto.BTC.address"})
	assert.Nil(t, err)
	assert.Equal(t, records, expectedRecords)
}

func TestUnsRecord(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	expectedRecord := "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2"
	record, err := uns.Record(testDomain, "crypto.ETH.address")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL2Record(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test-l2-domain-784391.wallet"
	expectedRecord := "0x6A1fd9a073256f14659fe59613bbf169Ed27CdcC"
	record, err := uns.Record(testDomain, "crypto.LINK.address")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsEmptyRecord(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	record, err := uns.Record(testDomain, "record-not-exist")
	assert.Nil(t, err)
	assert.Empty(t, record)
}

func TestUnsAddr(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecord := "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2"
	record, err := uns.Addr(testDomain, "ETH")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL2Addr(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test-l2-domain-784391.wallet"
	expectedRecord := "0x6A1fd9a073256f14659fe59613bbf169Ed27CdcC"
	record, err := uns.Addr(testDomain, "LINK")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsAddrLowerCaseTicker(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test.crypto"
	expectedRecord := "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2"
	record, err := uns.Addr(testDomain, "eth")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsEmail(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecord := "testing@example.com"
	record, err := uns.Email(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL2Email(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test-l2-domain-784391.wallet"
	expectedRecord := "l2email@l2mail.mail"
	record, err := uns.Email(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsResolver(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecord := "0x95AE1515367aa64C462c71e87157771165B1287A"
	record, err := uns.Resolver(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL2Resolver(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test-l2-domain-784391.wallet"
	expectedRecord := "0x2a93C52E7B6E7054870758e15A1446E769EdfB93"
	record, err := uns.Resolver(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsOwner(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecord := "0x95AE1515367aa64C462c71e87157771165B1287A"
	record, err := uns.Resolver(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL2Owner(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test-l2-domain-784391.wallet"
	expectedRecord := "0x2a93C52E7B6E7054870758e15A1446E769EdfB93"
	record, err := uns.Resolver(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsAddrVersion(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecord := "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2"
	record, err := uns.AddrVersion(testDomain, "USDT", "ERC20")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL2AddrVersion(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test-l2-domain-morerecords2231.wallet"
	expectedRecord := "0x499dD6D875787869670900a2130223D85d4F6Aa7"
	record, err := uns.AddrVersion(testDomain, "USDT", "ERC20")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsIpfs(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecord := "QmS23QDsc3Y26rUfME32Q7jawTrCH8bTrZ7iW8EGLJYMvD"
	record, err := uns.IpfsHash(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL2Ipfs(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test-l2-domain-784391.wallet"
	expectedRecord := "QmfRXG3CcM1eWiCUA89uzimCvQUnw4HzTKLo6hRZ47PYsN"
	record, err := uns.IpfsHash(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsHTTPUrl(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-redirect.crypto"
	expectedRecord := "https://example.com/home.html"
	record, err := uns.HTTPUrl(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL2HTTPUrl(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test-l2-domain-morerecords2231.wallet"
	expectedRecord := "https://L2.example.com/home.html"
	record, err := uns.HTTPUrl(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsHttpUrlLegacy(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-redirect-legacy.crypto"
	expectedRecord := "https://legacy-example.com/home.html"
	record, err := uns.HTTPUrl(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestDotCryptoAllRecords(t *testing.T) {
	t.Parallel()
	testDomain := "testing.crypto"
	expectedRecords := map[string]string{"crypto.ETH.address": "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2", "crypto.USDT.version.EOS.address": "karaarishmen", "crypto.USDT.version.ERC20.address": "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2", "crypto.USDT.version.OMNI.address": "1KvzMF2Vjy14d6JGY7dG7vjT5kfpmzSQXM", "crypto.USDT.version.TRON.address": "TRMJfXXbmwb3WFSRKbeRgKsYoD8o1a9xxV", "dns.A": "[\"10.0.0.1\", \"10.0.0.3\"]", "dns.A.ttl": "98", "dns.AAAA": "[]", "ipfs.html.value": "QmS23QDsc3Y26rUfME32Q7jawTrCH8bTrZ7iW8EGLJYMvD", "whois.email.value": "testing@example.com"}
	allRecords, err := uns.AllRecords(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, allRecords)
}

func TestUnsAllRecords(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-my-new-tls.wallet"
	expectedRecords := map[string]string{"crypto.BTC.address": "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", "crypto.ETH.address": "0x6EC0DEeD30605Bcd19342f3c30201DB263291589"}
	allRecords, err := uns.AllRecords(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, allRecords)
}

func TestUnsL2AllRecords(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-test-l2-domain-784391.wallet"
	expectedRecords := map[string]string{"crypto.LINK.address": "0x6A1fd9a073256f14659fe59613bbf169Ed27CdcC", "dweb.ipfs.hash": "QmfRXG3CcM1eWiCUA89uzimCvQUnw4HzTKLo6hRZ47PYsN", "whois.email.value": "l2email@l2mail.mail"}
	allRecords, err := uns.AllRecords(testDomain)
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
	allRecords, err := uns.AllRecords(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, allRecords)
}

func TestUnsAllRecordsStandardKeys(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-my-new-tls.wallet"

	expectedRecords := map[string]string{"crypto.BTC.address": "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", "crypto.ETH.address": "0x6EC0DEeD30605Bcd19342f3c30201DB263291589"}
	allRecords, err := uns.AllRecords(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, allRecords)
}

func TestUnsDnsA(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-dns.crypto"
	expectedRecords := []dnsrecords.Record{
		{Type: "A", TTL: 1800, Value: "10.0.0.1"},
		{Type: "A", TTL: 1800, Value: "10.0.0.2"},
	}
	dnsRecords, err := uns.DNS(testDomain, []dnsrecords.Type{"A"})
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedRecords, dnsRecords)
}

func TestUnsDnsCname(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-dns-cname.crypto"
	expectedRecords := []dnsrecords.Record{
		{Type: "CNAME", TTL: 1111, Value: "example.com."},
	}
	dnsRecords, err := uns.DNS(testDomain, []dnsrecords.Type{"CNAME"})
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedRecords, dnsRecords)
}

func TestUnsDnsGlobalTtl(t *testing.T) {
	t.Parallel()
	testDomain := "udtestdev-dns-global-ttl.crypto"
	expectedRecords := []dnsrecords.Record{
		{Type: "A", TTL: 1000, Value: "10.0.0.1"},
		{Type: "A", TTL: 1000, Value: "10.0.0.2"},
	}
	dnsRecords, err := uns.DNS(testDomain, []dnsrecords.Type{"A"})
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedRecords, dnsRecords)
}

func TestUnsIsSupportedDomain(t *testing.T) {
	t.Parallel()

	isSupportedDomain := func(domain string) bool {
		isSupported, _ := uns.IsSupportedDomain(domain)
		return isSupported
	}

	assert.True(t, isSupportedDomain("valid.crypto"))
	assert.True(t, isSupportedDomain("valid.qwdqwd.crypto"))
	assert.False(t, isSupportedDomain("invalid.zil"))
	assert.True(t, isSupportedDomain("invalid.wallet"))
	assert.True(t, isSupportedDomain("invalid.bitcoin"))
	assert.True(t, isSupportedDomain("invalid.x"))
	assert.True(t, isSupportedDomain("invalid.888"))
	assert.True(t, isSupportedDomain("invalid.blockchain"))
	assert.True(t, isSupportedDomain("invalid.dao"))
	assert.True(t, isSupportedDomain("invalid.nft"))
	assert.True(t, isSupportedDomain("invalid.coin"))
	assert.False(t, isSupportedDomain("invalid.com"))
	assert.False(t, isSupportedDomain("radomin-domain.com"))
	assert.False(t, isSupportedDomain("some-domain.net"))
	assert.False(t, isSupportedDomain("some-domain.wiowejfo.qwefwef"))
	assert.False(t, isSupportedDomain("some-domain.wiowejfo.qwd"))
	assert.False(t, isSupportedDomain("some-domain.wiowejfo.zil"))
}

func TestUnsDomainNotRegisteredError(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegisteredError
	_, err := uns.Data("invalid.zil", []string{"crypto.ETH.address"})
	assert.ErrorAs(t, err, &expectedError)
}

func TestUnsTokenURI(t *testing.T) {
	t.Parallel()
	tokenURI, err := uns.TokenURI("udtestdev-test.crypto")
	expectedTokenURI := "https://metadata.staging.unstoppabledomains.com/metadata/udtestdev-test.crypto"
	assert.Nil(t, err)
	assert.Equal(t, expectedTokenURI, tokenURI)
}

func TestUnsTokenURIDomainIsNotRegistered(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegisteredError
	_, err := uns.TokenURI("unregistered-domain-name.crypto")
	assert.ErrorAs(t, err, &expectedError)
}

func TestUnsTokenURIZilDomainIsNotRegistered(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegisteredError
	_, err := uns.TokenURI("invalid.zil")
	assert.ErrorAs(t, err, &expectedError)
}

func TestUnsTokenURIMetadata(t *testing.T) {
	t.Parallel()
	expectedMetadata := TokenMetadata{
		Name:        "udtestdev-test.crypto",
		Description: "A CNS or UNS blockchain domain. Use it to resolve your cryptocurrency addresses and decentralized websites.",
		ExternalUrl: "https://unstoppabledomains.com/search?searchTerm=udtestdev-test.crypto",
		Image:       "https://storage.googleapis.com/dot-crypto-metadata-api/images/unstoppabledomains.svg",
		Attributes: []TokenMetadataAttribute{
			{
				TraitType: "domain",
				Value:     "udtestdev-test.crypto",
			},
		},
	}
	metadata, err := uns.TokenURIMetadata("udtestdev-test.crypto")
	assert.Nil(t, err)
	assert.Equal(t, expectedMetadata.Name, metadata.Name)
	assert.Equal(t, expectedMetadata.Description, metadata.Description)
	assert.Equal(t, expectedMetadata.ExternalUrl, metadata.ExternalUrl)
	assert.Equal(t, expectedMetadata.Image, metadata.Image)
	assert.Contains(t, metadata.Attributes, expectedMetadata.Attributes[0])
}

func TestUnsTokenURIMetadataNotSupportedDomain(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegisteredError
	_, err := uns.TokenURIMetadata("very-unregistered-domain-name.crypto")
	assert.ErrorAs(t, err, &expectedError)
}

func TestUnsUnhashDotCrypto(t *testing.T) {
	t.Parallel()
	expectedDomainName := "testing.crypto"
	domainName, err := uns.Unhash("0xd52e0f8bfe7e039fddb362c7e00f3628e2dca805f191d8bef74a07ca0e848245")
	assert.Nil(t, err)
	assert.Equal(t, expectedDomainName, domainName)
}

func TestUnsUnhashWithout0xPrefixDotCrypto(t *testing.T) {
	t.Parallel()
	expectedDomainName := "testing.crypto"
	domainName, err := uns.Unhash("d52e0f8bfe7e039fddb362c7e00f3628e2dca805f191d8bef74a07ca0e848245")
	assert.Nil(t, err)
	assert.Equal(t, expectedDomainName, domainName)
}

func TestUnsUnhashDotWallet(t *testing.T) {
	t.Parallel()
	expectedDomainName := "udtestdev-my-new-tls.wallet"
	domainName, err := uns.Unhash("0x1586d090e1b5781399f988e4b4f5639f4c2775ef5ec093d1279bb95b9bceb1a0")
	assert.Nil(t, err)
	assert.Equal(t, expectedDomainName, domainName)
}

func TestUnsL2UnhashDotWallet(t *testing.T) {
	t.Parallel()
	expectedDomainName := "udtestdev-test-l2-domain-784391.wallet"
	domainName, err := uns.Unhash("0x40920d1d24c83454d9d64e6666927f3abb97b3fd67c7e1bf43de5c2f4297f3b8")
	assert.Nil(t, err)
	assert.Equal(t, expectedDomainName, domainName)
}

func TestUnsNamehash(t *testing.T) {
	t.Parallel()
	domainName := "udtestdev-my-new-tls.wallet"
	expectedNamehash := "0x1586d090e1b5781399f988e4b4f5639f4c2775ef5ec093d1279bb95b9bceb1a0"
	namehash, err := uns.Namehash(domainName)
	assert.Nil(t, err)
	assert.Equal(t, expectedNamehash, namehash)

	domainName = "udtestdev-my-new-tls.test.wallet"
	expectedNamehash = "0x126aa8f6239a84fe8bcfd7129b96176fa7ddd8c652e1c5bb30fe50ec595fd7a1"
	namehash, err = uns.Namehash(domainName)
	assert.Nil(t, err)
	assert.Equal(t, expectedNamehash, namehash)

	domainName = "wallet"
	expectedNamehash = "0x1e3f482b3363eb4710dae2cb2183128e272eafbe137f686851c1caea32502230"
	namehash, err = uns.Namehash(domainName)
	assert.Nil(t, err)
	assert.Equal(t, expectedNamehash, namehash)
}

func TestUnsUnhashWithout0xPrefixDotWallet(t *testing.T) {
	t.Parallel()
	expectedDomainName := "udtestdev-my-new-tls.wallet"
	domainName, err := uns.Unhash("1586d090e1b5781399f988e4b4f5639f4c2775ef5ec093d1279bb95b9bceb1a0")
	assert.Nil(t, err)
	assert.Equal(t, expectedDomainName, domainName)
}

func TestUnsSingleL1Locations(t *testing.T) {
	t.Parallel()
	testDomainL1 := "test-usdt-and-dns-records.crypto"
	expectedLocations := map[string]namingservice.Location{}
	expectedLocations[testDomainL1] = namingservice.Location{
		RegistryAddress:       "0xAad76bea7CFEc82927239415BB18D2e93518ecBB",
		ResolverAddress:       "0x95AE1515367aa64C462c71e87157771165B1287A",
		NetworkId:             4,
		Blockchain:            "ETH",
		OwnerAddress:          "0x499dD6D875787869670900a2130223D85d4F6Aa7",
		BlockchainProviderUrl: "https://rinkeby.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e",
	}
	locations, err := uns.Locations([]string{testDomainL1})

	assert.Nil(t, err)
	assert.Equal(t, expectedLocations, locations)
}
func TestUnsSingleL2Locations(t *testing.T) {
	t.Parallel()
	testDomainL2 := "udtestdev-test-l2-domain-784391.wallet"
	expectedLocations := map[string]namingservice.Location{}
	expectedLocations[testDomainL2] = namingservice.Location{
		RegistryAddress:       "0x2a93C52E7B6E7054870758e15A1446E769EdfB93",
		ResolverAddress:       "0x2a93C52E7B6E7054870758e15A1446E769EdfB93",
		NetworkId:             80001,
		Blockchain:            "MATIC",
		OwnerAddress:          "0x499dD6D875787869670900a2130223D85d4F6Aa7",
		BlockchainProviderUrl: "https://polygon-mumbai.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e",
	}
	locations, err := uns.Locations([]string{testDomainL2})

	assert.Nil(t, err)
	assert.Equal(t, expectedLocations, locations)
}
func TestUnsLocationsDomainOnBothNetworks(t *testing.T) {
	t.Parallel()
	testDomainL1AndL2 := "udtestdev-test-l1-and-l2-ownership.wallet"
	testDomainL2 := "udtestdev-test-l2-domain-784391.wallet"
	expectedLocations := map[string]namingservice.Location{}
	expectedLocations[testDomainL1AndL2] = namingservice.Location{
		RegistryAddress:       "0x2a93C52E7B6E7054870758e15A1446E769EdfB93",
		ResolverAddress:       "0x2a93C52E7B6E7054870758e15A1446E769EdfB93",
		NetworkId:             80001,
		Blockchain:            "MATIC",
		OwnerAddress:          "0x499dD6D875787869670900a2130223D85d4F6Aa7",
		BlockchainProviderUrl: "https://polygon-mumbai.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e",
	}
	expectedLocations[testDomainL2] = namingservice.Location{
		RegistryAddress:       "0x2a93C52E7B6E7054870758e15A1446E769EdfB93",
		ResolverAddress:       "0x2a93C52E7B6E7054870758e15A1446E769EdfB93",
		NetworkId:             80001,
		Blockchain:            "MATIC",
		OwnerAddress:          "0x499dD6D875787869670900a2130223D85d4F6Aa7",
		BlockchainProviderUrl: "https://polygon-mumbai.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e",
	}
	locations, err := uns.Locations([]string{testDomainL1AndL2, testDomainL2})

	assert.Nil(t, err)
	assert.Equal(t, expectedLocations, locations)
}
func TestUnsLocations(t *testing.T) {
	t.Parallel()
	testDomainL1 := "test-usdt-and-dns-records.crypto"
	testDomainL2 := "udtestdev-test-l2-domain-784391.wallet"
	expectedLocations := map[string]namingservice.Location{}
	expectedLocations[testDomainL1] = namingservice.Location{
		RegistryAddress:       "0xAad76bea7CFEc82927239415BB18D2e93518ecBB",
		ResolverAddress:       "0x95AE1515367aa64C462c71e87157771165B1287A",
		NetworkId:             4,
		Blockchain:            "ETH",
		OwnerAddress:          "0x499dD6D875787869670900a2130223D85d4F6Aa7",
		BlockchainProviderUrl: "https://rinkeby.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e",
	}
	expectedLocations[testDomainL2] = namingservice.Location{
		RegistryAddress:       "0x2a93C52E7B6E7054870758e15A1446E769EdfB93",
		ResolverAddress:       "0x2a93C52E7B6E7054870758e15A1446E769EdfB93",
		NetworkId:             80001,
		Blockchain:            "MATIC",
		OwnerAddress:          "0x499dD6D875787869670900a2130223D85d4F6Aa7",
		BlockchainProviderUrl: "https://polygon-mumbai.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e",
	}
	locations, err := uns.Locations([]string{testDomainL1, testDomainL2})

	assert.Nil(t, err)
	assert.Equal(t, expectedLocations, locations)
}

func TestUnsLocationsMainnet(t *testing.T) {
	var uns, _ = NewUnsBuilder().SetEthereumNetwork("mainnet").SetL2EthereumNetwork("polygon").Build()

	t.Parallel()
	testDomainL1 := "ryan.crypto"
	testDomainL2 := "fwefwf.crypto"
	expectedLocations := map[string]namingservice.Location{}
	expectedLocations[testDomainL1] = namingservice.Location{
		RegistryAddress:       "0xD1E5b0FF1287aA9f9A268759062E4Ab08b9Dacbe",
		ResolverAddress:       "0xb66DcE2DA6afAAa98F2013446dBCB0f4B0ab2842",
		NetworkId:             1,
		Blockchain:            "ETH",
		OwnerAddress:          "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2",
		BlockchainProviderUrl: "https://mainnet.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e",
	}
	expectedLocations[testDomainL2] = namingservice.Location{
		RegistryAddress:       "0xa9a6A3626993D487d2Dbda3173cf58cA1a9D9e9f",
		ResolverAddress:       "0xa9a6A3626993D487d2Dbda3173cf58cA1a9D9e9f",
		NetworkId:             137,
		Blockchain:            "MATIC",
		OwnerAddress:          "0x58cA45E932a88b2E7D0130712B3AA9fB7c5781e2",
		BlockchainProviderUrl: "https://polygon-mainnet.infura.io/v3/c5da69dfac9c4d9d96dd232580d4124e",
	}
	locations, err := uns.Locations([]string{testDomainL1, testDomainL2})

	assert.Nil(t, err)
	assert.Equal(t, expectedLocations, locations)
}
func TestUnsLocationsNullValues(t *testing.T) {

	t.Parallel()
	testDomainL1 := "invaliddomain.crypto"
	testDomainL2 := "invaliddomain2.crypto"
	expectedLocations := map[string]namingservice.Location{}
	expectedLocations[testDomainL1] = namingservice.Location{
		RegistryAddress:       "",
		ResolverAddress:       "",
		NetworkId:             0,
		Blockchain:            "",
		OwnerAddress:          "",
		BlockchainProviderUrl: "",
	}
	expectedLocations[testDomainL2] = namingservice.Location{
		RegistryAddress:       "",
		ResolverAddress:       "",
		NetworkId:             0,
		Blockchain:            "",
		OwnerAddress:          "",
		BlockchainProviderUrl: "",
	}
	locations, err := uns.Locations([]string{testDomainL1, testDomainL2})

	assert.Nil(t, err)
	assert.Equal(t, expectedLocations, locations)
}
