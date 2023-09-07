package resolution

import (
	"fmt"
	"strings"

	"github.com/unstoppabledomains/resolution-go/v3/dnsrecords"
	"github.com/unstoppabledomains/resolution-go/v3/namingservice"
	"github.com/unstoppabledomains/resolution-go/v3/utils"
)

// NamingService Unstoppable supports multiple naming services (.zil and .crypto).
// Each naming service implements shared interface and returns similar record types.
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
	// Deprecated: This method will be removed in future releases
	AllRecords(domainName string) (map[string]string, error)

	// Locations Retrieve locations of domains
	// Returns key-value map of domain names to location
	Locations(domainNames []string) (map[string]namingservice.Location, error)

	// DNS Retrieve the DNS records of a domain.
	// Returns a set of valid and filtered non-empty DNS records attached to domain.
	DNS(domainName string, types []dnsrecords.Type) ([]dnsrecords.Record, error)

	// IsSupportedDomain checks whether domain name is supported by the naming service.
	IsSupportedDomain(domainName string) (bool, error)

	// TokenURI returns ERC721 metadata token URI
	TokenURI(domainName string) (string, error)

	// TokenURIMetadata returns ERC721 metadata
	TokenURIMetadata(domainName string) (TokenMetadata, error)

	// Unhash returns a domain name from a hash using TokenMetadata.Name field and ensures it indeed matches the given hash.
	// domainHash should be in hex numeric string format, for example: "0x29bf1b111e709f0953848df35e419490fbad5d316690e4de61adc52695ddf9f3"
	// ERC721 Token ID could be passed to this method but should be converted to hex numeric string before usage.
	//
	// Examples of usage:
	//
	// domainName, err := NamingService.Unhash("0x29bf1b111e709f0953848df35e419490fbad5d316690e4de61adc52695ddf9f3")
	// domainName, err := NamingService.Unhash("0x691f36df38168d9297e784f45a87257a70c58c4040d469c6d0b91d253a837e32")
	//
	// Usage with ERC721 token id:
	//
	// var erc721TokenID big.Int
	// erc721TokenID.SetString("47548000072528700265403562077742902343248290986511625310517899838602191535666", 10)
	// domainHash := hex.EncodeToString(erc721TokenID.Bytes())
	// domainName, err := NamingService.Unhash(domainHash)
	//
	Unhash(domainHash string) (string, error)

	// Namehash returns a namehash of a domain following the EIP-137 standard
	Namehash(domainName string) (string, error)

	ReverseOf(addr string) (string, error)
}

// DetectNamingService helper to detect naming service type for provided domain.
// Returns ZNS or UNS for valid domain and error if domain is not valid or not supported by resolution-go library.
func DetectNamingService(domainName string) (string, error) {
	chunks := strings.Split(domainName, ".")
	if len(chunks) < 2 {
		return "", &DomainNotSupportedError{DomainName: domainName}
	}
	extension := chunks[len(chunks)-1]
	if len(extension) == 0 {
		return "", &DomainNotSupportedError{DomainName: domainName}
	}
	if extension == "zil" {
		return namingservice.ZNS, nil
	}
	return namingservice.UNS, nil
}

type Web3Domain struct {
	uns *Uns
	ens *Ens
}

func (w3d *Web3Domain) getNamingServiceForDomain(domain string) NamingService {
	normalizedDomain := utils.NormalizeName(domain)

	_, tld := utils.SplitDomain(normalizedDomain)

	switch tld {
	case "eth", "kred", "luxe", "xyz":
		return w3d.ens
	default:
		return w3d.uns
	}

}

func (w3d *Web3Domain) Owner(domain string) (string, error) {
	namingService := w3d.getNamingServiceForDomain(domain)

	if namingService == nil {
		return "", &DomainNotSupportedError{DomainName: domain}
	}

	return namingService.Owner(domain)
}

func (w3d *Web3Domain) Resolver(domain string) (string, error) {
	namingService := w3d.getNamingServiceForDomain(domain)

	if namingService == nil {
		return "", &DomainNotSupportedError{DomainName: domain}
	}

	return namingService.Resolver(domain)
}

func (w3d *Web3Domain) Namehash(domain string) (string, error) {
	namingService := w3d.getNamingServiceForDomain(domain)

	if namingService == nil {
		return "", &DomainNotSupportedError{DomainName: domain}
	}

	return namingService.Namehash(domain)
}

func (w3d *Web3Domain) ReverseOf(addr string) (string, error) {
	namingServices := []NamingService{w3d.uns, w3d.ens}

	for _, namingService := range namingServices {
		reverse, err := namingService.ReverseOf(addr)
		if err == nil {
			return reverse, nil
		}
	}

	return "", nil
}

func (w3d *Web3Domain) TokenURI(domain string) (string, error) {
	namingService := w3d.getNamingServiceForDomain(domain)

	if namingService == nil {
		return "", &DomainNotSupportedError{DomainName: domain}
	}

	return namingService.TokenURI(domain)
}

func (w3d *Web3Domain) Addr(domain, token string) (string, error) {
	namingService := w3d.getNamingServiceForDomain(domain)

	if namingService == nil {
		return "", &DomainNotSupportedError{DomainName: domain}
	}

	return namingService.Addr(domain, token)
}

func (w3d *Web3Domain) IpfsHash(domain string) (string, error) {
	namingService := w3d.getNamingServiceForDomain(domain)

	if namingService == nil {
		return "", &DomainNotSupportedError{DomainName: domain}
	}

	return namingService.IpfsHash(domain)
}

func (w3d *Web3Domain) HTTPUrl(domain string) (string, error) {
	namingService := w3d.getNamingServiceForDomain(domain)

	if namingService == nil {
		return "", &DomainNotSupportedError{DomainName: domain}
	}

	return namingService.HTTPUrl(domain)
}

func (w3d *Web3Domain) Email(domain string) (string, error) {
	namingService := w3d.getNamingServiceForDomain(domain)

	if namingService == nil {
		return "", &DomainNotSupportedError{DomainName: domain}
	}

	return namingService.Email(domain)
}

// maybe useless
func (w3d Web3Domain) Locations(domains []string) (map[string]namingservice.Location, error) {
	var unsDomains []string
	var ensDomains []string

	for _, domain := range domains {
		_, tld := utils.SplitDomain(domain)

		switch tld {
		case "eth", "kred", "luxe", "xyz":
			ensDomains = append(ensDomains, domain)
		default:
			unsDomains = append(unsDomains, domain)
		}
	}

	unsResult, err := w3d.uns.Locations(unsDomains)

	if err != nil {
		return nil, err
	}
	ensResult, err := w3d.ens.Locations(ensDomains)

	if err != nil {
		return nil, err
	}

	fmt.Println(unsResult)

	result := make(map[string]namingservice.Location)

	for k, v := range unsResult {
		result[k] = v
	}

	for k, v := range ensResult {
		result[k] = v
	}

	return result, nil
}

func (w3d *Web3Domain) DNS(domain string, types []dnsrecords.Type) ([]dnsrecords.Record, error) {
	namingService := w3d.getNamingServiceForDomain(domain)

	if namingService == nil {
		return nil, &DomainNotSupportedError{DomainName: domain}
	}

	return namingService.DNS(domain, types)
}
