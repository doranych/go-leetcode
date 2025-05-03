package count_prefix_and_suffix_pairs_i

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countPrefixSuffixPairs(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]string{"a", "aba", "ababa", "aa"}}, 4},
		{"", args{[]string{"pa", "papa", "ma", "mama"}}, 2},
		{"", args{[]string{"abab", "ab"}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPrefixSuffixPairs(tt.args.words)
			assert.Equal(t, tt.want, got)
		})
	}
}
