package map_of_highest_peak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_highestPeak(t *testing.T) {
	type args struct {
		isWater [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"", args{[][]int{{0, 1}, {0, 0}}}, [][]int{{1, 0}, {2, 1}}},
		{"", args{[][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}}}, [][]int{{1, 1, 0}, {0, 1, 1}, {1, 2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := highestPeak(tt.args.isWater)
			assert.Equal(t, tt.want, got)
		})
	}
}
