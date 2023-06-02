package namingservice

const (
	UNS = "UNS"
	ZNS = "ZNS"
)

type Location struct {
	RegistryAddress       string
	ResolverAddress       string
	NetworkId             int
	Blockchain            string
	OwnerAddress          string
	BlockchainProviderUrl string
}
