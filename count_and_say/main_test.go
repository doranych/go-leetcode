package count_and_say

import "testing"

func Test_countAndSay(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{"Test 1", 1, "1"},
		{"Test 2", 2, "11"},
		{"Test 3", 3, "21"},
		{"Test 4", 4, "1211"},
		{"Test 5", 5, "111221"},
		{"Test 6", 6, "312211"},
		{"Test 7", 7, "13112221"},
		{"Test 8", 8, "1113213211"},
		{"Test 9", 9, "31131211131221"},
		{"Test 10", 10, "13211311123113112211"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
