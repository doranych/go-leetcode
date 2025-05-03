package push_dominoes

import "testing"

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
			if got := pushDominoes(tt.dominoes); got != tt.want {
				t.Errorf("pushDominoes() = %v, want %v", got, tt.want)
			}
		})
	}
}
