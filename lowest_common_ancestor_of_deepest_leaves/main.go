package lowest_common_ancestor_of_deepest_leaves

import "github.com/doranych/go-leetcode/pkg/utils"

func lcaDeepestLeaves(root *utils.TreeNode) *utils.TreeNode {
	_, tree := dfs(root)
	return tree
}

func dfs(root *utils.TreeNode) (int, *utils.TreeNode) {
	if root == nil {
		return 0, nil
	}
	lDepth, lTree := dfs(root.Left)
	rDepth, rTree := dfs(root.Right)
	if lDepth > rDepth {
		return lDepth + 1, lTree
	}
	if rDepth > lDepth {
		return rDepth + 1, rTree
	}
	return lDepth + 1, root
}
