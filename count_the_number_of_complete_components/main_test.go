package count_the_number_of_complete_components

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countCompleteComponents(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  int
	}{
		{"", 6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}, 3},
		{"", 6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countCompleteComponents(tt.n, tt.edges))
		})
	}
}
