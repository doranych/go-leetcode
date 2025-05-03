package max_sum_of_a_pair_with_equal_sum_of_digits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{18, 43, 36, 13, 7}}, 54},
		{"", args{[]int{10, 12, 19, 14}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumSum(tt.args.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
