package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	var p = head
	var n = 0
	for p != nil {
		p = p.Next
		n++
	}
	if n == 0 {
		return head
	}
	k = k % n
	if k == 0 {
		return head
	}
	var fast, slow = head, head
	for k >= 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	p = slow.Next
	slow.Next = nil
	slow = p
	for slow.Next != nil {
		slow = slow.Next
	}
	slow.Next = head
	return p
}
