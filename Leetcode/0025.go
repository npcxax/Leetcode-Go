package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(left, right *ListNode) {
	var end = right.Next
	var start = left
	var pre = end
	for start != end {
		var nex = start.Next
		start.Next = pre
		pre = start
		start = nex
	}
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var pHead = &ListNode{Next: head}
	var p = pHead

	for head != nil {
		var end = head
		var count = 1
		for end != nil && count < k {
			end = end.Next
			count++
		}
		if count < k && end != nil {
			break
		}
		reverse(head, end)
		p.Next = end
		p = head
		head = head.Next
	}

	return pHead.Next
}
