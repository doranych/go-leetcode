package maximum_count_of_positive_integer_and_negative_integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumCount(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"", []int{1, 2, 3, 4, 5}, 5},
		{"", []int{-5, -4, -3, -2, -1, 0, 0, 1, 2, 3, 4, 5}, 5},
		{"", []int{-2, -1, -1, 1, 2, 3}, 3},
		{"", []int{-3, -2, -1, 0, 0, 1, 2}, 3},
		{"", []int{5, 20, 66, 1314}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maximumCount(tt.nums))
		})
	}
}
