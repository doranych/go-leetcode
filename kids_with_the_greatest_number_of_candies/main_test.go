package kids_with_the_greatest_number_of_candies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kidsWithCandies(t *testing.T) {
	tests := []struct {
		name         string
		candies      []int
		extraCandies int
		want         []bool
	}{
		{"", []int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{"", []int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kidsWithCandies(tt.candies, tt.extraCandies)
			assert.Equal(t, tt.want, got)
		})
	}
}
