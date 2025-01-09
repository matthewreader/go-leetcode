package best_sightseeing_pair_1014

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxScoreSightseeingPair(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 8,1,5,2,6, expect 11",
			args: args{
				values: []int{8, 1, 5, 2, 6},
			},
			want: 11,
		},
		{
			name: "test 1,2, expect 2",
			args: args{
				values: []int{1, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxScoreSightseeingPair(tt.args.values))
		})
	}
}
