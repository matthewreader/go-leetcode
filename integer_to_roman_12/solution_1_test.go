package integer_to_roman_12

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_intToRoman(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want string
	}{
		{
			name: "58 = LVIII",
			arg:  58,
			want: "LVIII",
		},
		{
			name: "1994 = MCMXCIV",
			arg:  1994,
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, intToRoman(tt.arg))
		})
	}
}
