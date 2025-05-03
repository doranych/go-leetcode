package minimum_length_of_string_after_operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"aabccabba"}, 5},
		{"", args{"bb"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumLength(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
