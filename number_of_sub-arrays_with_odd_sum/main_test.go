package number_of_sub_arrays_with_odd_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numOfSubarrays(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"", []int{1, 3, 5}, 4},
		{"", []int{2, 4, 6}, 0},
		{"", []int{1, 2, 3, 4, 5, 6, 7}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numOfSubarrays(tt.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
