package count_subarrays_with_score_less_than_k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countSubarrays(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int64
		want int64
	}{
		{"test1", []int{2, 1, 4, 3, 5}, 10, 6},
		{"test2", []int{1, 1, 1}, 5, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countSubarrays(tt.nums, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
