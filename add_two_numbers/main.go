package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	firstNode := &ListNode{}
	resultNode := firstNode
	carryOne := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			resultNode.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			resultNode.Val += l2.Val
			l2 = l2.Next
		}
		if resultNode.Val > 9 {
			resultNode.Val -= 10
			carryOne = 1
		}

		if l1 != nil || l2 != nil || carryOne != 0 {
			resultNode.Next = &ListNode{}
			resultNode = resultNode.Next
			resultNode.Val = carryOne
			carryOne = 0
		}
	}
	return firstNode
}
