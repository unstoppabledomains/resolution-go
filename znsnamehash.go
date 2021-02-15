package resolution

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func ZnsNameHash(domainName string) (string, error) {
	normalizedName := NormalizeName(domainName)
	domainNodes := strings.Split(normalizedName, ".")
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
