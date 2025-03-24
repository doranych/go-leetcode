package count_days_without_meetings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countDays(t *testing.T) {
	tests := []struct {
		name     string
		days     int
		meetings [][]int
		want     int
	}{
		{"", 7, [][]int{{7, 7}}, 6},
		{"", 7, [][]int{{1, 2}, {4, 5}}, 3},
		{"", 10, [][]int{{5, 7}, {1, 3}, {9, 10}}, 2},
		{"", 5, [][]int{{2, 4}, {1, 3}}, 1},
		{"", 6, [][]int{{1, 6}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countDays(tt.days, tt.meetings))
		})
	}
}
