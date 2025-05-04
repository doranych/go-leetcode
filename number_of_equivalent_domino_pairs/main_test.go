package number_of_equivalent_domino_pairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numEquivDominoPairs(t *testing.T) {
	tests := []struct {
		name     string
		dominoes [][]int
		want     int
	}{
		{"", [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}, 1},
		{"", [][]int{{1, 2}, {1, 2}, {1, 1}, {2, 2}, {1, 2}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numEquivDominoPairs(tt.dominoes)
			assert.Equal(t, tt.want, got)
		})
	}
}
