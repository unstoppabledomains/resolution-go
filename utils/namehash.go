package utils

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func UnsEnsNameHash(domainName string) common.Hash {
	node := common.Hash{}

	if len(domainName) > 0 {
		labels := strings.Split(domainName, ".")

		for i := len(labels) - 1; i >= 0; i-- {
			labelSha := crypto.Keccak256Hash([]byte(labels[i]))
			node = crypto.Keccak256Hash(node.Bytes(), labelSha.Bytes())
		}
	}

	return node
}

func Erc721Hash(label string) common.Hash {
	return crypto.Keccak256Hash([]byte(label))
}
