package _843_count_symmetric_integers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCount(t *testing.T) {
	type args struct {
		low  int
		high int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "low = 1, high = 100 yields 9",
			args: args{
				low:  1,
				high: 100,
			},
			want: 9,
		},
		{
			name: "low = 1200, high = 1230 yields 4",
			args: args{
				low:  1200,
				high: 1230,
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countSymmetricIntegers(tt.args.low, tt.args.high))
		})
	}
}
