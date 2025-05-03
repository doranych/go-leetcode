package count_the_number_of_ideal_arrays

import "testing"

func Test_idealArrays(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		maxValue int
		want     int
	}{
		{"", 2, 5, 10},
		{"", 5, 3, 11},
		{"", 2, 1, 1},
		{"", 4, 4, 19},
		{"", 100, 100, 959337187},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := idealArrays(tt.n, tt.maxValue); got != tt.want {
				t.Errorf("idealArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
