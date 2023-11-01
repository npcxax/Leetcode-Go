package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	var stk = make([]*TreeNode, 0)
	stk = append(stk, node)
	var count int = 0
	for len(stk) != 0 {
		var length = len(stk)
		for i := 0; i < length; i++ {
			node = stk[i]
			if node.Left != nil {
				stk = append(stk, node.Left)
			}
			if node.Right != nil {
				stk = append(stk, node.Right)
			}
		}
		stk = stk[length:]
		count++
	}
	return count
}

func dfs(node *TreeNode, height int) int {
	if node == nil {
		return height
	}
	return max(dfs(node.Left, height+1), dfs(node.Right, height+1))
}

func maxDepth(root *TreeNode) int {
	return bfs(root)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
