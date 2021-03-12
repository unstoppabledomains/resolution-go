package resolution

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/unstoppabledomains/resolution-go/dnsrecords"
)

var cns, _ = NewCnsBuilder().Build()

func TestCnsBuilder(t *testing.T) {
	t.Parallel()
	builder := NewCnsBuilder()
	_, err := builder.Build()
	assert.Nil(t, err)
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

func TestNewCnsWithSupportedKeys(t *testing.T) {
	t.Parallel()
	cnsService, _ := NewCnsBuilder().Build()
	deprecatedKeyName := cnsService.supportedKeys["crypto.ETH.address"]
	assert.Equal(t, "ETH", deprecatedKeyName.DeprecatedKeyName)
}

func TestCnsDataValue(t *testing.T) {
	t.Parallel()
	testDomain := "brad.crypto"
	expectedRecord := "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
	data, err := cns.Data(testDomain, []string{"crypto.ETH.address"})
	assert.Nil(t, err)
	assert.Equal(t, data.Values[0], expectedRecord)
}

func TestCnsData(t *testing.T) {
	t.Parallel()
	testDomain := "brad.crypto"
	expectedRecord := "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
	expectedOwner := common.HexToAddress("0x8aaD44321A86b170879d7A244c1e8d360c99DdA8")
	expectedResolver := common.HexToAddress("0xb66DcE2DA6afAAa98F2013446dBCB0f4B0ab2842")
	data, err := cns.Data(testDomain, []string{"crypto.ETH.address"})
	assert.Nil(t, err)
	assert.Equal(t, data.Values[0], expectedRecord)
	assert.Equal(t, expectedOwner, data.Owner)
	assert.Equal(t, expectedResolver, data.Resolver)
}

func TestCnsEmptyDataValues(t *testing.T) {
	t.Parallel()
	testDomain := "brad.crypto"
	data, _ := cns.Data(testDomain, []string{"empty record"})
	assert.Equal(t, data.Values[0], "")
	assert.Len(t, data.Values, 1)
}

func TestCnsDomainNotRegistered(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegistered
	testDomain := "not-registered-long-domain-name.crypto"
	_, err := cns.Data(testDomain, []string{"crypto.ETH.address"})
	assert.ErrorAs(t, err, &expectedError)
}

func TestCnsDomainNotConfigured(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotConfigured
	testDomain := "reseller-test-paul2.crypto"
	_, err := cns.Data(testDomain, []string{"crypto.ETH.address"})
	assert.ErrorAs(t, err, &expectedError)
}

func TestCnsRecords(t *testing.T) {
	t.Parallel()
	testDomain := "brad.crypto"
	expectedRecords := map[string]string{"crypto.ETH.address": "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8", "crypto.BTC.address": "bc1q359khn0phg58xgezyqsuuaha28zkwx047c0c3y"}
	records, err := cns.Records(testDomain, []string{"crypto.ETH.address", "crypto.BTC.address"})
	assert.Nil(t, err)
	assert.Equal(t, records, expectedRecords)
}

func TestCnsNoRecords(t *testing.T) {
	t.Parallel()
	testDomain := "brad.crypto"
	records, err := cns.Records(testDomain, []string{})
	assert.Nil(t, err)
	assert.Empty(t, records)
}

func TestCnsEmptyRecords(t *testing.T) {
	t.Parallel()
	testDomain := "brad.crypto"
	expectedRecords := map[string]string{"record-not-exist": "", "crypto.ETH.address": "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8", "crypto.BTC.address": "bc1q359khn0phg58xgezyqsuuaha28zkwx047c0c3y"}
	records, err := cns.Records(testDomain, []string{"record-not-exist", "crypto.ETH.address", "crypto.BTC.address"})
	assert.Nil(t, err)
	assert.Equal(t, records, expectedRecords)
}

func TestCnsRecord(t *testing.T) {
	t.Parallel()
	testDomain := "brad.crypto"
	expectedRecord := "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
	record, err := cns.Record(testDomain, "crypto.ETH.address")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestCnsEmptyRecord(t *testing.T) {
	t.Parallel()
	testDomain := "brad.crypto"
	record, err := cns.Record(testDomain, "record-not-exist")
	assert.Nil(t, err)
	assert.Empty(t, record)
}

func TestCnsAddr(t *testing.T) {
	t.Parallel()
	testDomain := "brad.crypto"
	expectedRecord := "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
	record, err := cns.Addr(testDomain, "ETH")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestCnsAddrLowerCaseTicker(t *testing.T) {
	t.Parallel()
	testDomain := "brad.crypto"
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
	testDomain := "brad.crypto"
	expectedRecords := map[string]string{
		"crypto.BTC.address":         "bc1q359khn0phg58xgezyqsuuaha28zkwx047c0c3y",
		"crypto.ETH.address":         "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8",
		"gundb.public_key.value":     "pqeBHabDQdCHhbdivgNEc74QO-x8CPGXq4PKWgfIzhY.7WJR5cZFuSyh1bFwx0GWzjmrim0T5Y6Bp0SSK0im3nI",
		"gundb.username.value":       "0x8912623832e174f2eb1f59cc3b587444d619376ad5bf10070e937e0dc22b9ffb2e3ae059e6ebf729f87746b2f71e5d88ec99c1fb3c7c49b8617e2520d474c48e1c",
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
	assert.True(t, cns.IsSupportedDomain("valid.crypto"))
	assert.False(t, cns.IsSupportedDomain("invalid.zil"))
	assert.False(t, cns.IsSupportedDomain("invalid.com"))
}

func TestCnsUnsupportedDomainError(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotSupported
	_, err := cns.Data("invalid.zil", []string{"crypto.ETH.address"})
	assert.ErrorAs(t, err, &expectedError)
}
