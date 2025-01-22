package map_of_heighest_peak

import (
	"reflect"
	"testing"
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
			if got := highestPeak(tt.args.isWater); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("highestPeak() = %v, want %v", got, tt.want)
			}
		})
	}
}
