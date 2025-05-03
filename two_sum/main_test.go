package two_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"", []int{3, 2, 4}, 6, []int{1, 2}},
		{"", []int{3, 3}, 6, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}
