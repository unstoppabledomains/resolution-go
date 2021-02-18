package resolution

import (
	"github.com/DeRain/resolution-go/dnsrecords"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDnsTypesToRecordKeys(t *testing.T) {
	t.Parallel()
	expectedRecords := []string{"dns.ttl", "dns.A", "dns.A.ttl", "dns.AAAA", "dns.AAAA.ttl"}
	records, err := DnsTypesToCryptoRecordKeys([]dnsrecords.Type{dnsrecords.A, dnsrecords.AAAA})
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, records)
}

func TestEmptyDnsTypesToRecordKeys(t *testing.T) {
	t.Parallel()
	expectedRecords := []string{"dns.ttl"}
	records, err := DnsTypesToCryptoRecordKeys([]dnsrecords.Type{})
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, records)
}

func TestCryptoRecordsToDns(t *testing.T) {
	t.Parallel()
	expectedRecords := []dnsrecords.Record{
		{Type: dnsrecords.A, TTL: 1800, Value: "10.0.0.1"},
		{Type: dnsrecords.A, TTL: 1800, Value: "10.0.0.2"},
		{Type: dnsrecords.AAAA, TTL: 1000, Value: "2400:cb00:2049:1::a29f:1804"},
		{Type: dnsrecords.AAAA, TTL: 1000, Value: "2001:db8::8a2e:370:7334"},
		{Type: dnsrecords.TXT, TTL: 666, Value: "unstoppable"},
		{Type: dnsrecords.TXT, TTL: 666, Value: "test"},
	}
	records, err := CryptoRecordsToDns(map[string]string{
		"dns.A":          "[\"10.0.0.1\",\"10.0.0.2\"]",
		"dns.AAAA":       "[\"2400:cb00:2049:1::a29f:1804\",\"2001:db8::8a2e:370:7334\"]",
		"dns.A.ttl":      "1800",
		"dns.ttl":        "1000",
		"invalid.record": "404",
		"dns.CNAME.ttl":  "100500",
		"dns.TXT":        "[\"unstoppable\",\"test\"]",
		"dns.TXT.ttl":    "666",
	})
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedRecords, records)
}

func TestCryptoRecordsToDnsInvalidKeys(t *testing.T) {
	t.Parallel()
	records, err := CryptoRecordsToDns(map[string]string{
		"crypto.BTC.address":         "bc1q359khn0phg58xgezyqsuuaha28zkwx047c0c3y",
		"crypto.ETH.address":         "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8",
		"gundb.public_key.value":     "pqeBHabDQdCHhbdivgNEc74QO-x8CPGXq4PKWgfIzhY.7WJR5cZFuSyh1bFwx0GWzjmrim0T5Y6Bp0SSK0im3nI",
		"gundb.username.value":       "0x8912623832e174f2eb1f59cc3b587444d619376ad5bf10070e937e0dc22b9ffb2e3ae059e6ebf729f87746b2f71e5d88ec99c1fb3c7c49b8617e2520d474c48e1c",
		"ipfs.html.value":            "Qme54oEzRkgooJbCDr78vzKAWcv6DDEZqRhhDyDtzgrZP6",
		"ipfs.redirect_domain.value": "https://abbfe6z95qov3d40hf6j30g7auo7afhp.mypinata.cloud/ipfs/Qme54oEzRkgooJbCDr78vzKAWcv6DDEZqRhhDyDtzgrZP6",
	})
	assert.Nil(t, err)
	assert.Len(t, records, 0)
}

func TestCryptoRecordsToDnsDefaultTTL(t *testing.T) {
	t.Parallel()
	expectedRecords := []dnsrecords.Record{
		{Type: dnsrecords.A, TTL: dnsrecords.DefaultTTL, Value: "10.0.0.1"},
		{Type: dnsrecords.AAAA, TTL: dnsrecords.DefaultTTL, Value: "2400:cb00:2049:1::a29f:1804"},
		{Type: dnsrecords.TXT, TTL: dnsrecords.DefaultTTL, Value: "test"},
	}
	records, err := CryptoRecordsToDns(map[string]string{
		"dns.A":    "[\"10.0.0.1\"]",
		"dns.AAAA": "[\"2400:cb00:2049:1::a29f:1804\"]",
		"dns.TXT":  "[\"test\"]",
	})
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedRecords, records)
}

func TestCryptoRecordsToDnsInvalidValue(t *testing.T) {
	t.Parallel()
	expectedRecords := []dnsrecords.Record{
		{Type: dnsrecords.AAAA, TTL: 1000, Value: "2400:cb00:2049:1::a29f:1804"},
		{Type: dnsrecords.TXT, TTL: 1000, Value: "test"},
	}
	records, err := CryptoRecordsToDns(map[string]string{
		"dns.A":    "invalid value",
		"dns.AAAA": "[\"2400:cb00:2049:1::a29f:1804\"]",
		"dns.TXT":  "[\"test\"]",
		"dns.ttl":  "1000",
	})
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedRecords, records)
}

func TestCryptoRecordsToDnsInvalidGlobalTTL(t *testing.T) {
	t.Parallel()
	expectedRecords := []dnsrecords.Record{
		{Type: dnsrecords.AAAA, TTL: dnsrecords.DefaultTTL, Value: "2400:cb00:2049:1::a29f:1804"},
		{Type: dnsrecords.TXT, TTL: dnsrecords.DefaultTTL, Value: "test"},
	}
	records, err := CryptoRecordsToDns(map[string]string{
		"dns.AAAA": "[\"2400:cb00:2049:1::a29f:1804\"]",
		"dns.TXT":  "[\"test\"]",
		"dns.ttl":  "invalid ttl",
	})
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedRecords, records)
}

func TestCryptoRecordsToDnsInvalidRecordTTL(t *testing.T) {
	t.Parallel()
	expectedRecords := []dnsrecords.Record{
		{Type: dnsrecords.AAAA, TTL: 500, Value: "2400:cb00:2049:1::a29f:1804"},
		{Type: dnsrecords.TXT, TTL: 500, Value: "test"},
	}
	records, err := CryptoRecordsToDns(map[string]string{
		"dns.AAAA":     "[\"2400:cb00:2049:1::a29f:1804\"]",
		"dns.AAAA.ttl": "invalid ttl",
		"dns.TXT":      "[\"test\"]",
		"dns.TXT.ttl":  "invalid ttl",
		"dns.ttl":      "500",
	})
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedRecords, records)
}

func TestCryptoRecordsToDnsEmpty(t *testing.T) {
	t.Parallel()
	records, err := CryptoRecordsToDns(map[string]string{})
	assert.Nil(t, err)
	assert.Empty(t, records)
}
