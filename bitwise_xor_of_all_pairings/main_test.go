package bitwise_xor_of_all_pairings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_xorAllNums(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{2}, []int{4}}, 6},
		{"", args{[]int{27, 6, 15}, []int{8, 23, 15, 16, 25, 5}}, 28},
		{"", args{[]int{10, 0, 15, 25, 29, 3, 2}, []int{12}}, 12},
		{"", args{
			[]int{947, 662, 559, 0, 967, 251, 90, 234, 864, 402, 101, 889, 146, 358, 940, 264, 411, 962, 448, 369, 485, 925, 558, 347, 53, 245, 598},
			[]int{440, 506, 176, 480, 787, 124, 989, 649, 750, 508, 494}}, 670},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := xorAllNums(tt.args.nums1, tt.args.nums2)
			assert.Equal(t, tt.want, got)
		})
	}
}
