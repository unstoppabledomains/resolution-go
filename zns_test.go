package resolution

import (
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/stretchr/testify/assert"
	"testing"
)

var zns = NewZnsWithDefaultProvider()

func TestNewZns(t *testing.T) {
	t.Parallel()
	znsProvider := provider.NewProvider("https://api.zilliqa.com")
	zns := NewZns(znsProvider)
	assert.IsType(t, &Zns{Provider: nil}, zns)
}

func TestNewZnsWithDefaultProvider(t *testing.T) {
	t.Parallel()
	zns := NewZnsWithDefaultProvider()
	assert.IsType(t, &Zns{Provider: nil}, zns)
}

func TestZnsState(t *testing.T) {
	t.Parallel()
	zns.State("brad.zil", []string{})
	// todo finish the test
}
