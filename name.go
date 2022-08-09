package resolution

import s "strings"

func normalizeName(domain string) string {
	return s.ToLower(s.TrimSpace(domain))
}

func normalizeAndVerifyName(domain string) (string, error) {
	domainName := normalizeName(domain)
	if s.HasSuffix(domainName, ".coin") {
		return "", &DomainNotSupportedError{DomainName: domainName}
	}
	return domainName, nil
}