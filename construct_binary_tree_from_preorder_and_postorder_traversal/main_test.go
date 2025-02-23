package construct_binary_tree_from_preorder_and_postorder_traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/doranych/go-leetcode/pkg/utils"
)

func Test_constructFromPrePost(t *testing.T) {
	tests := []struct {
		name      string
		preorder  []int
		postorder []int
		want      []int
	}{
		{"", []int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"", []int{1}, []int{1}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := constructFromPrePost(tt.preorder, tt.postorder)
			assert.Equal(t, tt.want, utils.DumpTree(got))
		})
	}
}
