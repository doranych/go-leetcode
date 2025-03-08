package minimum_recolors_to_get_k_consecutive_black_blocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumRecolors(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{"", "WWBBBWBWWBWWBBBWWWWWBWBBWBWBWBWBWBWBBBWBW", 19, 8},
		{"", "WBBBWWWWBBBWBWWBWWWWBWBBWBWWWBWWWWWWBWWWBBWB", 34, 21},
		{"", "WWBBBWWBBWWBBBBWWWWBWBBWWWBWWBBBBWBWWWWWBBBWBWBBBWBBWWWBWWWWBWWWBWBWWWWBWWBB", 6, 1},
		{"", "WBWBWBBWWWWBWBBBWBWBB", 19, 9},
		{"", "WBBBW", 4, 1},
		{"", "WWBWBWWBBWBWBWWBWWBWWBBBBWWBBWWWWBBWBWWWBBBBBWBWBWWWWBBWBBBBBWWBBBWBWWWBBBWWWWBBWWBWBB", 40, 17},
		{"", "B", 1, 0},
		{"", "BBWWWBBWBBBWWWWBBBBWWWWBBWWWBWBWWBBBWWBWWWB", 2, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minimumRecolors(tt.s, tt.k))
		})
	}
}
