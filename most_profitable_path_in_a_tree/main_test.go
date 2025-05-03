package most_profitable_path_in_a_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mostProfitablePath(t *testing.T) {
	tests := []struct {
		name   string
		edges  [][]int
		bob    int
		amount []int
		want   int
	}{
		{"", [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}, 3, []int{-2, 4, 2, -4, 6}, 6},
		{"", [][]int{{0, 1}}, 1, []int{-7280, 2350}, -7280},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mostProfitablePath(tt.edges, tt.bob, tt.amount)
			assert.Equal(t, tt.want, got)
		})
	}
}
