package tuple_with_same_product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tupleSameProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{2, 3, 4, 6}}, 8},
		{"", args{[]int{1, 2, 4, 5, 10}}, 16},
		{"", args{[]int{1, 3, 7, 21, 49, 147}}, 40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tupleSameProduct(tt.args.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
