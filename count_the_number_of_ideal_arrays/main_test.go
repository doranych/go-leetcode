package count_the_number_of_ideal_arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_idealArrays(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		maxValue int
		want     int
	}{
		{"", 2, 5, 10},
		{"", 5, 3, 11},
		{"", 2, 1, 1},
		{"", 4, 4, 19},
		{"", 100, 100, 959337187},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := idealArrays(tt.n, tt.maxValue)
			assert.Equal(t, tt.want, got)
		})
	}
}
