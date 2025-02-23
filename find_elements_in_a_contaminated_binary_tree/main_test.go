package find_elements_in_a_contaminated_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/doranych/go-leetcode/pkg/utils"
)

func Test_Find(t *testing.T) {
	tests := []struct {
		name       string
		makeTree   func() *utils.TreeNode
		makeResult func() ([]int, []bool)
	}{
		{
			"", func() *utils.TreeNode {
				root := &utils.TreeNode{Val: -1}
				root.Right = &utils.TreeNode{Val: -1}
				return root
			}, func() ([]int, []bool) {
				return []int{1, 2}, []bool{false, true}
			},
		},
		{
			"", func() *utils.TreeNode {
				root := &utils.TreeNode{Val: -1}
				root.Left = &utils.TreeNode{Val: -1}
				root.Right = &utils.TreeNode{Val: -1}
				root.Left.Left = &utils.TreeNode{Val: -1}
				root.Left.Right = &utils.TreeNode{Val: -1}
				return root
			}, func() ([]int, []bool) {
				return []int{1, 3, 5}, []bool{true, true, false}
			},
		},
		{
			"", func() *utils.TreeNode {
				root := &utils.TreeNode{Val: -1}
				root.Right = &utils.TreeNode{Val: -1}
				root.Right.Left = &utils.TreeNode{Val: -1}
				root.Right.Left.Left = &utils.TreeNode{Val: -1}
				return root
			}, func() ([]int, []bool) {
				return []int{2, 3, 4, 5}, []bool{true, false, false, true}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := tt.makeTree()
			obj := Constructor(root)
			result, expected := tt.makeResult()
			for i, r := range result {
				assert.Equal(t, expected[i], obj.Find(r))
			}
		})
	}
}
