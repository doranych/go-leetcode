package recover_a_tree_from_preorder_traversal

import "github.com/doranych/go-leetcode/pkg/utils"

func recoverFromPreorder(traversal string) *utils.TreeNode {
	stack := make([]*utils.TreeNode, 0)
	depth := 0
	curVal := 0
	for i, r := range traversal {
		if r >= '0' && r <= '9' {
			curVal = curVal*10 + int(r-'0')
		}
		if r == '-' || i == len(traversal)-1 {
			if curVal != 0 {
				node := &utils.TreeNode{Val: curVal}
				curVal = 0
				if len(stack) > depth {
					stack = stack[:depth]
				}
				if len(stack) > 0 {
					curNode := stack[len(stack)-1]
					if curNode.Left == nil {
						curNode.Left = node
					} else {
						curNode.Right = node
					}
				}
				stack = append(stack, node)
				depth = 0
			}
			depth++
		}
	}
	return stack[0]
}

/* non optimized solution switching current node when level is not matched
func recoverFromPreorder(traversal string) *TreeNode {
	parents := make(map[*TreeNode]*TreeNode, 0)
	root := &TreeNode{}
	dashCounter := 0
	curLevel := 0
	curNum := ""
	curNode := root
	for i, r := range traversal {
		if r >= '0' && r <= '9' {
			curNum += string(r)
			if dashCounter > 0 && dashCounter <curLevel {
				for dashCounter < curLevel {
					if parents[curNode] != nil {
						curNode = parents[curNode]
					} else {
						curNode = root
					}
					curLevel--
				}
			} else if dashCounter > curLevel && curLevel != 0 {
				if curNode.Right != nil {
					curNode = curNode.Right
				} else {
					curNode = curNode.Left
				}
			}
			curLevel = dashCounter
		}
		if r == '-' || i == len(traversal) -1 {
			if curNum != "" {
				node := &TreeNode{}
				if dashCounter == 0 { //first dash sending root val
					node = root
				}
				node.Val, _ = strconv.Atoi(curNum)
				curNum = ""
				if node != root {
					parents[node] = curNode
					if curNode.Left == nil {
						curNode.Left = node
					} else {
						curNode.Right = node
					}
				}
				dashCounter = 0
			}
			dashCounter++
		}
	}
	return root
}
*/
