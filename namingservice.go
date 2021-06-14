package resolution

import (
	"github.com/unstoppabledomains/resolution-go/dnsrecords"
	"github.com/unstoppabledomains/resolution-go/namingservice"
	"strings"
)

// NamingService Unstoppable supports multiple naming services (.zil and .crypto).
// NamingService Each naming service implements shared interface and returns similar record types.
type NamingService interface {
	// Records Retrieve records of domain.
	// Keys must be provided in raw format according to specification.
	// Keys specification: https://docs.unstoppabledomains.com/domain-registry-essentials/records-reference.
	// Supported keys reference: https://github.com/unstoppabledomains/dot-crypto/blob/master/src/supported-keys/supported-keys.json.
	// It returns key-value map of specified keys set on provided domain. Map can contain empty strings if keys are not found.
	Records(domainName string, keys []string) (map[string]string, error)

	// Record Retrieve single record of domain.
	// Keys must be provided in raw format according to specification.
	// Keys specification: https://docs.unstoppabledomains.com/domain-registry-essentials/records-reference.
	// Supported keys reference: https://github.com/unstoppabledomains/dot-crypto/blob/master/src/supported-keys/supported-keys.json.
	// Returns result in string or empty string if no result found.
	Record(domainName string, key string) (string, error)

	// Addr Retrieve the value of domain's currency ticker.
	// Ticker must contain cryptocurrency like: BTC, ETH.
	// Returns result in string or empty string if no result found.
	Addr(domainName string, ticker string) (string, error)

	// AddrVersion Retrieve the version value of domain's currency ticker.
	// This method should be used to query multi-chain currency like USDT.
	// Returns result in string or empty string if no result found.
	AddrVersion(domainName string, ticker string, version string) (string, error)

	// Email Retrieve the email of domain.
	// Returns result in string or empty string if no result found.
	Email(domainName string) (string, error)

	// Resolver Retrieve the resolver address.
	// Returns result or DomainNotConfiguredError if resolver is not found.
	Resolver(domainName string) (string, error)

	// Owner Retrieve the owner address.
	// Returns result or DomainNotRegisteredError if owner is not found.
	Owner(domainName string) (string, error)

	// IpfsHash Retrieve hash of IPFS website attached to domain.
	IpfsHash(domainName string) (string, error)

	// HTTPUrl Retrieve the http redirect url of a domain.
	HTTPUrl(domainName string) (string, error)

	// AllRecords Retrieve all records of a domain.
	// Returns result in string or empty string record is not found.
	AllRecords(domainName string) (map[string]string, error)

	// DNS Retrieve the DNS records of a domain.
	// Returns a set of valid and filtered non-empty DNS records attached to domain.
	DNS(domainName string, types []dnsrecords.Type) ([]dnsrecords.Record, error)

	// IsSupportedDomain checks whether domain name is supported by the naming service.
	IsSupportedDomain(domainName string) bool

	// TokenURI returns ERC721 metadata token URI
	TokenURI(domainName string) (string, error)
}

var supportedNamingServices = map[string]string{
	"crypto": namingservice.CNS,
	"zil":    namingservice.ZNS,
}

// DetectNamingService helper to detect naming service type for provided domain.
// Returns ZNS or CNS for valid domain and error if domain is not valid or not supported by resolution-go library.
func DetectNamingService(domainName string) (string, error) {
	chunks := strings.Split(domainName, ".")
	if len(chunks) == 0 {
		return "", &DomainNotSupportedError{DomainName: domainName}
	}
	extension := chunks[len(chunks)-1]
	service := supportedNamingServices[extension]
	if service == "" {
		return "", &DomainNotSupportedError{DomainName: domainName}
	}
	return service, nil
}
