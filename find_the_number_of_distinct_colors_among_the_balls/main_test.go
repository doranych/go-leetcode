package find_the_number_of_distinct_colors_among_the_balls

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_queryResults(t *testing.T) {
	type args struct {
		limit   int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{3, [][]int{{1, 2}, {2, 3}, {3, 4}}}, []int{1, 2, 3}},
		{"", args{4, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}}, []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := queryResults(tt.args.limit, tt.args.queries)
			assert.Equal(t, tt.want, got)
		})
	}
}
