package find_unique_binary_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDifferentBinaryString(t *testing.T) {
	type args struct {
		nums []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{[]string{"01", "10"}}, "00"},
		{"", args{[]string{"00", "01"}}, "10"},
		{"", args{[]string{"111", "011", "001"}}, "000"},
		{"", args{[]string{"111", "011", "101"}}, "000"},
		{"", args{[]string{"1110", "0011", "1010", "0010"}}, "0000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDifferentBinaryString(tt.args.nums)
			assert.Equal(t, tt.want, got)
		})
	}
}
