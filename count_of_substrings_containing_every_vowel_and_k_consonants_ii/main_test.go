package count_of_substrings_containing_every_vowel_and_k_consonants_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countOfSubstrings(t *testing.T) {
	tests := []struct {
		name string
		word string
		k    int
		want int64
	}{
		{"", "aeioqq", 1, 0},
		{"", "aeiou", 0, 1},
		{"", "ieaouqqieaouqq", 1, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, countOfSubstrings(tt.word, tt.k), tt.want)
		})
	}
}
