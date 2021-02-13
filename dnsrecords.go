package resolution

import (
	"fmt"
	"github.com/DeRain/resolution-go/dnsrecords"
	"strings"
)

func DnsTypesToRecordKeys(types []dnsrecords.Type) ([]string, error) {
	recordKeys := []string{"dns.ttl"}
	for _, dnsType := range types {
		var key strings.Builder
		var ttlKey strings.Builder
		_, err := fmt.Fprintf(&key, "dns.%v", dnsType)
		if err != nil {
			return nil, err
		}
		_, err = fmt.Fprintf(&ttlKey, "dns.%v.ttl", dnsType)
		if err != nil {
			return nil, err
		}
		recordKeys = append(recordKeys, key.String(), ttlKey.String())
	}

	return recordKeys, nil
}
