package dnsrecords

// Record DNS record
type Record struct {
	// Type DNS record type
	Type Type
	// TTL DNS record TTL in seconds
	TTL uint32
	// Value DNS record value
	Value string
}

// Type DNS record type according to specification (A, AAAA, CNAME, etc.)
type (
	Type string
)

// DefaultTTL Default DNS TTL in seconds if domain does not have TTL records attached
const DefaultTTL uint32 = 300
