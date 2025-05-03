package count_good_triplets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countGoodTriplets(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		a    int
		b    int
		c    int
		want int
	}{
		{"", []int{3, 0, 1, 1, 9, 7}, 7, 2, 3, 4},
		{"", []int{1, 1, 2, 2, 3}, 0, 0, 1, 0},
		{"", []int{1, 2, 3}, 0, 0, 0, 0},
		{"", []int{1, 1, 1}, 10, 10, 10, 1},
		{"", []int{1, 2, 3}, 2, 2, 2, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countGoodTriplets(tt.arr, tt.a, tt.b, tt.c)
			assert.Equal(t, tt.want, got)
		})
	}
}
