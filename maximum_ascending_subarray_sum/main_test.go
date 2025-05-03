package maximum_ascending_subarray_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxAscendingSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{10, 20, 30, 5, 10, 50}}, 65},
		{"", args{[]int{10, 20, 30, 40, 50}}, 150},
		{"", args{[]int{12, 17, 15, 13, 10, 11, 12}}, 33},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxAscendingSum(tt.args.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
