package minimum_operations_to_make_binary_array_elements_equal_to_one_i

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"", []int{0, 1, 1, 1, 0, 0}, 3},
		{"", []int{0, 1, 1, 1}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minOperations(tt.nums))
		})
	}
}
