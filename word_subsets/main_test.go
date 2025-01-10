package word_subsets

import (
	"reflect"
	"testing"
)

func Test_wordSubsets(t *testing.T) {
	type args struct {
		words1 []string
		words2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"", args{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"}}, []string{"facebook", "google", "leetcode"}},
		{"", args{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"l", "e"}}, []string{"apple", "google", "leetcode"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordSubsets(tt.args.words1, tt.args.words2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
