package longest_nice_subarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestNiceSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"", []int{1, 3, 8, 48, 10}, 3},
		{"", []int{3, 1, 5, 11, 13}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, longestNiceSubarray(tt.nums), tt.want)
		})
	}
}
