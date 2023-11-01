package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// O(n) extra space
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	var stk = make([]*Node, 0)
	p := root
	stk = append(stk, p)
	for len(stk) > 0 {
		sz := len(stk)
		for i := 0; i < sz; i++ {
			if i != sz-1 {
				stk[i].Next = stk[i+1]
			}
			if stk[i].Left != nil {
				stk = append(stk, stk[i].Left)
			}
			if stk[i].Right != nil {
				stk = append(stk, stk[i].Right)
			}
		}
		stk = stk[sz:]
	}

	return root
}

// O(1) extra space
func connect(root *Node) *Node {
	p := root
	for p != nil {
		var nextBegin *Node = nil
		var pre *Node = nil
		for p != nil {
			if p.Left != nil {
				if pre == nil {
					nextBegin = p.Left
					pre = p.Left
				} else {
					pre.Next = p.Left
					pre = pre.Next
				}
			}
			if p.Right != nil {
				if pre == nil {
					nextBegin = p.Right
					pre = p.Right
				} else {
					pre.Next = p.Right
					pre = pre.Next
				}
			}
			p = p.Next
		}
		p = nextBegin
	}

	return root
}

// official
func connect(root *Node) *Node {
	start := root
	for start != nil {
		var nextStart, last *Node
		var handle = func(node *Node) {
			if node == nil {
				return
			}
			if nextStart == nil {
				nextStart = node
			}
			if last != nil {
				last.Next = node
			}
			last = node
		}
		for ; start != nil; start = start.Next {
			handle(start.Left)
			handle(start.Right)
		}
		start = nextStart
	}
	return root
}
