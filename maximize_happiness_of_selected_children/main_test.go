package maximize_happiness_of_selected_children

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumHappinessSum(t *testing.T) {
	tests := []struct {
		name      string
		happiness []int
		k         int
		want      int64
	}{
		{"", []int{1, 2, 3}, 2, 4},
		{"", []int{1, 1, 1, 1, 1}, 2, 1},
		{"", []int{2, 3, 4, 5}, 1, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumHappinessSum(tt.happiness, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
