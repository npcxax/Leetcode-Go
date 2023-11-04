package leetcode

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LNR
func isValidBST(root *TreeNode) bool {
	var stk = make([]*TreeNode, 0)
	var pre = 0
	var flag = true

	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		fmt.Printf("root=%d,pre=%d\n", root.Val, pre)
		if !flag && pre >= root.Val {
			return false
		}
		flag = false
		pre = root.Val
		root = root.Right
	}
	return true
}

// official
func helper(node *TreeNode, low int, high int) bool {
	if node == nil {
		return true
	}
	if node.Val <= low || node.Val >= high {
		return false
	}
	return helper(node.Left, low, node.Val) && helper(node.Right, node.Val, high)
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt, math.MaxInt)
}
