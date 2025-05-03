package sum_of_all_subset_xor_totals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsetXORSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"", []int{1, 3}, 6},
		{"", []int{5, 1, 6}, 28},
		{"", []int{3, 4, 5, 6, 7, 8}, 480},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsetXORSum(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
