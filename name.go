package resolution

import s "strings"

// NormalizeName Normalizes domain
func NormalizeName(domain string) string {
	return s.ToLower(s.TrimSpace(domain))
}
