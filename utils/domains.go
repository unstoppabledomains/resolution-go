package utils

import (
	"strings"
)

func SplitDomain(domain string) (labels string, extension string) {
	chunks := strings.Split(domain, ".")

	if len(chunks) < 2 {
		extension = domain
		return labels, extension
	}

	extension = chunks[len(chunks)-1]

	labels = strings.Join(chunks[:len(chunks)-1], ".")

	return labels, extension
}

func IsSubdomain(domain string) bool {
	return len(strings.Split(domain, ".")) > 2
}

func GetParentDomain(domain string) string {
	chunks := strings.Split(domain, ".")

	if len(chunks) < 2 {
		return ""
	}

	return strings.Join(chunks[1:], ".")
}

func NormalizeName(domain string) string {
	return strings.ToLower(strings.TrimSpace(domain))
}
