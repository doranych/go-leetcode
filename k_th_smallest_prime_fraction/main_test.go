package k_th_smallest_prime_fraction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kthSmallestPrimeFraction(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		k    int
		want []int
	}{
		{"", []int{1, 2, 3, 5}, 3, []int{2, 5}},
		{"", []int{1, 7}, 1, []int{1, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kthSmallestPrimeFraction(tt.arr, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
