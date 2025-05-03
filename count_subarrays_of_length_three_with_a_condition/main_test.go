package count_subarrays_of_length_three_with_a_condition

import "testing"

func Test_countSubarrays(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"", []int{1, 2, 1, 4, 1}, 1},
		{"", []int{1, 1, 1}, 0},
		{"", []int{2, 3, 5, 7, 11}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarrays(tt.nums); got != tt.want {
				t.Errorf("countSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
