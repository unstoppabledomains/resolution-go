package resolution

import (
	"github.com/stretchr/testify/assert"
	"github.com/unstoppabledomains/resolution-go/namingservice"
	"testing"
)

func TestEnforceImplementInterface(t *testing.T) {
	t.Parallel()
	assert.Implements(t, (*NamingService)(nil), &Zns{provider: nil})
	assert.Implements(t, (*NamingService)(nil), &Cns{
		proxyReader:     nil,
		supportedKeys:   nil,
		contractBackend: nil,
	})
}

func TestDetectNamingServiceType(t *testing.T) {
	t.Parallel()
	var serviceType string
	serviceType, err := DetectNamingServiceType("test.zil")
	assert.Nil(t, err)
	assert.Equal(t, namingservice.ZNS, serviceType)

	serviceType, err = DetectNamingServiceType("test.crypto")
	assert.Nil(t, err)
	assert.Equal(t, namingservice.CNS, serviceType)
}

func TestDetectNamingServiceTypeInvalidDomain(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotSupported
	_, err := DetectNamingServiceType("aaaazzsd")
	assert.ErrorAs(t, err, &expectedError)
}

func TestDetectNamingServiceTypeUnsupportedDomain(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotSupported
	_, err := DetectNamingServiceType("google.com")
	assert.ErrorAs(t, err, &expectedError)
}
