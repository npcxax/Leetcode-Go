package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// my solution
func inorder(node *TreeNode, ans *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, ans)
	*ans = append(*ans, node.Val)
	inorder(node.Right, ans)
}

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	inorder(root, &ans)
	return ans
}

// official
func inorderTraversal(root *TreeNode) []int {
	var inorder func(*TreeNode)
	ans := make([]int, 0)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		ans = append(ans, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return ans
}

// iterator
func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	stk := make([]*TreeNode, 0)
	p := root
	for p != nil || len(stk) > 0 {
		for p != nil {
			stk = append(stk, p)
			p = p.Left
		}
		p = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		ans = append(ans, p.Val)
		p = p.Right
	}

	return ans
}
