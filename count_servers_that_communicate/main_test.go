package count_servers_that_communicate

import "testing"

func Test_countServers(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[][]int{{1, 0}, {0, 1}}}, 0},
		{"", args{[][]int{{1, 0}, {1, 1}}}, 3},
		{"", args{[][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countServers(tt.args.grid); got != tt.want {
				t.Errorf("countServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
