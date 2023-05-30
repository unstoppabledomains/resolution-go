package resolution

import (
	"net/http"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/unstoppabledomains/resolution-go/v3/dnsrecords"
	"github.com/unstoppabledomains/resolution-go/v3/namingservice"
)

type DomainData struct {
	Name string
}

var domains = map[string]DomainData{
	"DomainL1":           {Name: "reseller-test-udtesting-459239285.crypto"},
	"DomainL2":           {Name: "udtestdev-test-l2-domain-784391.wallet"},
	"DomainWallet":       {Name: "uns-devtest-265f8f.wallet"},
	"DomainNotRegistred": {Name: "not-registered-long-domain-name.wallet"},
}

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

func getUns() *Uns {
	builder := NewUnsBuilder()
	builder = builder.SetEthereumNetwork("goerli").SetContractBackendProviderUrl(getL1TestProviderUrl())
	builder = builder.SetL2EthereumNetwork("mumbai").SetL2ContractBackendProviderUrl(getL2TestProviderUrl())
	uns, _ := builder.Build()
	return uns
}

// TestNewUnsWithSupportedKeys uses default provider
func TestNewUnsWithSupportedKeys(t *testing.T) {
	t.Parallel()

	unsService := getUns()
	deprecatedKeyName := unsService.l1Service.supportedKeys["crypto.ETH.address"]
	assert.Equal(t, "ETH", deprecatedKeyName.DeprecatedKeyName)
}

func TestUnsL1DataValue(t *testing.T) {
	t.Parallel()
	expectedRecord := "0x084Ac37CDEfE1d3b68a63c08B203EFc3ccAB9742"

	uns := getUns()
	data, err := uns.Data(domains["DomainL1"].Name, []string{"crypto.ETH.address"})
	assert.Nil(t, err)
	assert.Equal(t, data.Values[0], expectedRecord)
}

func TestUnsL2DataValue(t *testing.T) {
	t.Parallel()
	expectedRecord := "0x6A1fd9a073256f14659fe59613bbf169Ed27CdcC"

	uns := getUns()
	data, err := uns.Data(domains["DomainL2"].Name, []string{"crypto.LINK.address"})
	assert.Nil(t, err)
	assert.Equal(t, data.Values[0], expectedRecord)
}

func TestUnsL1Data(t *testing.T) {
	t.Parallel()
	expectedRecord := "0x084Ac37CDEfE1d3b68a63c08B203EFc3ccAB9742"
	expectedOwner := common.HexToAddress("0xe586d5Bf4d7779498648DF67b73c88a712E4359d")
	expectedResolver := common.HexToAddress("0x0555344A5f440bd1d8CB6b42Db46C5E5d4070437")

	uns := getUns()
	data, err := uns.Data(domains["DomainL1"].Name, []string{"crypto.ETH.address"})
	assert.Nil(t, err)
	assert.Equal(t, data.Values[0], expectedRecord)
	assert.Equal(t, expectedOwner, data.Owner)
	assert.Equal(t, expectedResolver, data.Resolver)
}

func TestUnsL2Data(t *testing.T) {
	t.Parallel()
	expectedRecord := "0x6A1fd9a073256f14659fe59613bbf169Ed27CdcC"
	expectedOwner := common.HexToAddress("0x499dd6d875787869670900a2130223d85d4f6aa7")
	expectedResolver := common.HexToAddress("0x2a93C52E7B6E7054870758e15A1446E769EdfB93")

	uns := getUns()
	data, err := uns.Data(domains["DomainL2"].Name, []string{"crypto.LINK.address"})
	assert.Nil(t, err)
	assert.Equal(t, data.Values[0], expectedRecord)
	assert.Equal(t, expectedOwner, data.Owner)
	assert.Equal(t, expectedResolver, data.Resolver)
}

func TestUnsL1EmptyDataValues(t *testing.T) {
	t.Parallel()

	uns := getUns()
	data, _ := uns.Data(domains["DomainL1"].Name, []string{"empty record"})
	assert.Equal(t, data.Values[0], "")
	assert.Len(t, data.Values, 1)
}

