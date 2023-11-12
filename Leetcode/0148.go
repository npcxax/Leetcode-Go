package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// from bottom to top, my solution
func sortList(head *ListNode) *ListNode {
	var (
		length = 0
		pHead  = &ListNode{Next: head}
		merge  func(p1 *ListNode, p2 *ListNode) (*ListNode, *ListNode)
	)

	merge = func(p1, p2 *ListNode) (*ListNode, *ListNode) {
		var pHead = &ListNode{}
		var p = pHead
		for p1 != nil || p2 != nil {
			if p1 == nil {
				p.Next = p2
				p2 = p2.Next
			} else if p2 == nil {
				p.Next = p1
				p1 = p1.Next
			} else {
				if p1.Val > p2.Val {
					p.Next = p2
					p2 = p2.Next
				} else {
					p.Next = p1
					p1 = p1.Next
				}
			}
			p = p.Next
		}

		return pHead.Next, p
	}

	for p := head; p != nil; p = p.Next {
		length++

	}

	for subLength := 1; subLength < length; subLength <<= 1 {
		p := pHead
		for p.Next != nil {
			p1 := p
			var i = 0
			for i < subLength && p1.Next != nil {
				p1 = p1.Next
				i++
			}
			if i < subLength || (i == subLength && p1.Next == nil) {
				break
			}
			p2 := p1
			i = 0
			for i < subLength && p2.Next != nil {
				p2 = p2.Next
				i++
			}
			var p4 = p2
			if p2 != nil {
				p2 = p2.Next
				p4.Next = nil
			}
			p3 := p1.Next
			p1.Next = nil
			left, right := merge(p.Next, p3)
			p.Next = left
			p = right
			p.Next = p2
		}
	}

	return pHead.Next
}

// official, from bottom to top
func sortList(head *ListNode) *ListNode {
	var (
		length = 0
		merge  func(p1, p2 *ListNode) *ListNode
		pHead  = &ListNode{Next: head}
	)

	merge = func(p1, p2 *ListNode) *ListNode {
		var pHead = &ListNode{}
		var p = pHead
		for p1 != nil && p2 != nil {
			if p1.Val > p2.Val {
				p.Next = p2
				p2 = p2.Next
			} else {
				p.Next = p1
				p1 = p1.Next
			}
			p = p.Next
		}
		if p1 == nil {
			p.Next = p2
		} else {
			p.Next = p1
		}
		return pHead.Next
	}

	for p := head; p != nil; p = p.Next {
		length++
	}

	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := pHead, pHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}
			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength && cur != nil; i++ {
				cur = cur.Next
			}
			var next *ListNode = nil
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			prev.Next = merge(head1, head2)
			for prev.Next != nil {
				prev = prev.Next
			}

			cur = next
		}
	}

	return pHead.Next
}

// from top to bottom
func sortList(head *ListNode) *ListNode {
	var (
		sortL func(head, tail *ListNode) *ListNode
		merge func(p1, p2 *ListNode) *ListNode
	)

	merge = func(p1, p2 *ListNode) *ListNode {
		var pHead = &ListNode{}
		var p = pHead
		for p1 != nil && p2 != nil {
			if p1.Val > p2.Val {
				p.Next = p2
				p2 = p2.Next
			} else {
				p.Next = p1
				p1 = p1.Next
			}
			p = p.Next
		}
		if p1 == nil {
			p.Next = p2
		} else {
			p.Next = p1
		}
		return pHead.Next
	}

	sortL = func(head, tail *ListNode) *ListNode {
		if head == nil {
			return nil
		}
		if head.Next == tail {
			head.Next = nil
			return head
		}
		slow, fast := head, head
		for fast != tail {
			slow = slow.Next
			fast = fast.Next
			if fast != tail {
				fast = fast.Next
			}
		}
		mid := slow
		return merge(sortL(head, mid), sortL(mid, tail))
	}

	return sortL(head, nil)
}
