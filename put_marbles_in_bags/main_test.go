package put_marbles_in_bags

import "testing"

func Test_putMarbles(t *testing.T) {
	tests := []struct {
		name    string
		weights []int
		k       int
		want    int64
	}{
		{"", []int{1, 3, 5, 1}, 2, 4},
		{"", []int{1, 3, 5, 1}, 3, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := putMarbles(tt.weights, tt.k); got != tt.want {
				t.Errorf("putMarbles() = %v, want %v", got, tt.want)
			}
		})
	}
}
