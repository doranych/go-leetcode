package evaluate_boolean_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/doranych/go-leetcode/pkg/utils"
)

func Test_evaluateTree(t *testing.T) {
	tests := []struct {
		name string
		root *utils.TreeNode
		want bool
	}{
		{
			name: "Example 1",
			root: &utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 1},
				Right: &utils.TreeNode{Val: 3, Left: &utils.TreeNode{Val: 0}, Right: &utils.TreeNode{Val: 1}}},
			want: true,
		},
		{
			name: "Example 2",
			root: &utils.TreeNode{Val: 0},
			want: false,
		},
		{
			name: "Single AND node with false",
			root: &utils.TreeNode{Val: 3, Left: &utils.TreeNode{Val: 0}, Right: &utils.TreeNode{Val: 1}},
			want: false,
		},
		{
			name: "Single OR node with true",
			root: &utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 1}, Right: &utils.TreeNode{Val: 0}},
			want: true,
		},
		{
			name: "All leaf nodes are false",
			root: &utils.TreeNode{Val: 3, Left: &utils.TreeNode{Val: 0}, Right: &utils.TreeNode{Val: 0}},
			want: false,
		},
		{
			name: "Complex tree with mixed operations",
			root: &utils.TreeNode{Val: 2,
				Left: &utils.TreeNode{Val: 3, Left: &utils.TreeNode{Val: 1},
					Right: &utils.TreeNode{Val: 1}}, Right: &utils.TreeNode{Val: 3,
					Left: &utils.TreeNode{Val: 1}, Right: &utils.TreeNode{Val: 0},
				}},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := evaluateTree(tt.root)
			assert.Equal(t, tt.want, got)
		})
	}
}
