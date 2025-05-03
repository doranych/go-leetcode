package longest_palindromic_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"", "babad", "bab"},
		{"", "cbbd", "bb"},
		{"", "a", "a"},
		{"", "ac", "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
