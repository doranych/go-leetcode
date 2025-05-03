package divide_nodes_into_the_maximum_number_of_groups

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_magnificentSets(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{6, [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}}}, 4},
		{"", args{3, [][]int{{1, 2}, {2, 3}, {3, 1}}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := magnificentSets(tt.args.n, tt.args.edges)
			assert.Equal(t, tt.want, got)
		})
	}
}
