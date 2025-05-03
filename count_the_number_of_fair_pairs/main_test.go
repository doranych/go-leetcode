package count_the_number_of_fair_pairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countFairPairs(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		lower int
		upper int
		want  int64
	}{
		{"test1", []int{0, 1, 7, 4, 4, 5}, 3, 6, 6},
		{"test1", []int{1, 7, 9, 2, 5}, 11, 11, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countFairPairs(tt.nums, tt.lower, tt.upper)
			assert.Equal(t, tt.want, got)
		})
	}
}
