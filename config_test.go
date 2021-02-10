package resolution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSupportedKeysConfig(t *testing.T) {
	t.Parallel()
	config, err := NewSupportedKeysConfig()
	assert.Nil(t, err)
	assert.Equal(t, "1.1.0", config.Get("version"))
	assert.Equal(t, "BTC", config.Get("keys.crypto.BTC.address.deprecatedKeyName"))
}
