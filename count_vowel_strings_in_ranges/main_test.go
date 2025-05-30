package count_vowel_strings_in_ranges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_vowelStrings(t *testing.T) {
	type args struct {
		words   []string
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]string{"aba", "bcb", "ece", "aa", "e"},
			[][]int{{0, 2}, {1, 4}, {1, 1}}}, []int{2, 3, 0}},
		{"", args{[]string{"a", "e", "i"},
			[][]int{{0, 2}, {0, 1}, {2, 2}}}, []int{3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := vowelStrings(tt.args.words, tt.args.queries)
			assert.Equal(t, tt.want, got)
		})
	}
}
