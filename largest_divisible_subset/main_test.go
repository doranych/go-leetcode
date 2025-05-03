package largest_divisible_subset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestDivisibleSubset(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"", []int{1, 2, 3}, []int{1, 3}},
		{"", []int{1, 2, 4, 8}, []int{1, 2, 4, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestDivisibleSubset(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
