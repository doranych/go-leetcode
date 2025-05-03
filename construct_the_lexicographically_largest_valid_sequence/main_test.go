package construct_the_lexicographically_largest_valid_sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_constructDistancedSequence(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{1}, []int{1}},
		{"", args{2}, []int{2, 1, 2}},
		{"", args{3}, []int{3, 1, 2, 3, 2}},
		{"", args{4}, []int{4, 2, 3, 2, 4, 3, 1}},
		{"", args{5}, []int{5, 3, 1, 4, 3, 5, 2, 4, 2}},
		{"", args{10}, []int{10, 8, 6, 9, 3, 1, 7, 3, 6, 8, 10, 5, 9, 7, 4, 2, 5, 2, 4}},
		{"", args{20}, []int{20, 18, 19, 15, 13, 17, 10, 16, 7, 5, 3, 14, 12, 3, 5, 7, 10, 13, 15, 18,
			20, 19, 17, 16, 12, 14, 11, 9, 4, 6, 8, 2, 4, 2, 1, 6, 9, 11, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := constructDistancedSequence(tt.args.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
