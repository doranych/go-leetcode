package count_the_number_of_powerful_integers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numberOfPowerfulInt(t *testing.T) {
	tests := []struct {
		name string
		st   int64
		fin  int64
		l    int
		s    string
		want int64
	}{
		{name: "", st: 1, fin: 6000, l: 4, s: "124", want: 5},
		{name: "", st: 15, fin: 215, l: 6, s: "10", want: 2},
		{name: "", st: 1000, fin: 2000, l: 4, s: "3000", want: 0},
		{name: "", st: 1, fin: 9, l: 9, s: "5", want: 1},
		{name: "", st: 100, fin: 200, l: 3, s: "45", want: 1},
		{name: "", st: 1, fin: 1000000, l: 5, s: "123", want: 216},
		{name: "", st: 1, fin: 100, l: 9, s: "101", want: 0}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numberOfPowerfulInt(tt.st, tt.fin, tt.l, tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
