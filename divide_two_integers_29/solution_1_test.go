package divide_two_integers_29

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 10 / 3, expect 3",
			args: args{
				dividend: 10,
				divisor:  3,
			},
			want: 3,
		},
		{
			name: "test 7 / -3, expect -2",
			args: args{
				dividend: 7,
				divisor:  -3,
			},
			want: -2,
		},
		{
			name: "test -2147483648 / -1, expect 2147483647",
			args: args{
				dividend: -2147483648,
				divisor:  -1,
			},
			want: 2147483647,
		},
		{
			name: "test 2147483647 / 1, expect 2147483647",
			args: args{
				dividend: 2147483647,
				divisor:  1,
			},
			want: 2147483647,
		},
		{
			name: "test 2147483647 / 2, expect 1073741823",
			args: args{
				dividend: 2147483647,
				divisor:  2,
			},
			want: 1073741823,
		},
		{
			name: "test 2147483647 / 3, expect 715827882",
			args: args{
				dividend: 2147483647,
				divisor:  3,
			},
			want: 715827882,
		},
		{
			name: "test 2147483647 / 4, expect 536870911",
			args: args{
				dividend: 2147483647,
				divisor:  4,
			},
			want: 536870911,
		},
		{
			name: "test 2147483647 / 5, expect 429496729",
			args: args{
				dividend: 2147483647,
				divisor:  5,
			},
			want: 429496729,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, divide(tt.args.dividend, tt.args.divisor))
		})
	}
}
