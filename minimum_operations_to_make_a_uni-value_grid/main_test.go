package minimum_operations_to_make_a_uni_value_grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minOperations(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		x    int
		want int
	}{
		{"", [][]int{{2, 4}, {6, 8}}, 2, 4},
		{"", [][]int{{1, 5}, {2, 3}}, 1, 5},
		{"", [][]int{{1, 2}, {3, 4}}, 2, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minOperations(tt.grid, tt.x))
		})
	}
}
