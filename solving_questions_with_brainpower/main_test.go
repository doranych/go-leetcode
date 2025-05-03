package solving_questions_with_brainpower

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mostPoints(t *testing.T) {
	tests := []struct {
		name      string
		questions [][]int
		want      int64
	}{
		{"", [][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}}, 5},
		{"", [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mostPoints(tt.questions)
			assert.Equal(t, tt.want, got)
		})
	}
}
