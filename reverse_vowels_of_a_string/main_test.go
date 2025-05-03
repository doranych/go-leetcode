package reverse_vowels_of_a_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"", "IceCreAm", "AceCreIm"},
		{"", "leetcode", "leotcede"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseVowels(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
