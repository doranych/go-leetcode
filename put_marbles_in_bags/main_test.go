package put_marbles_in_bags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_putMarbles(t *testing.T) {
	tests := []struct {
		name    string
		weights []int
		k       int
		want    int64
	}{
		{"", []int{1, 3, 5, 1}, 2, 4},
		{"", []int{1, 3, 5, 1}, 3, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := putMarbles(tt.weights, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
