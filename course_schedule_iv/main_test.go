package course_schedule_iv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkIfPrerequisite(t *testing.T) {
	type args struct {
		n   int
		req [][]int
		q   [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{"", args{2, [][]int{{1, 0}}, [][]int{{0, 1}, {1, 0}}}, []bool{false, true}},
		{"", args{2, [][]int{}, [][]int{{0, 1}, {1, 0}}}, []bool{false, false}},
		{"", args{3, [][]int{{1, 2}, {1, 0}, {2, 0}}, [][]int{{1, 0}, {1, 2}}}, []bool{true, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkIfPrerequisite(tt.args.n, tt.args.req, tt.args.q)
			assert.Equal(t, tt.want, got)
		})
	}
}
