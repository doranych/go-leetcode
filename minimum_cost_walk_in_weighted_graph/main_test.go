package minimum_cost_walk_in_weighted_graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumCost(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		edges   [][]int
		queries [][]int
		want    []int
	}{
		{"", 5, [][]int{{0, 1, 7}, {1, 3, 7}, {1, 2, 1}},
			[][]int{{0, 3}, {3, 4}}, []int{1, -1}},
		{"", 3, [][]int{{0, 2, 7}, {0, 1, 15}, {1, 2, 6}, {1, 2, 1}},
			[][]int{{1, 2}}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minimumCost(tt.n, tt.edges, tt.queries))
		})
	}
}
