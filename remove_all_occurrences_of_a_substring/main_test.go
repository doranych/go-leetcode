package remove_all_occurrences_of_a_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeOccurrences(t *testing.T) {
	type args struct {
		s    string
		part string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"abc", "abc"}, ""},
		{"", args{"cb34", "34"}, "cb"},
		{"", args{"daabcbaabcbc", "abc"}, "dab"},
		{"", args{"axxxxyyyyb", "xy"}, "ab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeOccurrences(tt.args.s, tt.args.part)
			assert.Equal(t, tt.want, got)
		})
	}
}