func TestUnsL2EmptyDataValues(t *testing.T) {
	t.Parallel()

	uns := getUns()
	data, _ := uns.Data(domains["DomainL2"].Name, []string{"empty record"})
	assert.Equal(t, data.Values[0], "")
	assert.Len(t, data.Values, 1)
}

func TestUnsDomainNotRegistered(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegisteredError

	uns := getUns()
	_, err := uns.Data(domains["DomainNotRegistred"].Name, []string{"crypto.ETH.address"})
	assert.ErrorAs(t, err, &expectedError)
}

func TestUnsL1Records(t *testing.T) {
	t.Parallel()
	expectedRecords := map[string]string{"crypto.ETH.address": "0x084Ac37CDEfE1d3b68a63c08B203EFc3ccAB9742", "crypto.BTC.address": ""}

	uns := getUns()
	records, err := uns.Records(domains["DomainL1"].Name, []string{"crypto.ETH.address", "crypto.BTC.address"})
	assert.Nil(t, err)
	assert.Equal(t, records, expectedRecords)
}

func TestUnsL2Records(t *testing.T) {
	t.Parallel()
	expectedRecords := map[string]string{"crypto.LINK.address": "0x6A1fd9a073256f14659fe59613bbf169Ed27CdcC", "crypto.BTC.address": ""}

	uns := getUns()
	records, err := uns.Records(domains["DomainL2"].Name, []string{"crypto.LINK.address", "crypto.BTC.address"})
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, records)
}

func TestUnsEmptyRecords(t *testing.T) {
	t.Parallel()
	expectedRecords := map[string]string{"crypto.BTC.address": "", "crypto.ETH.address": "0x084Ac37CDEfE1d3b68a63c08B203EFc3ccAB9742", "record-not-exist": ""}

	uns := getUns()
	records, err := uns.Records(domains["DomainL1"].Name, []string{"record-not-exist", "crypto.ETH.address", "crypto.BTC.address"})
	assert.Nil(t, err)
	assert.Equal(t, records, expectedRecords)
}

