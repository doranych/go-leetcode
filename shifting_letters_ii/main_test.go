package shifting_letters_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shiftingLetters(t *testing.T) {
	type args struct {
		s      string
		shifts [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"abc", [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}}, "ace"},
		{"", args{"dztz", [][]int{{0, 0, 0}, {1, 1, 1}}}, "catz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shiftingLetters(tt.args.s, tt.args.shifts)
			assert.Equal(t, tt.want, got)
		})
	}
}
