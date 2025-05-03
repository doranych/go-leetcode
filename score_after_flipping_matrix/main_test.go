package score_after_flipping_matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_matrixScore(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"", [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}, 39},
		{"", [][]int{{0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := matrixScore(tt.grid)
			assert.Equal(t, tt.want, got)
		})
	}
}
