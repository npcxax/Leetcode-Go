package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var res = make([]float64, 0)
	var stk = make([]*TreeNode, 0)
	stk = append(stk, root)
	for len(stk) > 0 {
		sz := len(stk)
		total := 0.0
		for i := 0; i < sz; i++ {
			root = stk[i]
			total += float64(root.Val)
			if root.Left != nil {
				stk = append(stk, root.Left)
			}
			if root.Right != nil {
				stk = append(stk, root.Right)
			}
		}
		res = append(res, total/float64(sz))
		stk = stk[sz:]
	}
	return res
}
