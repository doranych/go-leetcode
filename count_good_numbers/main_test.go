package count_good_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countGoodNumbers(t *testing.T) {
	tests := []struct {
		name string
		n    int64
		want int
	}{
		{"", 1, 5},
		{"", 4, 400},
		{"", 50, 564908303},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countGoodNumbers(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
