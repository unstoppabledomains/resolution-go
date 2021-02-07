package resolution

type DomainNotRegistered struct {
	DomainName string
}

type DomainNotConfigured struct {
	DomainName string
}

func (e *DomainNotRegistered) Error() string { return e.DomainName + " is not registered" }
func (e *DomainNotConfigured) Error() string {
	return e.DomainName + " does not have configured Resolver"
}
