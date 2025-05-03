package can_place_flowers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canPlaceFlowers(t *testing.T) {
	tests := []struct {
		name      string
		flowerbed []int
		n         int
		want      bool
	}{
		{"", []int{1, 0, 0, 0, 1}, 1, true},
		{"", []int{1, 0, 0, 0, 1}, 2, false},
		{"", []int{0, 0, 0, 0, 0}, 2, true},
		{"", []int{1, 1, 1, 1, 1}, 1, false},
		{"", []int{0}, 1, true},
		{"", []int{1}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canPlaceFlowers(tt.flowerbed, tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
