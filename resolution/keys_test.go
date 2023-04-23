package resolution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSupportedKeys(t *testing.T) {
	t.Parallel()
	keys, err := newSupportedKeys()
	assert.Nil(t, err)
	assert.Equal(t, "ETH", keys["crypto.ETH.address"].DeprecatedKeyName)
	assert.Equal(t, "BTC", keys["crypto.BTC.address"].DeprecatedKeyName)
	assert.Equal(t, "MDX_HRC20", keys["crypto.MDX.version.HRC20.address"].DeprecatedKeyName)
}

func TestReturnFirstNonEmptyReturnFirst(t *testing.T) {
	t.Parallel()
	records := map[string]string{"crypto.ETH.address": "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8", "crypto.BTC.address": "bc1q359khn0phg58xgezyqsuuaha28zkwx047c0c3y"}
	result := returnFirstNonEmpty(records, []string{"crypto.ETH.address", "crypto.BTC.address"})
	assert.Equal(t, "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8", result)
}

func TestReturnFirstNonEmptyReturnSecond(t *testing.T) {
	t.Parallel()
	records := map[string]string{"crypto.ETH.address": "", "crypto.BTC.address": "bc1q359khn0phg58xgezyqsuuaha28zkwx047c0c3y"}
	result := returnFirstNonEmpty(records, []string{"crypto.ETH.address", "crypto.BTC.address"})
	assert.Equal(t, "bc1q359khn0phg58xgezyqsuuaha28zkwx047c0c3y", result)
}

func TestReturnFirstNonEmptyReturnEmpty(t *testing.T) {
	t.Parallel()
	records := map[string]string{"crypto.ETH.address": "", "crypto.BTC.address": ""}
	result := returnFirstNonEmpty(records, []string{"crypto.ETH.address", "crypto.BTC.address"})
	assert.Equal(t, "", result)
}

func TestReturnFirstNonEmptyNoRecords(t *testing.T) {
	t.Parallel()
	records := map[string]string{}
	result := returnFirstNonEmpty(records, []string{"crypto.ETH.address", "crypto.BTC.address"})
	assert.Equal(t, "", result)
}

func TestBuildCryptoKey(t *testing.T) {
	t.Parallel()
	result, err := buildCryptoKey("ETH")
	assert.Nil(t, err)
	assert.Equal(t, "crypto.ETH.address", result)
}

func TestBuildCryptoKeyVersion(t *testing.T) {
	t.Parallel()
	result, err := buildCryptoKeyVersion("USDT", "ERC20")
	assert.Nil(t, err)
	assert.Equal(t, "crypto.USDT.version.ERC20.address", result)
}
