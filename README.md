![Test](https://github.com/unstoppabledomains/resolution-go/workflows/Test/badge.svg?branch=master)
![Lint](https://github.com/unstoppabledomains/resolution-go/workflows/Lint/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/unstoppabledomains/resolution-go)](https://goreportcard.com/report/github.com/unstoppabledomains/resolution-go)
[![GoDoc](https://godoc.org/github.com/unstoppabledomains/resolution-go?status.svg)](https://pkg.go.dev/github.com/unstoppabledomains/resolution-go)
[![Unstoppable Domains Documentation](https://img.shields.io/badge/docs-unstoppabledomains.com-blue)](https://docs.unstoppabledomains.com/)
[![Get help on Discord](https://img.shields.io/badge/Get%20help%20on-Discord-blueviolet)](https://discord.gg/b6ZVxSZ9Hn)

# resolution-go

resolution-go is a library for interacting with blockchain domain names. It can be used to retrieve [payment addresses](https://unstoppabledomains.com/learn/how-to-send-crypto-using-your-domain) and IPFS hashes for [decentralized websites](https://support.unstoppabledomains.com/support/solutions/articles/48001181925-build-website).

resolution-go is primarily built and maintained by [Unstoppable Domains](https://unstoppabledomains.com/).

Resolution supports different decentralized domains. Please, refer to the [Top Level Domains List](https://resolve.unstoppabledomains.com/supported_tlds)

- [Installing Resolution](#installing-resolution-go)
- [Updating Resolution](#updating-resolution-go)
- [Using Resolution](#using-resolution)
  - [Web3Domain Resolution](#web3domain-resolution)
  - [Unstoppable Domain Resolution](#unstoppable-domain-resolution)
- [Contributions](#contributions)
- [Free advertising for integrated apps](#free-advertising-for-integrated-apps)

# Installing resolution-go

```shell
go get github.com/unstoppabledomains/resolution-go/v3
```

# Updating resolution-go

```shell
go get -u github.com/unstoppabledomains/resolution-go/v3
```

# Using Resolution

---
**NOTE**

From version `3.2.0`, we introduce a new interface to handle multiple naming services domain resolution.

If you wish to migrate to the new interface to be able resolve other naming services such as [ENS](https://ens.domains/), please reference the new interface.

---

## Web3Domain Resolution

We are currently support the following naming services. More naming service on other blockchain are coming...
* UNS: `.x`, `.crypto`, `.go` etc... 
* ZNS: `.zil`
* ENS: `.eth`, `.kred`, `.xyz`, `.luxe`

### Initialize with Unstoppable Domains' Proxy Provider

```go
package main

import (
	"fmt"
	"github.com/unstoppabledomains/resolution-go/v3"
)

func main() {
	// obtain a key from https://unstoppabledomains.com/partner-api-dashboard if you are a partner
	web3domain, _ := resolution.NewWeb3DomainBuilder().SetUdClient("<api_key>").Build()
}

```

## Initialize with Custom Providers Configuration

```go
package main

import (
	"fmt"

	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/unstoppabledomains/resolution-go/v3"
)

func main() {
	// obtain a key from https://www.infura.io
	ethereumUrl := "https://mainnet.infura.io/v3/<infura_api_key>"
	maticUrl := "https://polygon-mainnet.infura.io/v3/<infura_api_key>"

	web3DomainBuilder := resolution.NewWeb3DomainBuilder()
	ethBackend, _ := ethclient.Dial(ethereumUrl)
	maticBackend, _ := ethclient.Dial(maticUrl)

	web3DomainBuilder.SetEthContractBackend(backend)
	web3DomainBuilder.SetMaticContractBackend(backendL2)

	w3bDomain, _ = web3DomainBuilder.Build()
}
```
### Resolve web3 domains example

```go
package main

import (
	"fmt"

	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/unstoppabledomains/resolution-go/v3"
	"github.com/unstoppabledomains/resolution-go/v3/namingservice"
)

func main() {
	web3Domain, _ := resolution.NewWeb3DomainBuilder().SetUdClient("<api_key>").Build()

	// .x domain - UNS naming service
	unsDomain := "lisa.x"
	lisaAddress, _ := web3Domain.Owner(unsDomain)
	fmt.Printf("Owner of %s is %s\n", unsDomain, lisaAddress)
	lisaEthAddress, _ := web3Domain.Addr(unsDomain, "ETH")
	fmt.Printf("ETH address of %s is %s\n", unsDomain, lisaEthAddress)

	// .eth domain - ENS naming service
	ensDomain := "vitalik.eth"
	vitalikAddress, _ := web3Domain.Owner(ensDomain)
	fmt.Printf("Owner of %s is %s\n", ensDomain, vitalikAddress)
	vitalikEthAddress, _ := web3Domain.Addr(ensDomain, "ETH")
	fmt.Printf("ETH address of %s is %s\n", ensDomain, vitalikEthAddress)

	// .zil domain - ZNS naming service
	znsDomain := "brad.zil"
	bradAddress, _ := web3Domain.Owner(znsDomain)
	fmt.Printf("Owner of %s is %s\n", znsDomain, bradAddress)
	bradEthAddress, _ := web3Domain.Addr(znsDomain, "ETH")
	fmt.Printf("ETH address of %s is %s\n", znsDomain, bradEthAddress)

	//reverse resolution
	address1 := "0x05391f2407B664fbd1dcA5AEa9eEa89A29B946b4"
	foundDomain, _ := web3Domain.ReverseOf(address1)
	fmt.Printf("foundDomain for address %s is %s\n", address1, foundDomain) // tun.x

	address2 := "0x4309325e607de185bd6b091cc2e3cfc9f4d6e2e1"
	foundEthDomain, _ := web3Domain.ReverseOf(address2)
	fmt.Printf("foundDomain for address %s is %s\n", address2, foundEthDomain) // unimev.eth

	foundDomains, _ := web3Domain.MultiReverseOf(address1) 	// find a list of domains the address revere resolution to
	fmt.Printf("foundDomain for address %s is %v\n", address1, foundDomains) // [tun.x tu-nguyen.eth]


	//get addr records
	coinAddressUnsDomain, _ := web3Domain.Addr(unsDomain, "ETH")
	fmt.Printf("ETH address of %s is %s\n", unsDomain, coinAddressUnsDomain)

	coinAddressEnsDomain, _ := web3Domain.Addr(ensDomain, "ETH")
	fmt.Printf("ETH address of %s is %s\n", ensDomain, coinAddressEnsDomain)

	// For ENS domains, it's crucial to know that a domain is still valid
	ensExpiry, _ := web3Domain.DomainExpiry(ensDomain)
	fmt.Printf("Is %s expired %t\n", ensDomain, ensExpiry.Before(time.Now()))
}
```

## Unstoppable Domain Resolution

### Initialize with Unstoppable Domains' Proxy Provider

```go
package main

import (
	"fmt"
	"github.com/unstoppabledomains/resolution-go/v3"
)

func main() {
	// obtain a key from https://unstoppabledomains.com/partner-api-dashboard if you are a partner
	uns, _ := resolution.NewUnsBuilder().SetUdClient("<api_key>").Build()
}

```

### Initialize with Custom Ethereum Provider Configuration

```go
package main

import (
	"fmt"

	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/unstoppabledomains/resolution-go/v3"
)

func main() {
	// obtain a key from https://www.infura.io
	var ethereumUrl = "https://mainnet.infura.io/v3/<infura_api_key>"
	var ethereumL2Url = "https://polygon-mainnet.infura.io/v3/<infura_api_key>"

	unsBuilder := resolution.NewUnsBuilder()
	backend, _ := ethclient.Dial(ethereumUrl)
	backendL2, _ := ethclient.Dial(ethereumL2Url)

	unsBuilder.SetContractBackend(backend)
	unsBuilder.SetL2ContractBackend(backendL2)

	uns, _ = unsBuilder.Build()

	zilliqaProvider := provider.NewProvider("https://api.zilliqa.com")
	zns, _ := resolution.NewZnsBuilder().SetProvider(zilliqaProvider).Build()
}
```

### Resolve domains example

```go
package main

import (
	"fmt"

	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/unstoppabledomains/resolution-go/v3"
	"github.com/unstoppabledomains/resolution-go/v3/namingservice"
)

func main() {
	// Resolve .crypto
	uns, _ := resolution.NewUnsBuilder().SetUdClient("<api_key>").Build()
	ethAddress, _ := uns.Addr("brad.crypto", "ETH")
	fmt.Println("ETH address for brad.crypto is", ethAddress)

	// Resolve.zil
	zns, _ := resolution.NewZnsBuilder().Build()
	btcAddress, _ := zns.Addr("brad.zil", "BTC")
	fmt.Println("BTC address for brad.zil is", btcAddress)

	// Get locations of domains
	uns, _ = resolution.NewUnsBuilder().Build()
	locations, _ := uns.Locations([]string{"ryan.crypto", "brad.crypto"})
	fmt.Println("Locations for ryan.crypto and brad.crypto are", locations)

	// Detect domain naming service
	namingServices := map[string]resolution.NamingService{namingservice.UNS: uns, namingservice.ZNS: zns}
	domainToDetect := "ryan.crypto"
	namingServiceName, _ := resolution.DetectNamingService(domainToDetect)
	if namingServices[namingServiceName] != nil {
		resolvedAddress, _ := namingServices[namingServiceName].Addr(domainToDetect, "ETH")
		fmt.Println("ETH address for", domainToDetect, "is", resolvedAddress)
	}
}
```

### Resolve Wallet Address Examples

---
**NOTE**

These example are only applied to UNS using `NewUnsBuilder` function

---

#### Using **`Addr`**

This API is used to retrieve wallet address for single address record. (See
[Cryptocurrency payment](https://docs.unstoppabledomains.com/resolution/guides/records-reference/#cryptocurrency-payments)
section for the record format)

```go
func main() {
	//...
	// homecakes.crypto has `crypto.ETH.address` set to 0xe7474D07fD2FA286e7e0aa23cd107F8379085037
	domain := "homecakes.crypto"
	walletAddress, err := uns.Addr(domain, "ETH")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Addr for %s ticker %s: %s\n", domain, "ETH", walletAddress)
	// Addr for homecakes.crypto ticker ETH: 0xe7474D07fD2FA286e7e0aa23cd107F8379085037
}
```

#### Using **`GetAddr`**

This (beta) API can be used to resolve different formats

**Resolve single address format (similar to **`Addr`** API)**

With `homecakes.crypto` has a `crypto.ETH.address` record set on-chain:

```go
func main() {
	domain := "homecakes.crypto"
	walletAddress, err := uns.GetAddr(domain, "ETH", "ETH")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Addr for %s on network %s with token %s: %s\n", domain, "ETH", "ETH", walletAddress)
	// Addr for homecakes.crypto on network ETH with token ETH: 0xe7474D07fD2FA286e7e0aa23cd107F8379085037
}
```

**Resolve multi-chain currency address format (See
[multi-chain currency](https://docs.unstoppabledomains.com/resolution/guides/records-reference/#multi-chain-currencies))**

With `aaron.x` has a `crypto.AAVE.version.ERC20.address` record set to
`0xCD0DAdAb45bAF9a06ce1279D1342EcC3F44845af`. The `ERC20` indicates it's a token
on `ETH` network:

```go
func main() {
	domain := "aaron.x"
	walletAddress, err := uns.GetAddr(domain, "ETH", "AAVE")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Addr for %s on network %s with token %s: %s\n", domain, "ETH", "AAVE", walletAddress)
	// Addr for aaron.x on network ETH with token AAVE: 0xCD0DAdAb45bAF9a06ce1279D1342EcC3F44845af
}
```

**Derive wallet addresses within the same blockchain network and blockchain
family.**

The API can also be used by crypto exchanges to infer wallet addresses. In
centralized exchanges, users have same wallet addresses on different networks with same
wallet family. (See [Blockchain Family, Network, Token Level Addresses](https://apidocs.unstoppabledomains.com/resolution/guides/records-reference/#blockchain-family-network-token-level-addresses) section for the record format)

With `blockchain-family-keys.x` only has `token.EVM.address` record on-chain.
The API resolves to same wallet address for tokens live on EVM compatible
networks.

```go

func main() {
	domain := "blockchain-family-keys.x"
	aaveOnEthWallet, _ := uns.GetAddr(domain, "ETH", "AAVE")
	fmt.Printf("Addr for %s on network %s with token %s: %s\n", domain, "ETH", "AAVE", aaveOnEthWallet)
	// Addr for blockchain-family-keys.x on network ETH with token AAVE: 0xCD0DAdAb45bAF9a06ce1279D1342EcC3F44845af

	ethOnEthWallet, _ := uns.GetAddr(domain, "ETH", "ETH")
	fmt.Printf("Addr for %s on network %s with token %s: %s\n", domain, "ETH", "ETH", ethOnEthWallet)
	// Addr for blockchain-family-keys.x on network ETH with token ETH: 0xCD0DAdAb45bAF9a06ce1279D1342EcC3F44845af

	usdtOnAvax, _ := uns.GetAddr(domain, "AVAX", "USDT")
	fmt.Printf("Addr for %s on network %s with token %s: %s\n", domain, "AVAX", "USDT", usdtOnAvax)
	// Addr for blockchain-family-keys.x on network AVAX with token USDT: 0xCD0DAdAb45bAF9a06ce1279D1342EcC3F44845af
}
```

With `uns-devtest-nickshatilo-withdraw-test2.x` only has `token.EVM.ETH.address`
record on chain. The API resolves to the same wallet address for tokens
specifically on Ethereum network.

```go
func main() {
	domain := "uns-devtest-nickshatilo-withdraw-test2.x"
	aaveOnEthWallet, _ := uns.GetAddr(domain, "ETH", "AAVE")
	fmt.Printf("Addr for %s on network %s with token %s: %s\n", domain, "ETH", "AAVE", aaveOnEthWallet)
	// Addr for uns-devtest-nickshatilo-withdraw-test2.x on network ETH with token AAVE: 0xCD0DAdAb45bAF9a06ce1279D1342EcC3F44845af

	ethOnEthWallet, _ := uns.GetAddr(domain, "ETH", "ETH")
	fmt.Printf("Addr for %s on network %s with token %s: %s\n", domain, "ETH", "ETH", ethOnEthWallet)
	// Addr for uns-devtest-nickshatilo-withdraw-test2.x on network ETH with token ETH: 0xCD0DAdAb45bAF9a06ce1279D1342EcC3F44845af

	usdtOnAvax, _ := uns.GetAddr(domain, "AVAX", "USDT")
	// won't work
}
```

The API is compatible with other address formats. If a domain has multiple
address formats set, it will follow the algorithm described as follow:

if a domain has following records set:

```
token.EVM.address
crypto.USDC.version.ERC20.address
token.EVM.ETH.USDC.address
crypto.USDC.address
token.EVM.ETH.address
```

`getAddress(domain, 'ETH', 'USDC')` will lookup records in the following order:

```
1. token.EVM.ETH.USDC.address
2. crypto.USDC.address
3. crypto.USDC.version.ERC20.address
4. token.EVM.ETH.address
5. token.EVM.address
```

# Contributions

Contributions to this library are more than welcome. The easiest way to contribute is through GitHub issues and pull requests.

Use these commands to set up a local development environment (**macOS Terminal**
or **Linux shell**).

1. Recommended golang version

- go1.18

2. Clone the repository

   ```bash
   git clone https://github.com/unstoppabledomains/resolution-go.git
   cd resolution-go
   ```

3. Install dependencies

   ```bash
   go mod download
   ```

### Internal config

#### Unit tests:

**resolution-go** library relies on environment variables to load **TestNet** RPC Urls. This way, our keys don't expose directly to the code. In order to validate the code change, please set these variables to your local environment.

- L1_TEST_NET_RPC_URL
- L2_TEST_NET_RPC_URL

# Free advertising for integrated apps

Once your app has a working Unstoppable Domains integration, [register it here](https://unstoppabledomains.com/app-submission). Registered apps appear on the Unstoppable Domains [homepage](https://unstoppabledomains.com/) and [Applications](https://unstoppabledomains.com/apps) page — putting your app in front of tens of thousands of potential customers per day.

Also, every week we select a newly-integrated app to feature in the Unstoppable Update newsletter. This newsletter is delivered to straight into the inbox of ~100,000 crypto fanatics — all of whom could be new customers to grow your business.

# Get help

[Join our discord community](https://discord.gg/unstoppabledomains) and ask questions.

# Help us improve

We're always looking for ways to improve how developers use and integrate our products into their applications. We'd love to hear about your experience to help us improve by [taking our survey](https://form.typeform.com/to/uHPQyHO6).
