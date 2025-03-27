package minimum_index_of_a_valid_split

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"", []int{1, 2, 2, 2}, 2},
		{"", []int{1, 2, 3, 4}, -1},
		{"", []int{3, 3, 3, 3, 7, 2, 2}, -1},
		{"", []int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minimumIndex(tt.nums))
		})
	}
}
