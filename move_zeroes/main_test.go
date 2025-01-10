package move_zeroes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{0, 1, 0, 3, 12}}, []int{1, 3, 12, 0, 0}},
		{"", args{[]int{0}}, []int{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			require.Equal(t, tt.want, tt.args.nums)
		})
	}
}
