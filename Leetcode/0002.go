package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var add = 0
	var p = new(ListNode)
	var pHead = p
	for l1 != nil || l2 != nil {
		var n1, n2 = 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		p.Next = &ListNode{Val: (n1 + n2 + add) % 10}
		add = (n1 + n2 + add) / 10
		p = p.Next
	}
	if add == 1 {
		p.Next = &ListNode{Val: add}
	}

	return pHead.Next
}
