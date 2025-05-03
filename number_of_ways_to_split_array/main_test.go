package number_of_ways_to_split_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_waysToSplitArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{10, 4, -8, 7}}, 2},
		{"", args{[]int{2, 3, 1, 0}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := waysToSplitArray(tt.args.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
