package resolution

import (
	"testing"
)

func TestCnsWithDefaultProvider(t *testing.T) {
	_, err := NewCnsWithDefaultProvider()
	if err != nil {
		t.Error(err)
	}
}

func TestCnsRecord(t *testing.T) {
	testDomain := "brad.crypto"
	expectedRecord := "0x8aaD44321A86b170879d7A244c1e8d360c99DdA8"
	cns, _ := NewCnsWithDefaultProvider()
	record, err := cns.Record(testDomain, "crypto.ETH.address")
	if err != nil {
		t.Error(err)
	}
	if record != expectedRecord {
		t.Errorf("Record does not match with expected: %v (expected %v)\n", record, expectedRecord)
	}
}
