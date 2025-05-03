package special_array_i

import "testing"

func Test_isArraySpecial(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"", []int{1, 2, 3, 4}, true},
		{"", []int{1, 2, 3}, true},
		{"", []int{4, 3, 1, 6}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isArraySpecial(tt.nums); got != tt.want {
				t.Errorf("isArraySpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}
