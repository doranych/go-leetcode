package house_robber_iv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minCapability(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"", []int{2, 3, 5, 9}, 2, 9},
		{"", []int{2, 7, 9, 3, 1}, 2, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minCapability(tt.nums, tt.k))
		})
	}
}
