package relative_ranks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findRelativeRanks(t *testing.T) {
	tests := []struct {
		name  string
		score []int
		want  []string
	}{
		{"", []int{5, 4, 3, 2, 1}, []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}},
		{"", []int{10, 3, 8, 9, 4}, []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}},
		{"", []int{1, 2, 3, 4, 5}, []string{"5", "4", "Bronze Medal", "Silver Medal", "Gold Medal"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findRelativeRanks(tt.score)
			assert.Equal(t, tt.want, got)
		})
	}
}
