package _161_partition_array_according_to_given_pivot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pivotArray(t *testing.T) {
	type args struct {
		nums  []int
		pivot int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test [9,12,5,10,14,3,10] pivot=10, expect [9,5,3,10,10,12,14]",
			args: args{
				nums:  []int{9, 12, 5, 10, 14, 3, 10},
				pivot: 10,
			},
			want: []int{9, 5, 3, 10, 10, 12, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, pivotArray(tt.args.nums, tt.args.pivot))
		})
	}
}
