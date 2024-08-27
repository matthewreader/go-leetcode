package remove_element_27

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test (3, 2, 2, 3) remove 3, expect 2",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			name: "test (0, 1, 2, 2, 3, 0, 4, 2) remove 2, expect 5",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeElement(tt.args.nums, tt.args.val))
		})
	}
}
