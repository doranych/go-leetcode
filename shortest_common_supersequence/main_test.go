package shortest_common_supersequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shortestCommonSupersequence(t *testing.T) {
	tests := []struct {
		name string
		str1 string
		str2 string
		want string
	}{
		{"", "abac", "cab", "cabac"},
		{"", "abc", "ac", "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shortestCommonSupersequence(tt.str1, tt.str2)
			assert.Equal(t, tt.want, got)
		})
	}
}
