package special_array_i

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isArraySpecial(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"", []int{1, 2, 3, 4}, true},
		{"", []int{1, 2, 3}, true},
		{"", []int{4, 3, 1, 6}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isArraySpecial(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
