package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// my solution
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stk := make([]*TreeNode, 0)
	count := 0
	stk = append(stk, root)

	for len(stk) > 0 {
		sz := len(stk)
		for i := 0; i < sz; i++ {
			if stk[i].Left != nil {
				stk = append(stk, stk[i].Left)
			}
			if stk[i].Right != nil {
				stk = append(stk, stk[i].Right)
			}
			count++
		}
		stk = stk[sz:]
	}

	return count
}

// official
func countNodes(root *TreeNode) int {
	p := root
	h := 0.0
	for p != nil {
		p = p.Left
		h++
	}
	left, right := int(math.Pow(2, h)), int(math.Pow(2, h+1))-1
	for left < right {
		mid := (left + right) >> 1
		p = root
		for mid > 0 && p != nil {
			if mid&1 == 0 {
				p = p.Left
			} else {
				p = p.Right
			}
			mid = mid >> 1
		}
		if p == nil {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
