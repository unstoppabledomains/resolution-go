package resolution

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSupportedKeys(t *testing.T) {
	t.Parallel()
	keys, err := NewSupportedKeys()
	assert.Nil(t, err)
	assert.Equal(t, "ETH", keys["crypto.ETH.address"].DeprecatedKeyName)
	assert.Equal(t, "BTC", keys["crypto.BTC.address"].DeprecatedKeyName)
}
