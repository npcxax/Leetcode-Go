package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, curSum int, targetSum int) bool {
	if node == nil {
		return false
	}
	curSum += node.Val
	if curSum == targetSum && node.Left == nil && node.Right == nil {
		return true
	}
	return dfs(node.Left, curSum, targetSum) || dfs(node.Right, curSum, targetSum)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs(root, 0, targetSum)
}
