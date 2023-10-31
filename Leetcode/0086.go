package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var pHead = &ListNode{}
	var p = pHead
	var pXHead = &ListNode{}
	var pX = pXHead

	for head != nil {
		if head.Val < x {
			p.Next = head
			p = p.Next
		} else {
			pX.Next = head
			pX = pX.Next
		}
		head = head.Next
		if head == nil {
			pX.Next = nil
		}
	}
	p.Next = pXHead.Next

	return pHead.Next
}
