package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bfs
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res = make([]int, 0)
	var stk = make([]*TreeNode, 0)
	stk = append(stk, root)

	for len(stk) > 0 {
		sz := len(stk)
		res = append(res, stk[sz-1].Val)
		for i := 0; i < sz; i++ {
			root = stk[i]
			if root.Left != nil {
				stk = append(stk, root.Left)
			}
			if root.Right != nil {
				stk = append(stk, root.Right)
			}
		}
		stk = stk[sz:]
	}
	return res
}
