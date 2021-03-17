package resolution

import s "strings"

func normalizeName(domain string) string {
	return s.ToLower(s.TrimSpace(domain))
}
