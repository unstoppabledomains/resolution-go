package resolution

// DomainNotRegisteredError Error when domain is missing an owner
type DomainNotRegisteredError struct {
	DomainName string
}

// DomainNotConfiguredError Error when domain does not have a resolver set
type DomainNotConfiguredError struct {
	DomainName string
}

// DomainNotSupportedError Error when domain is not supported by the naming service
type DomainNotSupportedError struct {
	DomainName string
}

type MethodIsNotSupportedError struct {
	NamingServiceName string
}

func (e *DomainNotRegisteredError) Error() string {
	return e.DomainName + " is not registered"
}
func (e *DomainNotConfiguredError) Error() string {
	return e.DomainName + " does not have configured Resolver"
}
func (e *DomainNotSupportedError) Error() string {
	return e.DomainName + " is not supported by naming service"
}

func (e *MethodIsNotSupportedError) Error() string {
	return "Method is not supported in " + e.NamingServiceName + " naming service"
}
