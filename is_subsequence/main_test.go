package is_subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"abc", "ahbgdc"}, true},
		{"", args{"axc", "ahbgdc"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSubsequence(tt.args.s, tt.args.t)
			assert.Equal(t, tt.want, got)
		})
	}
}
