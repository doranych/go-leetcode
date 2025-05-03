package count_subarrays_where_max_element_appears_at_least_k_times

import "testing"

func Test_countSubarrays(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int64
	}{
		{"", []int{1, 3, 2, 3, 3}, 2, 6},
		{"", []int{1, 4, 2, 1}, 3, 0},
		{"", []int{5}, 1, 1},
		{"", []int{2, 2, 2, 2}, 2, 6},
		{"", []int{1, 2, 3, 4}, 5, 0},
		{"", []int{3, 3, 3, 3}, 4, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarrays(tt.nums, tt.k); got != tt.want {
				t.Errorf("countSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
