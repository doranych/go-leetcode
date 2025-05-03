package minimize_xor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimizeXor(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{3, 5}, 3},
		{"", args{1, 12}, 3},
		{"", args{65, 84}, 67},
		{"", args{255, 1000000000}, 8191},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimizeXor(tt.args.num1, tt.args.num2)
			assert.Equal(t, tt.want, got)
		})
	}
}
