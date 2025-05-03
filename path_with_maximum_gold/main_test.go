package path_with_maximum_gold

import "testing"

func Test_getMaximumGold(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"", [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}, 24},
		{"", [][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}}, 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumGold(tt.grid); got != tt.want {
				t.Errorf("getMaximumGold() = %v, want %v", got, tt.want)
			}
		})
	}
}
