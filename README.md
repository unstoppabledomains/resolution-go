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
go get github.com/unstoppabledomains/resolution-go/v2
```

# Updating resolution-go

```shell
go get -u github.com/unstoppabledomains/resolution-go/v2
```

# Usage

## Initialize with Default Ethereum Providers (deprecated)
> NOTE: The default Infura key provided is rate limited and should only be used for testing. For production applications, please bring your own Infura or Alchemy RPC URL to prevent downtime.

```go
package main

import (
	"fmt"
	"github.com/unstoppabledomains/resolution-go/v2"
)

func main() {
	// quick start using default RPC providers
	uns, _ := resolution.NewUnsBuilder().Build()
}
```

## Initialize with Unstoppable Domains' Proxy Provider

```go
package main

import (
	"fmt"
	"github.com/unstoppabledomains/resolution-go/v2"
)

// obtain a key from https://unstoppabledomains.com/partner-api-dashboard if you are a partner
uns, _ := resolution.NewUnsBuilder().SetUdClient("<api_key>").Build()

```

## Initialize with Custom Ethereum Provider Configuration

```go
package main

import (
	"fmt"

	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/unstoppabledomains/resolution-go/v2"
)

func main() {
	// obtain a key from https://www.infura.io
	var ethereumUrl = "https://mainnet.infura.io/v3/<infura_api_key>"
	var ethereumL2Url = "https://polygon-mainnet.infura.io/v3/<infura_api_key>"

	var unsBuilder := resolution.NewUnsBuilder()
	var backend, _ := ethclient.Dial(ethereumUrl)
	var backendL2, _ := ethclient.Dial(ethereumL2Url)

	unsBuilder.SetContractBackend(backend)
	unsBuilder.SetL2ContractBackend(backendL2)

	var uns, _ = unsBuilder.Build()

	zilliqaProvider := provider.NewProvider("https://api.zilliqa.com")
	zns, _ := resolution.NewZnsBuilder().SetProvider(zilliqaProvider).Build()
}
```

## Resolve domains example

```go
package main

import (
	"fmt"

	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/unstoppabledomains/resolution-go/v2"
	"github.com/unstoppabledomains/resolution-go/v2/namingservice"
)

func main() {
	// Resolve .crypto
	uns, _ := resolution.NewUnsBuilder().Build()
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

# Contributions

Contributions to this library are more than welcome. The easiest way to contribute is through GitHub issues and pull requests.

# Free advertising for integrated apps

Once your app has a working Unstoppable Domains integration, [register it here](https://unstoppabledomains.com/app-submission). Registered apps appear on the Unstoppable Domains [homepage](https://unstoppabledomains.com/) and [Applications](https://unstoppabledomains.com/apps) page — putting your app in front of tens of thousands of potential customers per day.

Also, every week we select a newly-integrated app to feature in the Unstoppable Update newsletter. This newsletter is delivered to straight into the inbox of ~100,000 crypto fanatics — all of whom could be new customers to grow your business.

# Get help

[Join our discord community](https://discord.gg/unstoppabledomains) and ask questions.

# Help us improve

We're always looking for ways to improve how developers use and integrate our products into their applications. We'd love to hear about your experience to help us improve by [taking our survey](https://form.typeform.com/to/uHPQyHO6).
