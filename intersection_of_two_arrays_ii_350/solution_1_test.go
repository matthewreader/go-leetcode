package intersection_of_two_arrays_ii_350

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test [1,2,2,1] and [2,2] intersect",
			args: args{
				nums1: []int{1, 2, 2, 1},
				nums2: []int{2, 2},
			},
			want: []int{2, 2},
		},
		{
			name: "test [4,9,5] and [9,4,9,8,4] intersect",
			args: args{
				nums1: []int{4, 9, 5},
				nums2: []int{9, 4, 9, 8, 4},
			},
			want: []int{4, 9},
		},
		{
			name: "test [1, 1, 1, 1] and [1, 1, 1, 2] intersect",
			args: args{
				nums1: []int{1, 1, 1, 1},
				nums2: []int{1, 1, 1, 2},
			},
			want: []int{1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, intersect(tt.args.nums1, tt.args.nums2))
		})
	}
}
