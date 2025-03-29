package apply_operations_to_maximize_score

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumScore(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"", []int{8, 3, 9, 3, 8}, 2, 81},
		{"", []int{19, 12, 14, 6, 10, 18}, 3, 4788},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maximumScore(tt.nums, tt.k))
		})
	}
}
