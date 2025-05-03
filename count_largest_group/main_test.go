package count_largest_group

import "testing"

func Test_countLargestGroup(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"", 13, 4},
		{"", 2, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countLargestGroup(tt.n); got != tt.want {
				t.Errorf("countLargestGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
