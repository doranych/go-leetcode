package count_largest_group

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countLargestGroup(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"", 13, 4},
		{"", 2, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countLargestGroup(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
