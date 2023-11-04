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
func maxPathSum(root *TreeNode) int {
	var ans = math.MinInt

	hashmap := make(map[*TreeNode]int)
	maxGain(root, hashmap)
	dfs(root, hashmap, &ans)
	return ans
}

func maxGain(node *TreeNode, hashmap map[*TreeNode]int) int {
	if node == nil {
		return 0
	}
	maxL, maxR := maxGain(node.Left, hashmap), maxGain(node.Right, hashmap)
	hashmap[node] = node.Val
	if maxL > 0 && maxR > 0 {
		hashmap[node] += max(maxL, maxR)
	} else if maxL > 0 {
		hashmap[node] += maxL
	} else if maxR > 0 {
		hashmap[node] += maxR
	}
	return hashmap[node]
}

func dfs(node *TreeNode, hashmap map[*TreeNode]int, ans *int) {
	if node == nil {
		return
	}
	total := node.Val
	if node.Left != nil && node.Right != nil {
		if hashmap[node.Left] > 0 {
			total += hashmap[node.Left]
		}
		if hashmap[node.Right] > 0 {
			total += hashmap[node.Right]
		}
	} else if node.Left != nil {
		if hashmap[node.Left] > 0 {
			total += hashmap[node.Left]
		}
	} else if node.Right != nil {
		if hashmap[node.Right] > 0 {
			total += hashmap[node.Right]
		}
	}
	if total > *ans {
		*ans = total
	}
	dfs(node.Left, hashmap, ans)
	dfs(node.Right, hashmap, ans)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// official
func maxPathSum(root *TreeNode) int {
	var maxGain func(node *TreeNode) int
	var ans = math.MinInt
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lGain := max(maxGain(node.Left), 0)
		rGain := max(maxGain(node.Right), 0)
		pathSum := lGain + rGain + node.Val
		if pathSum > ans {
			ans = pathSum
		}
		return node.Val + max(lGain, rGain)
	}
	maxGain(root)
	return ans
}
