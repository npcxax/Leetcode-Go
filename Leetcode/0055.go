package leetcoode

func canJump(nums []int) bool {
	var rightMost = 0
	var n = len(nums)

	for i := 0; i < n; i++ {
		if i <= rightMost {
			rightMost = max(rightMost, i+nums[i])
			if rightMost >= n-1 {
				return true
			}
		} else {
			break
		}
	}

	return false
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
