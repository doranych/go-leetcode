package maximum_number_of_tasks_you_can_assign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxTaskAssign(t *testing.T) {
	tests := []struct {
		name     string
		tasks    []int
		workers  []int
		pills    int
		strength int
		want     int
	}{
		{"", []int{3, 2, 1}, []int{0, 3, 3}, 1, 1, 3},
		{"", []int{5, 4}, []int{0, 0, 0}, 1, 5, 1},
		{"", []int{10, 15, 30}, []int{0, 10, 10, 10, 10}, 3, 10, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxTaskAssign(tt.tasks, tt.workers, tt.pills, tt.strength)
			assert.Equal(t, tt.want, got)
		})
	}
}
