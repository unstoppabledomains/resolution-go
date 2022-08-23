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

# Installing resolution-go

```shell
go get github.com/unstoppabledomains/resolution-go
```

# Updating resolution-go

```shell
go get -u github.com/unstoppabledomains/resolution-go
```

# Usage

```go
package main

import (
  "fmt"
  "github.com/Zilliqa/gozilliqa-sdk/provider"
  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/unstoppabledomains/resolution-go"
  "github.com/unstoppabledomains/resolution-go/namingservice"
)

func main() {
  // Resolve .wallet
	uns, _ := resolution.NewUnsBuilder().SetEthereumNetwork("goerli").SetL2EthereumNetwork("mumbai").Build()
	ethAddress, _ := uns.Resolver("uns-devtest-265f8f.wallet")
	fmt.Println("Address", ethAddress)

  // Resolve .zil
	zns, _ := resolution.NewZnsBuilder().Build()
	btcAddress, _ := zns.Addr("brad.zil", "BTC")
	fmt.Println("BTC address for brad.zil is", btcAddress)

  // Get locations of domains
	locations, _ := uns.Locations([]string{"ryan.crypto", "brad.crypto"})
	fmt.Println("Locations for ryan.crypto and brad.crypto are", locations)

  // Detect domain naming service
	namingServices := map[string]resolution.NamingService{namingservice.UNS: uns, namingservice.ZNS: zns}
	domainToDetect := "brad.zil"
	namingServiceName, _ := resolution.DetectNamingService(domainToDetect)
	if namingServices[namingServiceName] != nil {
		resolvedAddress, _ := namingServices[namingServiceName].Addr(domainToDetect, "ETH")
		fmt.Println("ETH address for", domainToDetect, "is", resolvedAddress)
	}

  // Set custom Ethereum endpoint for UNS service
	ethContractBackend, _ := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/RAQcwz7hhKhmwgoti6HYM_M_9nRJjEsQ")
	builder := resolution.NewUnsBuilder().SetEthereumNetwork("mainnet").SetL2EthereumNetwork("mumbai")
	uns, _ = builder.SetContractBackend(ethContractBackend).SetL2ContractBackend(ethContractBackend).Build()
	allUnsRecords, _ := uns.AllRecords("uns-devtest-265f8f.wallet")
	fmt.Println("Records", allUnsRecords)

  // Set custom Zilliqa endpoint for ZNS service
	zilliqaProvider := provider.NewProvider("https://api.zilliqa.com")
	znsWithCustomProvider, _ := resolution.NewZnsBuilder().SetProvider(zilliqaProvider).Build()
	allZnsRecords, _ := znsWithCustomProvider.AllRecords("brad.zil")
	fmt.Println("Records for brad.zil", allZnsRecords)
}
```

# Custom Ethereum provider configuration

```go
package main

import (
  "github.com/ethereum/go-ethereum/ethclient"
	"github.com/unstoppabledomains/resolution-go/v2"
)

func main() {
  var alchemyApiKey = ALCHEMY_PROJECT_ID
  var ethereumUrl = "https://eth-mainnet.g.alchemy.com/v2/" + alchemyApiKey
  var ethereumL2Url = "https://polygon-mainnet.g.alchemy.com/v2/" + alchemyApiKey

  var unsBuilder := resolution.NewUnsBuilder()
  var backend, _ := ethclient.Dial(ethereumUrl)
  var backendL2, _ := ethclient.Dial(ethereumL2Url)

  unsBuilder.SetContractBackend(backend)
  unsBuilder.SetL2ContractBackend(backendL2)

  var unsResolution, _ = unsBuilder.Build()
  var znsResolution, _ = resolution.NewZnsBuilder().Build()
}
```

# Contributions

Contributions to this library are more than welcome. The easiest way to contribute is through GitHub issues and pull requests.

# Free advertising for integrated apps

Once your app has a working Unstoppable Domains integration, [register it here](https://unstoppabledomains.com/app-submission). Registered apps appear on the Unstoppable Domains [homepage](https://unstoppabledomains.com/) and [Applications](https://unstoppabledomains.com/apps) page — putting your app in front of tens of thousands of potential customers per day.

Also, every week we select a newly-integrated app to feature in the Unstoppable Update newsletter. This newsletter is delivered to straight into the inbox of ~100,000 crypto fanatics — all of whom could be new customers to grow your business.

# Get help

[Join our discord community](https://discord.gg/unstoppabledomains) and ask questions.

# Help us improve

We're always looking for ways to improve how developers use and integrate our products into their applications. We'd love to hear about your experience to help us improve by [taking our survey](https://form.typeform.com/to/uHPQyHO6).
