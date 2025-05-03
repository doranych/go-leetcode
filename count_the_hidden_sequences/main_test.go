package count_the_hidden_sequences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numberOfArrays(t *testing.T) {
	tests := []struct {
		name        string
		differences []int
		lower       int
		upper       int
		want        int
	}{
		{"", []int{1, -3, 4}, 1, 6, 2},
		{"", []int{3, -4, 5, 1, -2}, -4, 5, 4},
		{"", []int{4, -7, 2}, 3, 6, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numberOfArrays(tt.differences, tt.lower, tt.upper)
			assert.Equal(t, tt.want, got)
		})
	}
}
