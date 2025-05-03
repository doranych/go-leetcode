package clear_digits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_clearDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test case 2", args{"abc"}, "abc"},
		{"Test case 2", args{"cb34"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := clearDigits(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
