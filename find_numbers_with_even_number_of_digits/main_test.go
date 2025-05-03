package find_numbers_with_even_number_of_digits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findNumbers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"", []int{12, 345, 2, 6, 7896}, 2},
		{"", []int{555, 901, 482, 1771}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findNumbers(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
