package resolution

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// ZnsNameHash Namehash for .zil domains
func ZnsNameHash(domainName string) (string, error) {
	domainNodes := strings.Split(domainName, ".")
	hash := sha256.New()
	var finalNamehash = make([]byte, 32)
	for i := len(domainNodes) - 1; i >= 0; i-- {
		hash.Reset()
		_, err := hash.Write([]byte(domainNodes[i]))
		if err != nil {
			return "", err
		}
		currentHash := hash.Sum(finalNamehash)
		hash.Reset()
		_, err = hash.Write(currentHash)
		if err != nil {
			return "", err
		}
		finalNamehash = hash.Sum(nil)
	}
	var hexNamehash strings.Builder
	_, err := fmt.Fprintf(&hexNamehash, "0x%v", hex.EncodeToString(finalNamehash))
	if err != nil {
		return "", err
	}

	return hexNamehash.String(), nil
}
