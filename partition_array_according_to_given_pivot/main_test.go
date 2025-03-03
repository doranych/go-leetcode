package partition_array_according_to_given_pivot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pivotArray(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		pivot int
		want  []int
	}{
		{"", []int{9, 12, 5, 10, 14, 3, 10}, 10, []int{9, 5, 3, 10, 10, 12, 14}},
		{"", []int{-3, 4, 3, 2}, 2, []int{-3, 2, 4, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pivotArray(tt.nums, tt.pivot)
			assert.Equal(t, tt.want, got)
		})
	}
}
