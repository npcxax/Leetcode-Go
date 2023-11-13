package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// my solution
func mergeKLists(lists []*ListNode) *ListNode {
	var (
		n     = len(lists)
		merge func(l1, l2 *ListNode) *ListNode
	)

	if n == 0 {
		return nil
	}

	merge = func(l1, l2 *ListNode) *ListNode {
		l := &ListNode{}
		p := l

		for l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				l.Next = l1
				l1 = l1.Next
			} else {
				l.Next = l2
				l2 = l2.Next
			}
			l = l.Next
		}

		if l1 != nil {
			l.Next = l1
		} else {
			l.Next = l2
		}

		return p.Next
	}

	for len(lists) > 1 {
		sz := len(lists)
		if sz%2 == 1 {
			sz -= 1
		}
		for i := 0; i < sz; i += 2 {
			lists = append(lists, merge(lists[i], lists[i+1]))
		}
		lists = lists[sz:]
	}

	return lists[0]
}
