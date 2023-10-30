package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var pHead = &ListNode{Next: head}
	var count = 1
	var start *ListNode = pHead
	for count < left {
		start = head
		head = head.Next
		count++
	}
	var begin = start
	var over = head
	for count <= right {
		var end = head.Next
		head.Next = start
		start = head
		head = end
		count++
	}
	begin.Next = start
	over.Next = head
	return pHead.Next
}
