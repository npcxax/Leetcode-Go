package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// my solution
// type BSTIterator struct {
// 	stk  []*TreeNode
// 	idx  int
// 	root *TreeNode
// }

// func solve(b *BSTIterator, node *TreeNode) {
// 	if node == nil {
// 		return
// 	}
// 	solve(b, node.Left)
// 	b.stk = append(b.stk, node)
// 	solve(b, node.Right)
// }

// func Constructor(root *TreeNode) BSTIterator {
// 	b := BSTIterator{
// 		stk:  make([]*TreeNode, 0),
// 		idx:  -1,
// 		root: root,
// 	}
// 	solve(&b, root)
// 	return b
// }

// func (this *BSTIterator) Next() int {
// 	this.idx++
// 	return this.stk[this.idx].Val
// }

// func (this *BSTIterator) HasNext() bool {
// 	return this.idx+1 < len(this.stk)
// }

// official
// type BSTIterator struct {
// 	stk []*TreeNode
// }

// func Constructor(root *TreeNode) BSTIterator {
// 	b := BSTIterator{stk: make([]*TreeNode, 0)}
// 	b.inorder(root)
// 	return b
// }

// func (this *BSTIterator) Next() int {
// 	val := this.stk[0]
// 	this.stk = this.stk[1:]
// 	return val.Val
// }

// func (this *BSTIterator) HasNext() bool {
// 	return len(this.stk) > 0
// }

// func (this *BSTIterator) inorder(node *TreeNode) {
// 	if node == nil {
// 		return
// 	}
// 	this.inorder(node.Left)
// 	this.stk = append(this.stk, node)
// 	this.inorder(node.Right)
// }

// iterator
type BSTIterator struct {
	stk  []*TreeNode
	node *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	b := BSTIterator{stk: make([]*TreeNode, 0), node: root}
	return b
}

func (this *BSTIterator) Next() int {
	for this.node != nil {
		this.stk = append(this.stk, this.node)
		this.node = this.node.Left
	}
	this.node = this.stk[len(this.stk)-1]
	this.stk = this.stk[:len(this.stk)-1]
	val := this.node.Val
	this.node = this.node.Right
	return val
}

func (this *BSTIterator) HasNext() bool {
	return this.node != nil || len(this.stk) != 0
}
