package count_equal_and_divisible_pairs_in_an_array

import "testing"

func Test_countPairs(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"", []int{3, 1, 2, 2, 2, 1, 3}, 2, 4},
		{"", []int{1, 2, 3, 4}, 1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.nums, tt.k); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
