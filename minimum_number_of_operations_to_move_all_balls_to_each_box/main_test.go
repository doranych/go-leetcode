package minimum_number_of_operations_to_move_all_balls_to_each_box

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minOperations(t *testing.T) {
	type args struct {
		boxes string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{"110"}, []int{1, 1, 3}},
		{"", args{"001011"}, []int{11, 8, 5, 4, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minOperations(tt.args.boxes)
			assert.Equal(t, tt.want, got)
		})
	}
}
