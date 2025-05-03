package count_the_hidden_sequences

import "testing"

func Test_numberOfArrays(t *testing.T) {
	tests := []struct {
		name        string
		differences []int
		lower       int
		upper       int
		want        int
	}{
		{"", []int{1, -3, 4}, 1, 6, 2},
		{"", []int{3, -4, 5, 1, -2}, -4, 5, 4},
		{"", []int{4, -7, 2}, 3, 6, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArrays(tt.differences, tt.lower, tt.upper); got != tt.want {
				t.Errorf("numberOfArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
