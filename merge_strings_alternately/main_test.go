package merge_strings_alternately

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeAlternately(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  string
	}{
		{"", "abc", "pqr", "apbqcr"},
		{"", "ab", "pqrs", "apbqrs"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeAlternately(tt.word1, tt.word2)
			assert.Equal(t, tt.want, got)
		})
	}
}
