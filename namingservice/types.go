package namingservice

const (
	UNS = "UNS"
	ZNS = "ZNS"
	ENS = "ENS"
)

type Location struct {
	RegistryAddress       string
	ResolverAddress       string
	NetworkId             int
	Blockchain            string
	OwnerAddress          string
	BlockchainProviderUrl string
}
