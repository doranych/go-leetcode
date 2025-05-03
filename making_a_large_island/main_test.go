package making_a_large_island

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestIsland(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{grid: [][]int{{1}}}, want: 1},
		{name: "", args: args{grid: [][]int{{1, 0}, {0, 1}}}, want: 3},
		{name: "", args: args{grid: [][]int{{1, 1}, {1, 1}}}, want: 4},
		{name: "", args: args{grid: [][]int{{0, 0}, {0, 0}}}, want: 1},
		{name: "", args: args{grid: [][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 1}}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestIsland(tt.args.grid)
			assert.Equal(t, tt.want, got)
		})
	}
}
