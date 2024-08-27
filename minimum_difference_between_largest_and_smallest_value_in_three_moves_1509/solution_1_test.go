package minimum_difference_between_largest_and_smallest_value_in_three_moves_1509

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minDifference(t *testing.T) {
	tests := []struct {
		name string
		arg  []int
		want int
	}{
		{
			name: "test 1",
			arg:  []int{5, 3, 2, 4},
			want: 0,
		},
		{
			name: "test 2",
			arg:  []int{1, 5, 0, 10, 14},
			want: 1,
		},
		{
			name: "test 3",
			arg:  []int{6, 6, 0, 1, 1, 4, 6},
			want: 2,
		},
		{
			name: "test 4",
			arg:  []int{1, 5, 6, 14, 15},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minDifference(tt.arg))
		})
	}
}
