package utils

import (
	"encoding/hex"
	"errors"
	"math/big"
	"strings"
)

var base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func EncodeBase58(input string) (string, error) {
	bytes, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}
	intInput := big.NewInt(0).SetBytes(bytes)
	result := ""

	// Set up a big.Int for the divisor
	divisor := big.NewInt(58)

	// Create big.Int for quotient and remainder
	quotient := big.NewInt(0)
	remainder := big.NewInt(0)

	for intInput.Cmp(big.NewInt(0)) > 0 {
		quotient, remainder = quotient.QuoRem(intInput, divisor, remainder)
		result = string(base58Alphabet[remainder.Int64()]) + result
		intInput = quotient
	}

	// handle leading zeros
	for _, b := range bytes {
		if b != 0 {
			break
		}
		result = "1" + result
	}

	return result, nil
}

func DecodeENSContentHash(contentHash string) (string, error) {
	contentHash = strings.TrimPrefix(contentHash, "0x")

	switch {
	case strings.HasPrefix(contentHash, "e3010170"): // IPFS
		hashPart := contentHash[8:]
		ipfsHash, err := EncodeBase58(hashPart)
		if err != nil {
			return "", err
		}
		return "ipfs://" + ipfsHash, nil
	case strings.HasPrefix(contentHash, "e40101fa011b20"): // Swarm
		hashPart := contentHash[16:]
		return "bzz://" + hashPart, nil
	default:
		return "", errors.New("unknown content hash type")
	}
}
