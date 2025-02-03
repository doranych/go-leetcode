package longest_strictly_increasing_or_strictly_decreasing_subarray

import "testing"

func Test_longestMonotonicSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{1, 4, 3, 3, 2}}, 2},
		{"", args{[]int{3, 3, 3, 3}}, 1},
		{"", args{[]int{3, 2, 1}}, 3},
		{"", args{[]int{24, 47, 24, 23, 14, 6, 9, 2, 3, 19}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestMonotonicSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestMonotonicSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
