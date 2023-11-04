package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var stk = make([]*TreeNode, 0)
	var res = make([][]int, 0)
	stk = append(stk, root)

	for len(stk) > 0 {
		sz := len(stk)
		t := make([]int, 0)
		for i := 0; i < sz; i++ {
			root = stk[i]
			t = append(t, root.Val)
			if root.Left != nil {
				stk = append(stk, root.Left)
			}
			if root.Right != nil {
				stk = append(stk, root.Right)
			}
		}
		stk = stk[sz:]
		res = append(res, t)
	}

	return res
}
