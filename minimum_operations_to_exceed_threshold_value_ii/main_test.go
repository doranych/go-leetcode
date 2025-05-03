package minimum_operations_to_exceed_threshold_value_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minOperations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{2, 11, 10, 1, 3}, 10}, 2},
		{"", args{[]int{1, 1, 2, 4, 9}, 20}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minOperations(tt.args.nums, tt.args.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
