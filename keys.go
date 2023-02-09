package resolution

import (
	"encoding/json"
	"fmt"
	"strings"
  "embed"
)

// supportedKeys struct of supported keys
type supportedKeys map[string]struct {
	DeprecatedKeyName string
	Deprecated        bool
	ValidationRegex   string
}

const emailKey = "whois.email.value"

var ipfsKeys = []string{"dweb.ipfs.hash", "ipfs.html.value"}
var redirectUrlKeys = []string{"browser.redirect_url", "ipfs.redirect_domain.value"}

// buildCryptoKey returns raw key for crypto currency which is used to query blockchain
func buildCryptoKey(ticker string) (string, error) {
	var key strings.Builder
	_, err := fmt.Fprintf(&key, "crypto.%s.address", strings.ToUpper(ticker))
	if err != nil {
		return "", err
	}
	return key.String(), nil
}

// buildCryptoKeyVersion returns raw key for multi-chain currency which is used to query blockchain
func buildCryptoKeyVersion(ticker string, version string) (string, error) {
	var key strings.Builder
	_, err := fmt.Fprintf(&key, "crypto.%s.version.%s.address", strings.ToUpper(ticker), strings.ToUpper(version))
	if err != nil {
		return "", err
	}
	return key.String(), nil
}

// returnFirstNonEmpty returns first not empty elements from provided records and keys order
func returnFirstNonEmpty(records map[string]string, keysSequence []string) string {
	for _, key := range keysSequence {
		if records[key] != "" {
			return records[key]
		}
	}

	return ""
}

// newSupportedKeys returns supportedKeys
func newSupportedKeys() (supportedKeys, error) {
	var keysObject struct {
		Keys supportedKeys
	}
	err := json.Unmarshal(supportedKeysJSON, &keysObject)
	if err != nil {
		return nil, err
	}
	return keysObject.Keys, nil
}

//go:embed uns/resolver-keys.json
var f embed.FS
var supportedKeysJSON, _ := f.ReadFile("uns/resolver-keys.json")
