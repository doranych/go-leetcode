package minimum_number_of_operations_to_make_elements_in_array_distinct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"", []int{1, 2, 3, 4, 2, 3, 3, 5, 7}, 2},
		{"", []int{4, 5, 6, 4, 4}, 2},
		{"", []int{6, 7, 8, 9}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumOperations(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
