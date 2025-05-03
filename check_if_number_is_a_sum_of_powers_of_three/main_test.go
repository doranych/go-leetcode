package check_if_number_is_a_sum_of_powers_of_three

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkPowersOfThree(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"", 12, true},
		{"", 91, true},
		{"", 21, false},
		{"", 1, true},
		{"", 10000000, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkPowersOfThree(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
