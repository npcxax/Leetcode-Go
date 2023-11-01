package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// my solution
func isSymmetric(root *TreeNode) bool {
	var stk = make([]*TreeNode, 0)
	stk = append(stk, root)

	for len(stk) > 0 {
		sz := len(stk)
		var arr = make([]int, 0)
		for i := 0; i < sz; i++ {
			root = stk[i]
			if root != nil {
				arr = append(arr, root.Val)
				stk = append(stk, root.Left)
				stk = append(stk, root.Right)
			} else {
				arr = append(arr, -101)
			}
		}
		stk = stk[sz:]
		var left, right = 0, len(arr) - 1
		for left < right {
			if arr[left] != arr[right] {
				return false
			}
			left++
			right--
		}
	}

	return true
}

// official
func isSymmetric(root *TreeNode) bool {
	var stk = make([]*TreeNode, 0)
	u, v := root, root
	stk = append(stk, u)
	stk = append(stk, v)
	for len(stk) > 0 {
		u, v := stk[0], stk[1]
		stk = stk[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		stk = append(stk, u.Left)
		stk = append(stk, v.Right)
		stk = append(stk, u.Right)
		stk = append(stk, v.Left)
	}
	return true
}

func dfs(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	return dfs(root, root)
}
