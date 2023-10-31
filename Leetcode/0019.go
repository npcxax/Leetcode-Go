package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var pHead = &ListNode{Next: head}
	var fast, slow = pHead, pHead
	for n >= 0 {
		fast = fast.Next
		n--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return pHead.Next
}
