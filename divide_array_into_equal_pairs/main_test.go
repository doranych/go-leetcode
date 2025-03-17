package divide_array_into_equal_pairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_divideArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"", []int{1, 2, 3, 4}, false},
		{"", []int{1, 1, 1, 1}, true},
		{"", []int{1, 2, 2, 1}, true},
		{"", []int{3, 2, 3, 2, 2, 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, divideArray(tt.nums))
		})
	}
}
