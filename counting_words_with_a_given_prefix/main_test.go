package counting_words_with_a_given_prefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_prefixCount(t *testing.T) {
	type args struct {
		words []string
		pref  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]string{"pay", "attention", "practice", "attend"}, "at"}, 2},
		{"", args{[]string{"leetcode", "win", "loops", "success"}, "code"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prefixCount(tt.args.words, tt.args.pref)
			assert.Equal(t, tt.want, got)
		})
	}
}
