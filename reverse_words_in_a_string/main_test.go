package reverse_words_in_a_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"", "the sky is blue", "blue is sky the"},
		{"", "  hello world!  ", "world! hello"},
		{"", "a good   example", "example good a"},
		{"", "  Bob    Loves  Alice   ", "Alice Loves Bob"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseWords(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
