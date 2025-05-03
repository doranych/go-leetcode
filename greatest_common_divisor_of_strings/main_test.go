package greatest_common_divisor_of_strings

import "testing"

func Test_gcdOfStrings(t *testing.T) {
	tests := []struct {
		name string
		str1 string
		str2 string
		want string
	}{
		{"", "ABCABC", "ABC", "ABC"},
		{"", "ABABAB", "ABAB", "AB"},
		{"", "LEET", "CODE", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcdOfStrings(tt.str1, tt.str2); got != tt.want {
				t.Errorf("gcdOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
