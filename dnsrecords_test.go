package resolution

import (
	"github.com/DeRain/resolution-go/dnsrecords"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDnsTypesToRecordKeys(t *testing.T) {
	t.Parallel()
	expectedRecords := []string{"dns.ttl", "dns.A", "dns.A.ttl", "dns.AAAA", "dns.AAAA.ttl"}
	records, err := DnsTypesToRecordKeys([]dnsrecords.Type{dnsrecords.A, dnsrecords.AAAA})
	assert.Nil(t, err)
	assert.Equal(t, expectedRecords, records)
}
