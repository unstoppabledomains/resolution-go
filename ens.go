package resolution

import (
	"math/big"
	"strconv"
	"strings"

	"github.com/unstoppabledomains/resolution-go/v3/utils"
)

// Ens is a naming service handles Ethereum naming service resolution.
type Ens struct {
	service EnsService
}

func (e *Ens) IsSupportedDomain(domainName string) (bool, error) {
	if domainName == "" {
		return false, nil
	}

	_, extension := utils.SplitDomain(domainName)

	return e.service.domainExists(e.service.namehash(extension))
}

func (e *Ens) Namehash(domainName string) string {
	normalizeName := normalizeName(domainName)
	return e.service.namehash(normalizeName).String()
}

func (e *Ens) GetResolver(domainName string) (string, error) {
	normalizeName := normalizeName(domainName)
	resolverAddress, err := e.service.resolver(e.service.namehash(normalizeName))

	if err != nil {
		return "", err
	}

	if resolverAddress == NullAddress {
		return "", &DomainNotConfiguredError{DomainName: normalizeName}
	}

	return resolverAddress, nil
}

func (e *Ens) ReverseOf(addr string) (string, error) {
	return e.service.reverseOf(addr)
}

func (e *Ens) Owner(domainName string) (string, error) {
	normalizeName := normalizeName(domainName)
	return e.service.ownerOf(e.service.namehash(normalizeName))
}

func (e *Ens) Addr(domainName string) (string, error) {
	normalizeName := normalizeName(domainName)

	resolverAddress, err := e.GetResolver(domainName)

	if err != nil {
		return "", err
	}

	return e.service.addrRecord(resolverAddress, e.service.namehash(normalizeName))
}

func (e *Ens) CoinAddress(domainName string, coin string) (string, error) {
	normalizeName := normalizeName(domainName)

	var coinNum *big.Int
	if strings.HasPrefix(coin, "0x8") { // hexadecimal representation
		num, err := strconv.ParseInt(coin[3:], 16, 64)

		if err != nil {
			return "", err
		}

		coinNum = big.NewInt(num)
	} else {
		coinNum = new(big.Int)
		_, ok := coinNum.SetString(coin, 10)

		if !ok {
			return "", &EnsInvalidCoinType{CoinType: coin}
		}
	}

	resolverAddress, err := e.GetResolver(domainName)

	if err != nil {
		return "", err
	}

	return e.service.addrCoinRecord(resolverAddress, e.service.namehash(normalizeName), coinNum)
}

func (e *Ens) ContentHash(domainName string) (string, error) {
	normalizeName := normalizeName(domainName)
	resolverAddress, err := e.GetResolver(domainName)
	if err != nil {
		return "", err
	}

	return e.service.contenthashRecord(resolverAddress, e.service.namehash(normalizeName))
}

func (e *Ens) TextRecord(domainName, key string) (string, error) {
	normalizeName := normalizeName(domainName)
	resolverAddress, err := e.GetResolver(domainName)

	if err != nil {
		return "", err
	}

	return e.service.textRecord(resolverAddress, e.service.namehash(normalizeName), key)
}
