package maximum_number_of_points_from_grid_queries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxPoints(t *testing.T) {
	tests := []struct {
		name    string
		grid    [][]int
		queries []int
		want    []int
	}{
		{"", [][]int{{1, 2, 3}, {2, 5, 7}, {3, 5, 1}}, []int{5, 6, 2}, []int{5, 8, 1}},
		{"", [][]int{{5, 2, 1}, {1, 1, 2}}, []int{3}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxPoints(tt.grid, tt.queries))
		})
	}
}
