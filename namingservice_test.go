package resolution

import (
	"github.com/stretchr/testify/assert"
	"github.com/unstoppabledomains/resolution-go/namingservice"
	"testing"
)

func TestEnforceImplementInterface(t *testing.T) {
	t.Parallel()
	zns := NewZnsWithDefaultProvider()
	cns, _ := NewCnsWithDefaultBackend()
	assert.Implements(t, (*NamingService)(nil), zns)
	assert.Implements(t, (*NamingService)(nil), cns)
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
