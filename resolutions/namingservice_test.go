package resolution

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unstoppabledomains/resolution-go/v4/resolutions/namingservice"
)

func TestEnforceImplementInterface(t *testing.T) {
	t.Parallel()
	assert.Implements(t, (*NamingService)(nil), &Zns{provider: nil})
	assert.Implements(t, (*NamingService)(nil), &Uns{
		l1Service: UnsService{},
		l2Service: UnsService{},
	})
}

func TestDetectNamingServiceType(t *testing.T) {
	t.Parallel()
	var serviceType string
	serviceType, err := DetectNamingService("test.zil")
	assert.Nil(t, err)
	assert.Equal(t, namingservice.ZNS, serviceType)

	serviceType, err = DetectNamingService("test.crypto")
	assert.Nil(t, err)
	assert.Equal(t, namingservice.UNS, serviceType)

	serviceType, err = DetectNamingService("test.asdasdas")
	assert.Nil(t, err)
	assert.Equal(t, namingservice.UNS, serviceType)

	serviceType, err = DetectNamingService("test.wallet")
	assert.Nil(t, err)
	assert.Equal(t, namingservice.UNS, serviceType)
}

func TestDetectNamingServiceTypeInvalidDomain(t *testing.T) {
	t.Parallel()
	var expectedError *DomainNotSupportedError
	_, err := DetectNamingService("aaaazzsd..")
	assert.ErrorAs(t, err, &expectedError)

	_, err = DetectNamingService("aaaazzsd")
	assert.ErrorAs(t, err, &expectedError)
}
