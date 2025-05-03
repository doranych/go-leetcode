package maximum_score_after_splitting_a_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxScore(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"011101"}, 5},
		{"", args{"00111"}, 5},
		{"", args{"1111"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxScore(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
