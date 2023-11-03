package leetcode

import (
	"math"
)

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
	height := 0
	for p != nil {
		p = p.Left
		height++
	}
	height -= 1
	var exist = func(node *TreeNode, height int, num int) bool {
		b := 1 << (height - 1)
		p := node
		for b > 0 && p != nil {
			if num&b == 0 {
				p = p.Left
			} else {
				p = p.Right
			}
			b >>= 1
		}
		return p != nil
	}
	var left, right = int(math.Pow(2, float64(height))), int(math.Pow(2, float64(height+1)) - 1)
	for left < right {
		mid := (right-left+1)>>1 + left // up bound
		if exist(root, height, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

// other
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	heightL := 0
	p := root
	for p != nil {
		p = p.Left
		heightL++
	}
	p = root
	heightR := 0
	for p != nil {
		p = p.Right
		heightR++
	}
	if heightL == heightR {
		return int(math.Pow(2, float64(heightL)) - 1)
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
