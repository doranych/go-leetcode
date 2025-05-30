package check_if_one_string_swap_can_make_strings_equal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_areAlmostEqual(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"bank", "kanb"}, true},
		{"", args{"attack", "defend"}, false},
		{"", args{"kelb", "kelb"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := areAlmostEqual(tt.args.s1, tt.args.s2)
			assert.Equal(t, tt.want, got)
		})
	}
}
