package _8_find_the_index_of_the_first_occurrence_in_a_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test hello, ll, expect 2",
			args: args{
				haystack: "hello",
				needle:   "ll",
			},
			want: 2,
		},
		{
			name: "test a, a, expect 0",
			args: args{
				haystack: "a",
				needle:   "a",
			},
			want: 0,
		},
		{
			name: "test a, b, expect -1",
			args: args{
				haystack: "a",
				needle:   "b",
			},
			want: -1,
		},
		{
			name: "test a, , expect 0",
			args: args{
				haystack: "a",
				needle:   "",
			},
			want: 0,
		},
		{
			name: "test , a, expect -1",
			args: args{
				haystack: "",
				needle:   "a",
			},
			want: -1,
		},
		{
			name: "test , , expect 0",
			args: args{
				haystack: "",
				needle:   "",
			},
			want: 0,
		},
		{
			name: "test abc, c, expect 2",
			args: args{
				haystack: "abc",
				needle:   "c",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, strStr(tt.args.haystack, tt.args.needle))
		})
	}
}
