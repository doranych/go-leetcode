package string_to_integer_atoi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"42", args{"42"}, 42},
		{"   -42", args{"   -42"}, -42},
		{"4193 with words", args{"4193 with words"}, 4193},
		{"0-1", args{"0-1"}, 0},
		{"+-12", args{"+-12"}, 0},
		{"words and 987", args{"words and 987"}, 0},
		{"2147483648", args{"2147483648"}, 2147483647},
		{"-91283472332", args{"-91283472332"}, -2147483648},
		{"91283472332", args{"91283472332"}, 2147483647},
		{"20000000000000000000", args{"20000000000000000000"}, 2147483647},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := myAtoi(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
