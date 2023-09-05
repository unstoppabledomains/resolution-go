package utils

import (
	"strings"
)

func SplitDomain(domain string) (labels []string, extension string) {
	chunks := strings.Split(domain, ".")

	if len(chunks) < 2 {
		extension = domain
		return labels, extension
	}

	extension = chunks[len(chunks)-1]

	labels = chunks[:len(chunks)-1]

	return labels, extension
}
