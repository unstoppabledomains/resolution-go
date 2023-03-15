package udclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUdClientDial(t *testing.T) {
	t.Parallel()

	client, err := Dial("test")
	assert.Nil(t, err)

	assert.NotNil(t, client)
	assert.NotNil(t, client.L1ContractBackend)
	assert.NotNil(t, client.L2ContractBackend)
}
