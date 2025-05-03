package rabbits_in_forest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numRabbits(t *testing.T) {
	tests := []struct {
		name    string
		answers []int
		want    int
	}{
		{"", []int{1, 1, 2}, 5},
		{"", []int{10, 10, 10}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numRabbits(tt.answers)
			assert.Equal(t, tt.want, got)
		})
	}
}
