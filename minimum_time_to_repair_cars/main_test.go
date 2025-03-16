package minimum_time_to_repair_cars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_repairCars(t *testing.T) {
	tests := []struct {
		name  string
		ranks []int
		cars  int
		want  int64
	}{
		{"", []int{4, 2, 3, 1}, 10, 16},
		{"", []int{5, 1, 8}, 6, 16},
		{"", []int{1, 1, 3, 3}, 74, 576},
		{"", []int{1, 3, 1, 2, 1, 1}, 21, 18},
		{"", []int{3, 1, 3, 3, 1, 3}, 42, 108},
		{"", []int{3, 3, 1, 2, 1, 1, 3, 2, 1}, 58, 75},
		{"", []int{2, 2, 3, 3, 1, 3, 3, 1, 3}, 32, 32},
		{"", []int{2, 2, 1, 3, 1, 2, 1, 1, 1, 3}, 36, 25},
		{"", []int{5, 3, 1, 2, 6, 5, 3, 1, 5, 7, 1, 4, 5, 1, 6}, 2306, 61347},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, repairCars(tt.ranks, tt.cars))
		})
	}
}
