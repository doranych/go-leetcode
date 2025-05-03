package count_symmetric_integers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countSymmetricIntegers(t *testing.T) {
	tests := []struct {
		name string
		low  int
		high int
		want int
	}{
		{"test1", 1, 100, 9},
		{"test2", 1200, 1230, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countSymmetricIntegers(tt.low, tt.high)
			assert.Equal(t, tt.want, got)
		})
	}
}
