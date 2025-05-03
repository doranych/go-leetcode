package container_with_most_water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"", args{[]int{1, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.args.height)
			assert.Equal(t, tt.want, got)
		})
	}
}
