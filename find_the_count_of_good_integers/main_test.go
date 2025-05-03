package find_the_count_of_good_integers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countGoodIntegers(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want int64
	}{
		{"", 3, 5, 27},
		{"", 1, 4, 2},
		{"", 5, 6, 2468},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countGoodIntegers(tt.n, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
