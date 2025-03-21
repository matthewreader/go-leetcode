package container_with_most_water_11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			want: 49,
		},
		{
			name: "test 2",
			args: args{
				height: []int{1, 1},
			},
			want: 1,
		},
		{
			name: "test 3",
			args: args{
				height: []int{4, 3, 2, 1, 4},
			},
			want: 16,
		},
		{
			name: "test 4",
			args: args{
				height: []int{1, 2, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxArea(tt.args.height))
		})
	}
}
