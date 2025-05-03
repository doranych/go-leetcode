package maximum_candies_allocated_to_k_children

import "testing"

func Test_maximumCandies(t *testing.T) {
	tests := []struct {
		name    string
		candies []int
		k       int64
		want    int
	}{
		{"", []int{5, 8, 6}, 3, 5},
		{"", []int{2, 5}, 11, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCandies(tt.candies, tt.k); got != tt.want {
				t.Errorf("maximumCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
