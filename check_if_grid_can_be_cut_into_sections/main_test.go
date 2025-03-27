package check_if_grid_can_be_cut_into_sections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkValidCuts(t *testing.T) {
	tests := []struct {
		name       string
		n          int
		rectangles [][]int
		want       bool
	}{
		{"", 5, [][]int{{1, 0, 5, 2}, {0, 2, 2, 4}, {3, 2, 5, 3}, {0, 4, 4, 5}}, true},
		{"", 4, [][]int{{0, 0, 1, 1}, {2, 0, 3, 4}, {0, 2, 2, 3}, {3, 0, 4, 3}}, true},
		{"", 4, [][]int{{0, 2, 2, 4}, {1, 0, 3, 2}, {2, 2, 3, 4}, {3, 0, 4, 2}, {3, 2, 4, 4}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, checkValidCuts(tt.n, tt.rectangles))
		})
	}
}
