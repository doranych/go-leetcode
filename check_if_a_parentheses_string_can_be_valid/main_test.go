package check_if_a_parentheses_string_can_be_valid

import "testing"

func Test_canBeValid(t *testing.T) {
	type args struct {
		s      string
		locked string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"()", "00"}, true},
		{"", args{"))()))", "010100"}, true},
		{"", args{"()()", "0000"}, true},
		{"", args{")", "0"}, false},
		{"", args{"(", "0"}, false},
		{"", args{"()", "11"}, true},
		{"", args{"()))", "0010"}, true},
		{"", args{"((())(((()()", "100000111010"}, false},
		{"", args{"()()()()(((()(", "10000000000001"}, false},
		{"", args{"((()(()()))()((()()))))(((()(()", "1011110010010100111010000001001"}, false},
		{"",
			args{"())(()(()(())()())(())((())(()())((())))))(((((((())(()))))(",
				"100011110110011011010111100111011101111110000101001101001111"},
			false,
		},
		{"", args{"((()(()()))()((()()))))()((()(()", "10111100100101001110100010001001"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeValid(tt.args.s, tt.args.locked); got != tt.want {
				t.Errorf("canBeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
