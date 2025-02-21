package find_elements_in_a_contaminated_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	t *TreeNode
	m map[int]bool
}

type qe struct {
	t *TreeNode
	v int
}

func Constructor(root *TreeNode) FindElements {
	q := []qe{{root, 0}}
	m := make(map[int]bool)
	for len(q) > 0 {
		te := q[0]
		q = q[1:]
		te.t.Val = te.v
		m[te.v] = true
		if te.t.Left != nil {
			q = append(q, qe{te.t.Left, 2*te.t.Val + 1})
		}
		if te.t.Right != nil {
			q = append(q, qe{te.t.Right, 2*te.t.Val + 2})
		}
	}
	return FindElements{root, m}
}

func (f *FindElements) Find(target int) bool {
	return f.m[target]
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
