package resolution

import (
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/stretchr/testify/assert"
	"testing"
)

var zns = NewZnsWithDefaultProvider()

func TestNewZns(t *testing.T) {
	t.Parallel()
	znsProvider := provider.NewProvider("https://api.zilliqa.com")
	zns := NewZns(znsProvider)
	assert.IsType(t, &Zns{Provider: nil}, zns)
}

func TestNewZnsWithDefaultProvider(t *testing.T) {
	t.Parallel()
	zns := NewZnsWithDefaultProvider()
	assert.IsType(t, &Zns{Provider: nil}, zns)
}

func TestZnsStateAllRecords(t *testing.T) {
	t.Parallel()
	expectedRecords := map[string]string{
		"ipfs.html.value":            "QmVaAtQbi3EtsfpKoLzALm6vXphdi2KjMgxEDKeGg6wHuK",
		"crypto.BCH.address":         "qrq4sk49ayvepqz7j7ep8x4km2qp8lauvcnzhveyu6",
		"crypto.BTC.address":         "1EVt92qQnaLDcmVFtHivRJaunG2mf2C3mB",
		"crypto.ETH.address":         "0x45b31e01AA6f42F0549aD482BE81635ED3149abb",
		"crypto.LTC.address":         "LetmswTW3b7dgJ46mXuiXMUY17XbK29UmL",
		"crypto.XMR.address":         "447d7TVFkoQ57k3jm3wGKoEAkfEym59mK96Xw5yWamDNFGaLKW5wL2qK5RMTDKGSvYfQYVN7dLSrLdkwtKH3hwbSCQCu26d",
		"crypto.ZEC.address":         "t1h7ttmQvWCSH1wfrcmvT4mZJfGw2DgCSqV",
		"crypto.ZIL.address":         "zil1yu5u4hegy9v3xgluweg4en54zm8f8auwxu0xxj",
		"crypto.DASH.address":        "XnixreEBqFuSLnDSLNbfqMH1GsZk7cgW4j",
		"ipfs.redirect_domain.value": "www.unstoppabledomains.com",
	}
	expectedOwner := "0x2d418942dce1afa02d0733a2000c71b371a6ac07"
	expectedResolver := "0xdac22230adfe4601f00631eae92df6d77f054891"
	state, err := zns.State("brad.zil")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, state.Records)
	assert.Equal(t, expectedOwner, state.Owner)
	assert.Equal(t, expectedResolver, state.Resolver)
}

func TestZnsStateDomainNotRegistered(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotRegistered
	_, err := zns.State("long-not-registered-name.zil")
	assert.ErrorAs(t, err, &expectedError)
}

func TestZnsStateDomainNotConfigured(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotConfigured
	_, err := zns.State("1010.zil")
	assert.ErrorAs(t, err, &expectedError)
}

func TestZnsRecords(t *testing.T) {
	t.Parallel()
	expectedRecords := map[string]string{
		"ipfs.html.value":    "QmVaAtQbi3EtsfpKoLzALm6vXphdi2KjMgxEDKeGg6wHuK",
		"crypto.BCH.address": "qrq4sk49ayvepqz7j7ep8x4km2qp8lauvcnzhveyu6",
	}
	records, err := zns.Records("brad.zil", []string{"ipfs.html.value", "crypto.BCH.address"})
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, records)
}

func TestZnsEmptyRecords(t *testing.T) {
	t.Parallel()
	expectedRecords := map[string]string{
		"ipfs.html.value":    "QmVaAtQbi3EtsfpKoLzALm6vXphdi2KjMgxEDKeGg6wHuK",
		"crypto.BCH.address": "qrq4sk49ayvepqz7j7ep8x4km2qp8lauvcnzhveyu6",
		"key-not-exist":      "",
	}
	records, err := zns.Records("brad.zil", []string{"ipfs.html.value", "crypto.BCH.address", "key-not-exist"})
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, records)
}

func TestZnsRecord(t *testing.T) {
	t.Parallel()
	expectedRecord := "0x45b31e01AA6f42F0549aD482BE81635ED3149abb"
	record, err := zns.Record("brad.zil", "crypto.ETH.address")
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, record)
}

func TestZnsEmptyRecord(t *testing.T) {
	t.Parallel()
	record, err := zns.Record("brad.zil", "non-existent-key")
	assert.Nil(t, err)
	assert.Empty(t, record)
}

func TestZnsOwner(t *testing.T) {
	t.Parallel()
	expectedOwner := "0x2d418942dce1afa02d0733a2000c71b371a6ac07"
	owner, err := zns.Owner("brad.zil")
	assert.Nil(t, err)
	assert.Equal(t, expectedOwner, owner)
}

func TestZnsResolver(t *testing.T) {
	t.Parallel()
	expectedResolver := "0xdac22230adfe4601f00631eae92df6d77f054891"
	resolver, err := zns.Resolver("brad.zil")
	assert.Nil(t, err)
	assert.Equal(t, expectedResolver, resolver)
}
