package apply_operations_to_an_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_applyOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"", []int{1, 2, 2, 1, 1, 0}, []int{1, 4, 2, 0, 0, 0}},
		{"", []int{0, 1}, []int{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := applyOperations(tt.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
