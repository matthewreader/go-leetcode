package _191_minimum_operations_to_make_binary_array_elements_equal_to_one_i

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test [0,1,1,1,0,0], expect 3",
			args: args{
				nums: []int{0, 1, 1, 1, 0, 0},
			},
			want: 3,
		},
		{
			name: "test [0,1,1,1], expect -1",
			args: args{
				nums: []int{0, 1, 1, 1},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minOperations(tt.args.nums))
		})
	}
}
