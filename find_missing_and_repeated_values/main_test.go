package find_missing_and_repeated_values

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMissingAndRepeatedValues(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want []int
	}{
		{"", [][]int{{1, 2, 3}, {3, 5, 6}, {8, 7, 9}}, []int{3, 4}},
		{"", [][]int{{1, 1}, {3, 4}}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMissingAndRepeatedValues(tt.grid)
			assert.Equal(t, tt.want, got)
		})
	}
}
