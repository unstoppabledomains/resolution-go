package resolution

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/unstoppabledomains/resolution-go/v3/dnsrecords"
)

func dnsTypesToCryptoRecordKeys(types []dnsrecords.Type) ([]string, error) {
	recordKeys := []string{"dns.ttl"}
	var key strings.Builder
	var ttlKey strings.Builder
	for _, dnsType := range types {
		key.Reset()
		ttlKey.Reset()
		dnsTypeUppercase := strings.ToUpper(string(dnsType))
		_, err := fmt.Fprintf(&key, "dns.%s", dnsTypeUppercase)
		if err != nil {
			return nil, err
		}
		_, err = fmt.Fprintf(&ttlKey, "dns.%s.ttl", dnsTypeUppercase)
		if err != nil {
			return nil, err
		}
		recordKeys = append(recordKeys, key.String(), ttlKey.String())
	}

	return recordKeys, nil
}

func cryptoRecordsToDNS(cryptoRecords map[string]string) ([]dnsrecords.Record, error) {
	var globalTTL = dnsrecords.DefaultTTL
	if cryptoRecords["dns.ttl"] != "" {
		parsedTTL, err := strconv.ParseUint(cryptoRecords["dns.ttl"], 10, 32)
		if err == nil {
			globalTTL = uint32(parsedTTL)
		}
	}
	var ttlKey strings.Builder
	var parsedDNSRecords []dnsrecords.Record
	for cryptoKey, cryptoValue := range cryptoRecords {
		if strings.Index(cryptoKey, "dns.") == 0 {
			keyParts := strings.Split(cryptoKey, ".")
			if len(keyParts) == 2 {
				recordTTL := globalTTL
				recordType := dnsrecords.Type(keyParts[1])
				ttlKey.Reset()
				_, err := fmt.Fprintf(&ttlKey, "dns.%v.ttl", recordType)
				if err != nil {
					return nil, err
				}
				if cryptoRecords[ttlKey.String()] != "" {
					parsedTTL, err := strconv.ParseUint(cryptoRecords[ttlKey.String()], 10, 32)
					if err == nil {
						recordTTL = uint32(parsedTTL)
					}
				}
				var parsedDNSRecordValues []string
				err = json.Unmarshal([]byte(cryptoValue), &parsedDNSRecordValues)
				if err == nil {
					for _, dnsRecordValue := range parsedDNSRecordValues {
						parsedDNSRecords = append(parsedDNSRecords, dnsrecords.Record{Type: recordType, TTL: recordTTL, Value: dnsRecordValue})
					}
				}
			}
		}
	}

	return parsedDNSRecords, nil
}
