package recover_a_tree_from_preorder_traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/doranych/go-leetcode/pkg/utils"
)

func Test_recoverFromPreorder(t *testing.T) {
	tests := []struct {
		name      string
		traversal string
		want      []int
	}{
		{"", "1-2--3--4-5--6--7", []int{1, 2, 5, 3, 4, 6, 7}},
		{"", "1-2--3---4-5--6---7", []int{1, 2, 5, 3, 0, 6, 0, 4, 0, 7}},
		{"", "1-401--349---90--88", []int{1, 401, 0, 349, 88, 90}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := recoverFromPreorder(tt.traversal)
			assert.Equal(t, tt.want, utils.DumpTree(got))
		})
	}
}
