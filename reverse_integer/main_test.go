package reverse_integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"", 123, 321},
		{"", -123, -321},
		{"", 120, 21},
		{"", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse(tt.x)
			assert.Equal(t, tt.want, got)
		})
	}
}
