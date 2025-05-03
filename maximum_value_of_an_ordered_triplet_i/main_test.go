package maximum_value_of_an_ordered_triplet_i

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumTripletValue(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int64
	}{
		{"", []int{12, 6, 1, 2, 7}, 77},
		{"", []int{1, 10, 3, 4, 19}, 133},
		{"", []int{1, 2, 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumTripletValue(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
