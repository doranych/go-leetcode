package evaluate_boolean_binary_tree

import "github.com/doranych/go-leetcode/pkg/utils"

func evaluateTree(root *utils.TreeNode) bool {
	switch root.Val {
	case 0:
		return false
	case 1:
		return true
	case 2:
		l := evaluateTree(root.Left)
		if l {
			return true
		}
		r := evaluateTree(root.Right)
		return r
	case 3:
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	default:
		return false
	}
}
