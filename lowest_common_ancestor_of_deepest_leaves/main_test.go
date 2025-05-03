package lowest_common_ancestor_of_deepest_leaves

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/doranych/go-leetcode/pkg/utils"
)

func Test_lcaDeepestLeaves(t *testing.T) {
	tests := []struct {
		name string
		root *utils.TreeNode
		want *utils.TreeNode
	}{
		{
			"",
			&utils.TreeNode{Val: 3, Left: &utils.TreeNode{Val: 5, Left: &utils.TreeNode{Val: 6},
				Right: &utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 7}, Right: &utils.TreeNode{Val: 4}},
			}, Right: &utils.TreeNode{Val: 1, Left: &utils.TreeNode{Val: 0}, Right: &utils.TreeNode{Val: 8}}},
			&utils.TreeNode{Val: 2, Left: &utils.TreeNode{Val: 7}, Right: &utils.TreeNode{Val: 4}},
		},
		{
			"",
			&utils.TreeNode{Val: 1}, &utils.TreeNode{Val: 1},
		},
		{
			"",
			&utils.TreeNode{Val: 0, Left: &utils.TreeNode{Val: 1, Right: &utils.TreeNode{Val: 2}}, Right: &utils.TreeNode{Val: 3}},
			&utils.TreeNode{Val: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lcaDeepestLeaves(tt.root)
			assert.Equal(t, tt.want, got)
		})
	}
}
