package count_number_of_bad_pairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countBadPairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"", args{[]int{1, 2, 3, 4, 5}}, 0},
		{"", args{[]int{4, 1, 3, 3}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countBadPairs(tt.args.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
