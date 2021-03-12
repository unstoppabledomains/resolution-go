package dnsrecords

type Record struct {
	Type  Type
	TTL   uint32
	Value string
}

type (
	Type string
)

const DefaultTTL uint32 = 300
