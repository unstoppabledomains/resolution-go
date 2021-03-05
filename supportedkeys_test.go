package resolution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSupportedKeys(t *testing.T) {
	t.Parallel()
	keys, err := NewSupportedKeys()
	assert.Nil(t, err)
	assert.Equal(t, "ETH", keys["crypto.ETH.address"].DeprecatedKeyName)
	assert.Equal(t, "BTC", keys["crypto.BTC.address"].DeprecatedKeyName)
}
