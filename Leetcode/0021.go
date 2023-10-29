package leetcode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var p = &ListNode{}
	var pHead = p

	for list1 != nil || list2 != nil {
		if list1 == nil {
			p.Next = list2
			list2 = list2.Next
		} else if list2 == nil {
			p.Next = list1
			list1 = list1.Next
		} else {
			if list1.Val < list2.Val {
				p.Next = list1
				list1 = list1.Next
			} else {
				p.Next = list2
				list2 = list2.Next
			}
		}
		p = p.Next
	}

	return pHead.Next
}
