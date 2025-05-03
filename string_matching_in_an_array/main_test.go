package string_matching_in_an_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stringMatching(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"", args{[]string{"mass", "as", "hero", "superhero"}}, []string{"as", "hero"}},
		{"", args{[]string{"leetcode", "et", "code"}}, []string{"et", "code"}},
		{"", args{[]string{"blue", "green", "bu"}}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringMatching(tt.args.words)
			assert.Equal(t, tt.want, got)
		})
	}
}
