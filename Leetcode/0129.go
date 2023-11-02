package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, cur int) int {
	total := 0
	cur = cur*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return cur
	}
	if node.Left != nil {
		total += dfs(node.Left, cur)
	}
	if node.Right != nil {
		total += dfs(node.Right, cur)
	}
	return total
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}
