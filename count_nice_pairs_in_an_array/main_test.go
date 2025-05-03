package count_nice_pairs_in_an_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countNicePairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{42, 11, 1, 97}}, 2},
		{"", args{[]int{13, 10, 35, 24, 76}}, 4},
		{"", args{[]int{1, 1, 1, 1, 1}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countNicePairs(tt.args.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
