package resolution

import (
	"github.com/unstoppabledomains/resolution-go/dnsrecords"
	"github.com/unstoppabledomains/resolution-go/namingservice"
	"strings"
)

type NamingService interface {
	// Records retrieve records of domain
	Records(domainName string, keys []string) (map[string]string, error)

	// Record Retrieve single record of domain
	Record(domainName string, key string) (string, error)

	// Addr Retrieve the value of domain's currency ticker
	Addr(domainName string, ticker string) (string, error)

	// AddrVersion Retrieve the version value of domain's currency ticker - useful for multichain currencies
	AddrVersion(domainName string, ticker string, version string) (string, error)

	// Email Retrieve the email of domain
	Email(domainName string) (string, error)

	// Resolver Retrieve the resolver set for a domain
	Resolver(domainName string) (string, error)

	// Owner Retrieve the owner of a domain
	Owner(domainName string) (string, error)

	// IpfsHash Retrieve the ipfs hash of a domain
	IpfsHash(domainName string) (string, error)

	// HTTPUrl Retrieve the http redirect url of a domain
	HTTPUrl(domainName string) (string, error)

	// AllRecords Retrieve all records of a domain
	AllRecords(domainName string) (map[string]string, error)

	// DNS Retrieve the DNS records of a domain
	DNS(domainName string, types []dnsrecords.Type) ([]dnsrecords.Record, error)

	// IsSupportedDomain checks whether domain name is supported by the naming service
	IsSupportedDomain(domainName string) bool
}

var supportedNamingServices = map[string]string{
	"crypto": namingservice.CNS,
	"zil":    namingservice.ZNS,
}

// DetectNamingServiceType helper to detect naming service type for provided domain
// Returns ZNS or CNS for valid domain and error if domain is not valid or not supported by resolution-go library
func DetectNamingServiceType(domainName string) (string, error) {
	chunks := strings.Split(domainName, ".")
	if len(chunks) == 0 {
		return "", &DomainNotSupported{DomainName: domainName}
	}
	extension := chunks[len(chunks)-1]
	service := supportedNamingServices[extension]
	if service == "" {
		return "", &DomainNotSupported{DomainName: domainName}
	}
	return service, nil
}
