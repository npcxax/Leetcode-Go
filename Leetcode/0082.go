package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// my solution
func deleteDuplicates(head *ListNode) *ListNode {
	var pHead = &ListNode{Next: head}
	var p = pHead

	for head != nil {
		var count = 0
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
			count++
		}
		if count > 0 {
			p.Next = head.Next
		} else {
			p.Next = head
			p = p.Next
		}
		head = p.Next
	}

	return pHead.Next
}

// official
func deleteDuplicates(head *ListNode) *ListNode {
	var dummy = &ListNode{Next: head}
	var cur = dummy

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
