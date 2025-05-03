package grid_game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_gridGame(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"", args{[][]int{{2, 5, 4}, {1, 5, 1}}}, 4},
		{"", args{[][]int{{3, 3, 1}, {8, 5, 2}}}, 4},
		{"", args{[][]int{{1, 3, 1, 15}, {1, 3, 3, 1}}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gridGame(tt.args.grid)
			assert.Equal(t, tt.want, got)
		})
	}
}
