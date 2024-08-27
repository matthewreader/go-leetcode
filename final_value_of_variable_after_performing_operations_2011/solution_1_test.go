package final_value_of_variable_after_performing_operations_2011

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_finalValueAfterOperations(t *testing.T) {
	type args struct {
		operations []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test ['--X','X++','X++'], expect 1",
			args: args{
				operations: []string{"--X", "X++", "X++"},
			},
			want: 1,
		},
		{
			name: "test ['++X','++X','X++'], expect 3",
			args: args{
				operations: []string{"++X", "++X", "X++"},
			},
			want: 3,
		},
		{
			name: "test ['X++','++X','--X','X--'], expect 0",
			args: args{
				operations: []string{"X++", "++X", "--X", "X--"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, finalValueAfterOperations(tt.args.operations))
		})
	}
}
