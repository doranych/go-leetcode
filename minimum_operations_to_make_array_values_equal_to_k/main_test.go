package minimum_operations_to_make_array_values_equal_to_k

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"test1", []int{5, 2, 5, 4, 5}, 2, 2},
		{"test1", []int{2, 1, 2}, 2, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minOperations(tt.nums, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
