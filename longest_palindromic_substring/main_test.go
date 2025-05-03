package longest_palindromic_substring

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"", "babad", "bab"},
		{"", "cbbd", "bb"},
		{"", "a", "a"},
		{"", "ac", "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
