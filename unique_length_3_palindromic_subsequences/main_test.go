package unique_length_3_palindromic_subsequences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countPalindromicSubsequence(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"aabca"}, 3},
		{"", args{"adc"}, 0},
		{"", args{"bbcbaba"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPalindromicSubsequence(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
