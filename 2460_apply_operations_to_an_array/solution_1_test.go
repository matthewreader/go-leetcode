package _460_apply_operations_to_an_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_applyOperations(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test [1,2,2,1,1,0], expect [1,4,2,0,0,0]",
			args: args{
				values: []int{1, 2, 2, 1, 1, 0},
			},
			want: []int{1, 4, 2, 0, 0, 0},
		},
		{
			name: "test [0,1], expect [1,0]",
			args: args{
				values: []int{0, 1},
			},
			want: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, applyOperations(tt.args.values))
		})
	}
}