func TestUnsL1Record(t *testing.T) {
	t.Parallel()
	expectedRecord := "0x084Ac37CDEfE1d3b68a63c08B203EFc3ccAB9742"

	uns := getUns()
	record, err := uns.Record(domains["DomainL1"].Name, "crypto.ETH.address")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL2Record(t *testing.T) {
	t.Parallel()
	expectedRecord := "0x6A1fd9a073256f14659fe59613bbf169Ed27CdcC"

	uns := getUns()
	record, err := uns.Record(domains["DomainL2"].Name, "crypto.LINK.address")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsEmptyRecord(t *testing.T) {
	t.Parallel()

	uns := getUns()
	record, err := uns.Record(domains["DomainL1"].Name, "record-not-exist")
	assert.Nil(t, err)
	assert.Empty(t, record)
}

func TestUnsL1Addr(t *testing.T) {
	t.Parallel()
	expectedRecord := "0x084Ac37CDEfE1d3b68a63c08B203EFc3ccAB9742"

	uns := getUns()
	record, err := uns.Addr(domains["DomainL1"].Name, "ETH")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL2Addr(t *testing.T) {
	t.Parallel()
	expectedRecord := "0x6A1fd9a073256f14659fe59613bbf169Ed27CdcC"

	uns := getUns()
	record, err := uns.Addr(domains["DomainL2"].Name, "LINK")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsZilOnL1Owner(t *testing.T) {
	t.Parallel()
	testDomain := "uns-devtest-testdomain303030.zil"
	expectedRecord := "0x499dD6D875787869670900a2130223D85d4F6Aa7"

	uns := getUns()
	record, err := uns.Owner(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsZilOnZilliqaOwner(t *testing.T) {
	t.Parallel()
	testDomain := "brad.zil"
	expectedRecord := "0x2d418942dce1afa02d0733a2000c71b371a6ac07"

	uns := getUns()
	record, err := uns.Owner(testDomain)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsAddrLowerCaseTicker(t *testing.T) {
	t.Parallel()
	expectedRecord := "0x084Ac37CDEfE1d3b68a63c08B203EFc3ccAB9742"

	uns := getUns()
	record, err := uns.Addr(domains["DomainL1"].Name, "eth")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL1Email(t *testing.T) {
	t.Parallel()
	expectedRecord := "testing@example.com"

	uns := getUns()
	record, err := uns.Email(domains["DomainWallet"].Name)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL2Email(t *testing.T) {
	t.Parallel()
	expectedRecord := "l2email@l2mail.mail"

	uns := getUns()
	record, err := uns.Email(domains["DomainL2"].Name)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL1Resolver(t *testing.T) {
	t.Parallel()
	expectedRecord := "0x0555344A5F440Bd1d8cb6B42db46c5e5D4070437"

	uns := getUns()
	record, err := uns.Resolver(domains["DomainL1"].Name)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL2Resolver(t *testing.T) {
	t.Parallel()
	expectedRecord := "0x2a93C52E7B6E7054870758e15A1446E769EdfB93"

	uns := getUns()
	record, err := uns.Resolver(domains["DomainL2"].Name)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL1Owner(t *testing.T) {
	t.Parallel()
	expectedRecord := "0xe586d5Bf4d7779498648DF67b73c88a712E4359d"

	uns := getUns()
	record, err := uns.Owner(domains["DomainL1"].Name)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL2Owner(t *testing.T) {
	t.Parallel()
	expectedRecord := "0x499dD6D875787869670900a2130223D85d4F6Aa7"

	uns := getUns()
	record, err := uns.Owner(domains["DomainL2"].Name)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsOwnerWithoutOwner(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegisteredError

	uns := getUns()
	_, err := uns.Owner(domains["DomainNotRegistred"].Name)
	assert.ErrorAs(t, err, &expectedError)
}

func TestUnsL1AddrVersion(t *testing.T) {
	t.Parallel()
	expectedRecord := "0xe7474D07fD2FA286e7e0aa23cd107F8379085037"

	uns := getUns()
	record, err := uns.AddrVersion(domains["DomainWallet"].Name, "USDT", "ERC20")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL1Ipfs(t *testing.T) {
	t.Parallel()
	expectedRecord := "QmdyBw5oTgCtTLQ18PbDvPL8iaLoEPhSyzD91q9XmgmAjb"

	uns := getUns()
	record, err := uns.IpfsHash(domains["DomainWallet"].Name)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL2Ipfs(t *testing.T) {
	t.Parallel()
	expectedRecord := "QmfRXG3CcM1eWiCUA89uzimCvQUnw4HzTKLo6hRZ47PYsN"

	uns := getUns()
	record, err := uns.IpfsHash(domains["DomainL2"].Name)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL1HTTPUrl(t *testing.T) {
	t.Parallel()
	expectedRecord := ""

	uns := getUns()
	record, err := uns.HTTPUrl(domains["DomainL1"].Name)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsL2HTTPUrl(t *testing.T) {
	t.Parallel()
	expectedRecord := ""

	uns := getUns()
	record, err := uns.HTTPUrl(domains["DomainL2"].Name)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestUnsDnsA(t *testing.T) {
	t.Parallel()
	expectedRecords := []dnsrecords.Record{
		{Type: "A", TTL: 98, Value: "10.0.0.1"},
		{Type: "A", TTL: 98, Value: "10.0.0.3"},
	}

	uns := getUns()
	dnsRecords, err := uns.DNS(domains["DomainWallet"].Name, []dnsrecords.Type{"A"})
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedRecords, dnsRecords)
}

func TestUnsDnsCname(t *testing.T) {
	t.Parallel()

	uns := getUns()
	expectedRecords := []dnsrecords.Record{}
	dnsRecords, err := uns.DNS(domains["DomainWallet"].Name, []dnsrecords.Type{"CNAME"})
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedRecords, dnsRecords)
}

func TestUnsIsSupportedDomain(t *testing.T) {
	t.Parallel()

	uns := getUns()
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

	uns := getUns()
	_, err := uns.Data("invalid.zil", []string{"crypto.ETH.address"})
	assert.ErrorAs(t, err, &expectedError)
}

func TestUnsTokenURI(t *testing.T) {
	t.Parallel()
	uns := getUns()
	tokenURI, err := uns.TokenURI(domains["DomainWallet"].Name)
	expectedTokenURI := "https://api.ud-staging.com/metadata/6304531997610998161237844647282663196661123000121147597890468333969432655810"
	assert.Nil(t, err)
	assert.Equal(t, expectedTokenURI, tokenURI)
}

func TestUnsTokenURIMetadata(t *testing.T) {
	t.Parallel()
	expectedMetadata := TokenMetadata{
		Name:        "uns-devtest-265f8f.wallet",
		Description: "A CNS or UNS blockchain domain. Use it to resolve your cryptocurrency addresses and decentralized websites.\nhttps://gateway.pinata.cloud/ipfs/QmdyBw5oTgCtTLQ18PbDvPL8iaLoEPhSyzD91q9XmgmAjb",
		ExternalUrl: "https://unstoppabledomains.com/search?searchTerm=uns-devtest-265f8f.wallet",
		Image:       "https://api.ud-staging.com/metadata/image-src/uns-devtest-265f8f.wallet.svg",
		Attributes: []TokenMetadataAttribute{
			{
				TraitType: "domain",
				Value:     "uns-devtest-265f8f.wallet",
			},
		},
	}

	uns := getUns()
	metadata, err := uns.TokenURIMetadata(domains["DomainWallet"].Name)
	assert.Nil(t, err)
	assert.Equal(t, expectedMetadata.Name, metadata.Name)
	assert.Equal(t, expectedMetadata.Description, metadata.Description)
	assert.Equal(t, expectedMetadata.ExternalUrl, metadata.ExternalUrl)
	assert.Equal(t, expectedMetadata.Image, metadata.Image)
}

func TestUnsUnhashDotCrypto(t *testing.T) {
	t.Parallel()
	expectedDomainName := "udtestdev-johnnytest.wallet"

	uns := getUns()
	domainName, err := uns.Unhash("0x684c51201935fdd42fbaebe43b1986f13984b94569c4c4827beda913232d066f")
	assert.Nil(t, err)
	assert.Equal(t, expectedDomainName, domainName)
}

func TestUnsUnhashWithout0xPrefixDotCrypto(t *testing.T) {
	t.Parallel()
	expectedDomainName := "udtestdev-johnnytest.wallet"

	uns := getUns()
	domainName, err := uns.Unhash("684c51201935fdd42fbaebe43b1986f13984b94569c4c4827beda913232d066f")
	assert.Nil(t, err)
	assert.Equal(t, expectedDomainName, domainName)
}

func TestUnsUnhashDotWallet(t *testing.T) {
	t.Parallel()
	expectedDomainName := "uns-devtest-265f8f.wallet"

	uns := getUns()
	domainName, err := uns.Unhash("0x0df03d18a0a02673661da22d06f43801a986840e5812989139f0f7a2c41037c2")
	assert.Nil(t, err)
	assert.Equal(t, expectedDomainName, domainName)
}

func TestUnsL2UnhashDotWallet(t *testing.T) {
	t.Parallel()
	expectedDomainName := "udtestdev-test-l2-domain-784391.wallet"

	uns := getUns()
	domainName, err := uns.Unhash("0x40920d1d24c83454d9d64e6666927f3abb97b3fd67c7e1bf43de5c2f4297f3b8")
	assert.Nil(t, err)
	assert.Equal(t, expectedDomainName, domainName)
}

func TestUnsNamehash(t *testing.T) {
	t.Parallel()
	expectedNamehash := "0x4fe5c8229795fec5cab66bf7e2c301f2f54cada203afb9b7b8b1d01213ede26d"

	uns := getUns()
	namehash, err := uns.Namehash(domains["DomainL1"].Name)
	assert.Nil(t, err)
	assert.Equal(t, expectedNamehash, namehash)

	expectedNamehash = "0x0df03d18a0a02673661da22d06f43801a986840e5812989139f0f7a2c41037c2"
	namehash, err = uns.Namehash(domains["DomainWallet"].Name)
	assert.Nil(t, err)
	assert.Equal(t, expectedNamehash, namehash)

	expectedNamehash = "0x1e3f482b3363eb4710dae2cb2183128e272eafbe137f686851c1caea32502230"
	namehash, err = uns.Namehash("wallet")
	assert.Nil(t, err)
	assert.Equal(t, expectedNamehash, namehash)
}

func TestUnsUnhashWithout0xPrefixDotWallet(t *testing.T) {
	t.Parallel()
	expectedDomainName := domains["DomainWallet"].Name

	uns := getUns()
	domainName, err := uns.Unhash("0df03d18a0a02673661da22d06f43801a986840e5812989139f0f7a2c41037c2")
	assert.Nil(t, err)
	assert.Equal(t, expectedDomainName, domainName)
}

func TestUnsSingleL1Locations(t *testing.T) {
	t.Parallel()
	expectedLocations := map[string]namingservice.Location{}
	expectedLocations[domains["DomainL1"].Name] = namingservice.Location{
		RegistryAddress:       "0x801452cFAC27e79a11c6b185986fdE09e8637589",
		ResolverAddress:       "0x0555344A5F440Bd1d8cb6B42db46c5e5D4070437",
		NetworkId:             5,
		Blockchain:            "ETH",
		OwnerAddress:          "0xe586d5Bf4d7779498648DF67b73c88a712E4359d",
		BlockchainProviderUrl: getL1TestProviderUrl(),
	}

	uns := getUns()
	locations, err := uns.Locations([]string{domains["DomainL1"].Name})

	assert.Nil(t, err)
	assert.Equal(t, expectedLocations, locations)
}

func TestUnsSingleL2Locations(t *testing.T) {
	t.Parallel()
	expectedLocations := map[string]namingservice.Location{}
	expectedLocations[domains["DomainL2"].Name] = namingservice.Location{
		RegistryAddress:       "0x2a93C52E7B6E7054870758e15A1446E769EdfB93",
		ResolverAddress:       "0x2a93C52E7B6E7054870758e15A1446E769EdfB93",
		NetworkId:             80001,
		Blockchain:            "MATIC",
		OwnerAddress:          "0x499dD6D875787869670900a2130223D85d4F6Aa7",
		BlockchainProviderUrl: getL2TestProviderUrl(),
	}

	uns := getUns()
	locations, err := uns.Locations([]string{domains["DomainL2"].Name})

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

	uns := getUns()
	locations, err := uns.Locations([]string{testDomainL1, testDomainL2})

	assert.Nil(t, err)
	assert.Equal(t, expectedLocations, locations)
}

func TestUnsLocationsNullValueForUnsupportedTLD(t *testing.T) {
	t.Parallel()
	testDomainL1 := "invaliddomain.eth"
	testDomainL2 := "invaliddomain2.dwe"
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

	uns := getUns()
	locations, err := uns.Locations([]string{testDomainL1, testDomainL2})

	assert.Nil(t, err)
	assert.Equal(t, expectedLocations, locations)
}

func TestUnsLocationsNoResolver(t *testing.T) {
	t.Parallel()
	testDomainL1 := "udtestdev-test-l2-domain-784391.wallet"
	expectedLocations := map[string]namingservice.Location{}
	expectedLocations[testDomainL1] = namingservice.Location{
		RegistryAddress:       "0x2a93C52E7B6E7054870758e15A1446E769EdfB93",
		ResolverAddress:       "0x2a93C52E7B6E7054870758e15A1446E769EdfB93",
		NetworkId:             80001,
		Blockchain:            "MATIC",
		OwnerAddress:          "0x499dD6D875787869670900a2130223D85d4F6Aa7",
		BlockchainProviderUrl: getL2TestProviderUrl(),
	}

	uns := getUns()
	locations, err := uns.Locations([]string{testDomainL1})

	assert.Nil(t, err)
	assert.Equal(t, expectedLocations, locations)
}
