package construct_binary_tree_from_preorder_and_postorder_traversal

import "github.com/doranych/go-leetcode/pkg/utils"

func constructFromPrePost(preorder []int, postorder []int) *utils.TreeNode {
	stack := make([]*utils.TreeNode, 0)

	for i := range preorder {
		node := &utils.TreeNode{Val: preorder[i]}
		if len(stack) > 0 {
			curNode := stack[len(stack)-1]
			if curNode.Left == nil {
				curNode.Left = node
			} else {
				curNode.Right = node
			}
		}
		stack = append(stack, node)
		for stack[len(stack)-1].Val == postorder[0] && len(stack) > 1 {
			postorder = postorder[1:]
			stack = stack[:len(stack)-1]
		}
	}
	return stack[0]
}
