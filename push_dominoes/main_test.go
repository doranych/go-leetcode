package push_dominoes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pushDominoes(t *testing.T) {
	tests := []struct {
		name     string
		dominoes string
		want     string
	}{
		{"", "RR.L", "RR.L"},
		{"", ".L.R...LR..L..", "LL.RR.LLRRLL.."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pushDominoes(tt.dominoes)
			assert.Equal(t, tt.want, got)
		})
	}
}
