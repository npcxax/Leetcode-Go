package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LNR
func getMinimumDifference(root *TreeNode) int {
	var stk = make([]*TreeNode, 0)
	var res = 0x7fffffff
	var pre = -1

	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		if pre == -1 {
			pre = root.Val
		} else {
			if root.Val-pre < res {
				res = root.Val - pre
			}
			pre = root.Val
		}
		stk = stk[:len(stk)-1]
		root = root.Right
	}

	return res
}

// dfs LNR
func getMinimumDifference(root *TreeNode) int {
	var res = math.MaxInt64
	var pre = -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && res > node.Val-pre {
			res = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return res
}
