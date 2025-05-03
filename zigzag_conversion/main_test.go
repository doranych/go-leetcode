package zigzag_conversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{"", args{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convert(tt.args.s, tt.args.numRows)
			assert.Equal(t, tt.want, got)
		})
	}
}
