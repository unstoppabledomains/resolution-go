package resolution

// Ens is a naming service handles Ethereum naming service resolution.
type Ens struct {
	service EnsService
}

func (e *Ens) Owner(domainName string) (string, error) {
	return e.service.owner(domainName)
}

func (e *Ens) GetResolver(domainName string) (string, error) {
	return e.service.resolver(domainName)
}

func (e *Ens) Namehash(domainName string) (string, error) {
	return e.service.namehash(domainName)
}
