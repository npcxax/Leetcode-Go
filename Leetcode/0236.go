package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans *TreeNode = nil

func f(node, p, q *TreeNode) bool {
	if node == nil {
		return false
	}
	fp, fq := f(node.Left, p, q), f(node.Right, p, q)
	if (fp && fq) || ((node == p || node == q) && (fp || fq)) {
		ans = node
	}
	return fp || fq || (p.Val == node.Val || node.Val == q.Val)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	f(root, p, q)
	return ans
}

// official
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	lson := lowestCommonAncestor(root.Left, p, q)
	rson := lowestCommonAncestor(root.Right, p, q)
	if lson != nil && rson != nil {
		return root
	}
	if lson == nil {
		return rson
	}
	return lson
}
