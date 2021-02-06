package resolution

import (
	"testing"
)

func TestNormalizeName(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"     ", ""},
		{"   beresnev.crypto   ", "beresnev.crypto"},
		{"BeResNev.crypto", "beresnev.crypto"},
		{"   BeResNev.crypto", "beresnev.crypto"},
	}

	for _, tt := range tests {
		result := NormalizeName(tt.input)
		if tt.output != result {
			t.Errorf("Failure: %v => %v (expected %v)\n", tt.input, result, tt.output)
		}
	}
}
