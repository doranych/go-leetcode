package minimum_domino_rotations_for_equal_row

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minDominoRotations(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		B    []int
		want int
	}{
		{"", []int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2}, 2},
		{"", []int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDominoRotations(tt.A, tt.B)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
