package _717_maximum_score_substring_removal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximumGain(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		x        int
		y        int
		expected int
	}{
		{
			name:     "Example 1 - ba scores higher",
			s:        "cdbcbbaaabab",
			x:        4,
			y:        5,
			expected: 19,
		},
		{
			name:     "Example 2 - ab scores higher",
			s:        "aabbaaxybbaabb",
			x:        5,
			y:        4,
			expected: 20,
		},
		{
			name:     "Equal scores",
			s:        "abab",
			x:        3,
			y:        3,
			expected: 6,
		},
		{
			name:     "Only ab substrings",
			s:        "ababab",
			x:        10,
			y:        1,
			expected: 30,
		},
		{
			name:     "Only ba substrings",
			s:        "bababa",
			x:        1,
			y:        10,
			expected: 30,
		},
		{
			name:     "No valid substrings",
			s:        "cdefgh",
			x:        5,
			y:        5,
			expected: 0,
		},
		{
			name:     "Single character",
			s:        "a",
			x:        10,
			y:        10,
			expected: 0,
		},
		{
			name:     "Empty string",
			s:        "",
			x:        5,
			y:        5,
			expected: 0,
		},
		{
			name:     "Complex nested pattern",
			s:        "aabbbaaba",
			x:        3,
			y:        7,
			expected: 24, // ba=7, ba=7, ba=7, ab=3 = 24
		},
		{
			name:     "Large x value prioritized",
			s:        "abbaab",
			x:        100,
			y:        1,
			expected: 201, // ab=100, ab=100, ba=1 = 201
		},
		{
			name:     "Large y value prioritized",
			s:        "abbaab",
			x:        1,
			y:        100,
			expected: 201, // ba=100, ba=100, ab=1 = 201
		},
		{
			name:     "Overlapping opportunities",
			s:        "aabbaa",
			x:        2,
			y:        3,
			expected: 6, // ba=3, ab=2, ab=1 = 6
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximumGain(tt.s, tt.x, tt.y)
			assert.Equal(t, tt.expected, result,
				"maximumGainOptimized(%q, %d, %d) = %d, want %d",
				tt.s, tt.x, tt.y, result, tt.expected)
		})
	}
}
