package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func build(nums []int, left int, right int) *TreeNode {
	if left <= right {
		mid := (left + right) >> 1
		node := &TreeNode{
			Val: nums[mid],
		}
		node.Left = build(nums, left, mid-1)
		node.Right = build(nums, mid+1, right)
		return node
	}
	return nil
}

func sortedArrayToBST(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}
