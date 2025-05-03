package minimum_cost_to_make_at_least_one_valid_path_in_a_grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minCost(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{"", args{[][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}}}, 3},
		{"", args{[][]int{{1, 1, 3}, {3, 2, 2}, {1, 1, 4}}}, 0},
		{"", args{[][]int{{1, 2}, {4, 3}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCost(tt.args.grid)
			assert.Equal(t, tt.want, got)
		})
	}
}
