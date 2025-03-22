package minimum_cost_walk_in_weighted_graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumCost(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		n       int
		edges   [][]int
		queries [][]int
		want    []int
	}{
		{},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, minimumCost(tt.n, tt.edges, tt.queries))
	}
}
