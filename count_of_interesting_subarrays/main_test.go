package count_of_interesting_subarrays

import "testing"

func Test_countInterestingSubarrays(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		nums []int
		mod  int
		k    int
		want int64
	}{
		{name: "", nums: []int{3, 2, 4}, mod: 2, k: 1, want: 3},
		{name: "", nums: []int{3, 1, 9, 6}, mod: 3, k: 0, want: 2},
		{name: "", nums: []int{1, 2, 3}, mod: 5, k: 4, want: 0},
		{name: "", nums: []int{5, 10, 15}, mod: 5, k: 0, want: 0},
		{name: "", nums: []int{7}, mod: 3, k: 1, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countInterestingSubarrays(tt.nums, tt.mod, tt.k); got != tt.want {
				t.Errorf("countInterestingSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
