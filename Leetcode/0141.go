package leetcode

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var fast, slow = head.Next, head

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	return true
}
