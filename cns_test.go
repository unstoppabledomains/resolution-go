package resolution

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"testing"
)

var cns, _ = NewCnsWithDefaultBackend()

func TestNewCnsWithDefaultProvider(t *testing.T) {
	t.Parallel()
	_, err := NewCnsWithDefaultBackend()
	assert.Nil(t, err)
}

func TestNewCns(t *testing.T) {
	t.Parallel()
	backend, _ := ethclient.Dial(DefaultProvider)
	_, err := NewCns(backend)
	assert.Nil(t, err)
}

func TestCnsData(t *testing.T) {
	t.Parallel()
	testDomain := "brad.crypto"
	expectedRecord := "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
	data, err := cns.Data(testDomain, []string{"crypto.ETH.address"})
	assert.Nil(t, err)
	assert.Equal(t, data.Values[0], expectedRecord)
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
	expectedRecords := []string{"0x8aaD44321A86b170879d7A244c1e8d360c99DdA8", "bc1q359khn0phg58xgezyqsuuaha28zkwx047c0c3y"}
	records, err := cns.Records(testDomain, []string{"crypto.ETH.address", "crypto.BTC.address"})
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, err)
	assert.ObjectsAreEqual(records, expectedRecords)
}

func TestCnsRecord(t *testing.T) {
	t.Parallel()
	testDomain := "brad.crypto"
	expectedRecord := "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
	record, err := cns.Record(testDomain, "crypto.ETH.address")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}
