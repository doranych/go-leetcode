package largest_local_values_in_a_matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestLocal(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want [][]int
	}{
		{"", [][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}},
			[][]int{{9, 9}, {8, 6}}},
		{"", [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 2, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}},
			[][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestLocal(tt.grid)
			assert.Equal(t, tt.want, got)
		})
	}
}
