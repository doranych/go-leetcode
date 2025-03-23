package number_of_ways_to_arrive_at_destination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countPaths(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		roads [][]int
		want  int
	}{
		{"", 7, [][]int{
			{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3},
			{6, 3, 3}, {3, 5, 1}, {6, 5, 1}, {2, 5, 1},
			{0, 4, 5}, {4, 6, 2},
		}, 4},
		{"", 2, [][]int{{1, 0, 10}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countPaths(tt.n, tt.roads))
		})
	}
}
