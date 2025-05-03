package count_good_numbers

import "testing"

func Test_countGoodNumbers(t *testing.T) {
	tests := []struct {
		name string
		n    int64
		want int
	}{
		{"", 1, 5},
		{"", 4, 400},
		{"", 50, 564908303},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodNumbers(tt.n); got != tt.want {
				t.Errorf("countGoodNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
