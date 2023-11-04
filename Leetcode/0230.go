package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LNR
func kthSmallest(root *TreeNode, k int) int {
	var stk = make([]*TreeNode, 0)
	var count = 0
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		count++
		if count == k {
			return root.Val
		}
		root = root.Right
	}
	return -1
}

// TODO: AVL tree
