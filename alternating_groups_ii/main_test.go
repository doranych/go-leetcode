package alternating_groups_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numberOfAlternatingGroups(t *testing.T) {
	tests := []struct {
		name   string
		colors []int
		k      int
		want   int
	}{
		{"", []int{0, 1, 0, 1, 0}, 3, 3},
		{"", []int{0, 1, 0, 1, 0}, 4, 2},
		{"", []int{1, 1, 0, 1}, 4, 0},
		{"", []int{0, 1, 0, 0, 1, 0, 1}, 6, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, numberOfAlternatingGroups(tt.colors, tt.k))
		})
	}
}
