package maximum_absolute_sum_of_any_subarray

import "testing"

func Test_maxAbsoluteSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"", []int{1, -3, 2, 3, -4}, 5},
		{"", []int{2, -5, 1, -4, 3, -2}, 8},
		{"", []int{3, -1, 0, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAbsoluteSum(tt.nums); got != tt.want {
				t.Errorf("maxAbsoluteSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
