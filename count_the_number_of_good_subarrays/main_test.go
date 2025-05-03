package count_the_number_of_good_subarrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countGood(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int64
	}{
		{"", []int{1, 1, 1, 1, 1}, 10, 1},
		{"", []int{3, 1, 4, 3, 2, 2, 4}, 2, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countGood(tt.nums, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
