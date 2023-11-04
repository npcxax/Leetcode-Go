package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// my solution
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var stk = make([]*TreeNode, 0)
	var res = make([][]int, 0)
	var flag = true
	stk = append(stk, root)

	for len(stk) > 0 {
		sz := len(stk)
		t := make([]int, 0)
		if flag {
			for i := 0; i < sz; i++ {
				root = stk[i]
				t = append(t, root.Val)
				if root.Left != nil {
					stk = append(stk, root.Left)
				}
				if root.Right != nil {
					stk = append(stk, root.Right)
				}
			}
		} else {
			for i := sz - 1; i >= 0; i-- {
				root = stk[i]
				t = append(t, root.Val)
				if root.Right != nil {
					stk = append(stk[:sz], append([]*TreeNode{root.Right}, stk[sz:]...)...)
				}
				if root.Left != nil {
					stk = append(stk[:sz], append([]*TreeNode{root.Left}, stk[sz:]...)...)
				}
			}
		}
		flag = !flag
		res = append(res, t)
		stk = stk[sz:]
	}

	return res
}

// reverse
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var stk = make([]*TreeNode, 0)
	var res = make([][]int, 0)
	stk = append(stk, root)
	flag := true

	for len(stk) > 0 {
		sz := len(stk)
		t := make([]int, 0)
		for i := 0; i < sz; i++ {
			root = stk[i]
			t = append(t, root.Val)
			if root.Left != nil {
				stk = append(stk, root.Left)
			}
			if root.Right != nil {
				stk = append(stk, root.Right)
			}
		}
		if !flag {
			left, right := 0, len(t)-1
			for left < right {
				t[left], t[right] = t[right], t[left]
				left++
				right--
			}
		}
		stk = stk[sz:]
		flag = !flag
		res = append(res, t)
	}

	return res
}
