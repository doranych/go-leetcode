package count_total_number_of_colored_cells

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_coloredCells(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int64
	}{
		{"", 1, 1},
		{"", 2, 5},
		{"", 3, 13},
		{"", 15, 421},
		{"", 3333, 22211113},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coloredCells(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
