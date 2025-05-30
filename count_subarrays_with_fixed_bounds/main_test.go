package count_subarrays_with_fixed_bounds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countSubarrays(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		minK int
		maxK int
		want int64
	}{
		{"", []int{1, 3, 5, 2, 7, 5}, 1, 5, 2},
		{"", []int{1, 1, 1, 1}, 1, 1, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countSubarrays(tt.nums, tt.minK, tt.maxK)
			assert.Equal(t, tt.want, got)
		})
	}
}
