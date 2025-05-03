package the_k_th_lexicographical_string_of_all_happy_strings_of_length_n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getHappyString(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{1, 3}, "c"},
		{"", args{1, 1}, "a"},
		{"", args{3, 9}, "cab"},
		{"", args{2, 7}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getHappyString(tt.args.n, tt.args.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
