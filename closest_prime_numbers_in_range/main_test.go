package closest_prime_numbers_in_range

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_closestPrimes(t *testing.T) {
	tests := []struct {
		name  string
		left  int
		right int
		want  []int
	}{
		{"", 10, 20, []int{11, 13}},
		{"", 3, 10, []int{3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := closestPrimes(tt.left, tt.right)
			assert.Equal(t, tt.want, got)
		})
	}
}
