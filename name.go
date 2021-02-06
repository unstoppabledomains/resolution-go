package resolution

import s "strings"

func NormalizeName(domain string) string {
	return s.ToLower(s.TrimSpace(domain))
}
