package count_the_number_of_good_subarrays

import "testing"

func Test_countGood(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int64
	}{
		{"", []int{1, 1, 1, 1, 1}, 10, 1},
		{"", []int{3, 1, 4, 3, 2, 2, 4}, 2, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGood(tt.nums, tt.k); got != tt.want {
				t.Errorf("countGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
