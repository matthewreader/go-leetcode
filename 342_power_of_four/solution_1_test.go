package _42_power_of_four

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfFour(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected bool
	}{
		{
			name:     "Power of four - 16",
			n:        16,
			expected: true,
		},
		{
			name:     "Not a power of four - 15",
			n:        15,
			expected: false,
		},
		{
			name:     "Power of four - 64",
			n:        64,
			expected: true,
		},
		{
			name:     "Negative number",
			n:        -4,
			expected: false,
		},
		{
			name:     "Testcase 1061",
			n:        4194304,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, isPowerOfFour(tt.n))
		})
	}
}
