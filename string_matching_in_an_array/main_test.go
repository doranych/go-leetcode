package string_matching_in_an_array

import (
	"reflect"
	"testing"
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
			if got := stringMatching(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}
