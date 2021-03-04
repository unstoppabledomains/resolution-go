package resolution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZnsNameHash(t *testing.T) {
	t.Parallel()
	namehash, err := ZnsNameHash("brad.zil")
	assert.Nil(t, err)
	assert.Equal(t, "0x5fc604da00f502da70bfbc618088c0ce468ec9d18d05540935ae4118e8f50787", namehash)
}

func TestZnsNameHashSubdomain(t *testing.T) {
	t.Parallel()
	namehash, err := ZnsNameHash("subdomain.brad.zil")
	assert.Nil(t, err)
	assert.Equal(t, "0x5ec61cc1b660ba2fbc2b8a02ee73e386b8c9e3a6c77d2acfe61d04ad95f3d495", namehash)
}

func TestZnsNameHashRoot(t *testing.T) {
	t.Parallel()
	namehash, err := ZnsNameHash("zil")
	assert.Nil(t, err)
	assert.Equal(t, "0x9915d0456b878862e822e2361da37232f626a2e47505c8795134a95d36138ed3", namehash)
}
