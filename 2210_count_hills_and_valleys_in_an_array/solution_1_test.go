package _210_count_hills_and_valleys_in_an_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountHillValley(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{2, 4, 1, 1, 6, 5},
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{6, 6, 5, 5, 4, 1},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countHillValley(tt.args.nums))
		})
	}
}
