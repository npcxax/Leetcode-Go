package leetcode

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var p = &Node{Next: head}

	// create new nodes
	for head != nil {
		var tp = &Node{Next: head.Next, Val: head.Val}
		head.Next = tp
		head = tp.Next
	}

	head = p.Next

	for head != nil {
		if head.Random != nil {
			head.Next.Random = head.Random.Next
		}
		head = head.Next.Next
	}

	head = p.Next
	var pHead = p

	for head != nil {
		p.Next = head.Next
		head.Next = head.Next.Next
		p = p.Next
		head = head.Next
	}

	return pHead.Next
}
