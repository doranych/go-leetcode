package number_of_substrings_containing_all_three_characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numberOfSubstrings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"", "abc", 1},
		{"", "abcabc", 10},
		{"", "aaacb", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, numberOfSubstrings(tt.s))
		})
	}
}
