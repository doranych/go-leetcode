package count_good_triplets_in_an_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_goodTriplets(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  int64
	}{
		{"", []int{2, 0, 1, 3}, []int{0, 1, 2, 3}, 1},
		{"", []int{4, 0, 1, 3, 2}, []int{4, 1, 0, 2, 3}, 4},
		{"", []int{0, 1, 2}, []int{2, 1, 0}, 0},
		{"", []int{0, 1, 2, 3}, []int{0, 1, 2, 3}, 4},
		{"", []int{0, 1, 2}, []int{0, 1, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := goodTriplets(tt.nums1, tt.nums2)
			assert.Equal(t, tt.want, got)
		})
	}
}
