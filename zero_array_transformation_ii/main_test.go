package zero_array_transformation_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minZeroArray(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		queries [][]int
		want    int
	}{
		{"", []int{2, 0, 2}, [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}}, 2},
		{"", []int{4, 3, 2, 1}, [][]int{{1, 3, 2}, {0, 2, 1}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minZeroArray(tt.nums, tt.queries))
		})
	}
}
