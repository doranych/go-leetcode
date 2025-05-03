package count_complete_subarrays_in_an_array

import "testing"

func Test_countCompleteSubarrays(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"", []int{1}, 1},
		{"", []int{1, 2, 3, 4}, 1},
		{"", []int{1, 2, 1, 2, 1}, 10},
		{"", []int{1, 3, 1, 2, 2}, 4},
		{"", []int{5, 5, 5, 5}, 10},
		{"", []int{1, 2, 3, 1, 2, 3}, 10},
		{"", []int{1, 1, 2, 3, 2, 1}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCompleteSubarrays(tt.nums); got != tt.want {
				t.Errorf("countCompleteSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
