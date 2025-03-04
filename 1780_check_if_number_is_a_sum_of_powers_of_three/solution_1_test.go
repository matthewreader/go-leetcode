package _780_check_if_number_is_a_sum_of_powers_of_three

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_checkPowersOfThree(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 12, expect true",
			args: args{
				n: 12,
			},
			want: true,
		},
		{
			name: "test 91, expect true",
			args: args{
				n: 91,
			},
			want: true,
		},
		{
			name: "test 11, expect false",
			args: args{
				n: 11,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, checkPowersOfThree(tt.args.n))
		})
	}
}
