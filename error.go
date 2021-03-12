package resolution

// DomainNotRegistered Error when domain is missing an owner

type DomainNotRegistered struct {
	DomainName string
}

// DomainNotConfigured Error when domain does not have a resolver set
type DomainNotConfigured struct {
	DomainName string
}

// DomainNotSupported Error when domain is not supported by the naming service
type DomainNotSupported struct {
	DomainName string
}

func (e *DomainNotRegistered) Error() string {
	return e.DomainName + " is not registered"
}
func (e *DomainNotConfigured) Error() string {
	return e.DomainName + " does not have configured Resolver"
}
func (e *DomainNotSupported) Error() string {
	return e.DomainName + " is not supported by naming service"
}
