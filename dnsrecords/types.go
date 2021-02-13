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

const (
	A          Type = "A"
	AAAA            = "AAAA"
	AFSDB           = "AFSDB"
	APL             = "APL"
	CAA             = "CAA"
	CDNSKEY         = "CDNSKEY"
	CDS             = "CDS"
	CERT            = "CERT"
	CNAME           = "CNAME"
	CSYNC           = "CSYNC"
	DHCID           = "DHCID"
	DLV             = "DLV"
	DNAME           = "DNAME"
	DNSKEY          = "DNSKEY"
	DS              = "DS"
	EUI48           = "EUI48"
	EUI64           = "EUI64"
	HINFO           = "HINFO"
	HIP             = "HIP"
	HTTPS           = "HTTPS"
	IPSECKEY        = "IPSECKEY"
	KEY             = "KEY"
	KX              = "KX"
	LOC             = "LOC"
	MX              = "MX"
	NAPTR           = "NAPTR"
	NS              = "NS"
	NSEC            = "NSEC"
	NSEC3           = "NSEC3"
	NSEC3PARAM      = "NSEC3PARAM"
	OPENPGPKEY      = "OPENPGPKEY"
	PTR             = "PTR"
	RP              = "RP"
	RRSIG           = "RRSIG"
	SIG             = "SIG"
	SMIMEA          = "SMIMEA"
	SOA             = "SOA"
	SRV             = "SRV"
	SSHFP           = "SSHFP"
	SVCB            = "SVCB"
	TA              = "TA"
	TKEY            = "TKEY"
	TLSA            = "TLSA"
	TSIG            = "TSIG"
	TXT             = "TXT"
	URI             = "URI"
	ZONEMD          = "ZONEMD"
)
