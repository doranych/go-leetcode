package maximum_number_of_fish_in_a_grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMaxFish(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}}}, 7},
		{"", args{[][]int{{1, 2, 3, 4}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}}, 10},
		{"", args{[][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 1}}}, 1},
		{"", args{[][]int{{0, 8, 3, 1}}}, 12},
		{"", args{[][]int{{8, 6}, {2, 6}}}, 22},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxFish(tt.args.grid)
			assert.Equal(t, tt.want, got)
		})
	}
}
