package find_all_numbers_disappeared_in_an_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDisappearedNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{4, 3, 2, 7, 8, 2, 3, 1}}, []int{5, 6}},
		{"", args{[]int{1, 1}}, []int{2}},
		{"", args{[]int{1, 1, 2, 2}}, []int{3, 4}},
		{"", args{[]int{1, 1, 2, 2, 3, 3}}, []int{4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDisappearedNumbers(tt.args.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
