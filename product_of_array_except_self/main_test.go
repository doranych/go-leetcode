package product_of_array_except_self

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_productExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"", []int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
		{"", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := productExceptSelf(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
