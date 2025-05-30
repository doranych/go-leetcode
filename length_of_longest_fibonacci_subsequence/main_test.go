package length_of_longest_fibonacci_subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lenLongestFibSubseq(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"", []int{1, 2, 3, 4, 5, 6, 7, 8}, 5},
		{"", []int{1, 3, 7, 11, 12, 14, 18}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lenLongestFibSubseq(tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
