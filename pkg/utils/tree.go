package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DumpTree is a function that outputs array of TreeNode values in level order
func DumpTree(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{node}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current != nil {
			result = append(result, current.Val)
			queue = append(queue, current.Left, current.Right)
		} else {
			result = append(result, 0)
		}
	}

	// Remove trailing zeros
	for len(result) > 0 && result[len(result)-1] == 0 {
		result = result[:len(result)-1]
	}

	return result
}
