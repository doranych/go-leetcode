package recover_a_tree_from_preorder_traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			assert.Equal(t, tt.want, dumpTree(got))
		})
	}
}

// dumpTree is a function that outputs array of TreeNode values in level order
func dumpTree(node *TreeNode) []int {
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
