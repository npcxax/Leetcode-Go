package leetcoode

func solve(nums []int, left int, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func rotate(nums []int, k int) {
	var n = len(nums)
	k %= n

	solve(nums, 0, n-k-1)
	solve(nums, n-k, n-1)
	solve(nums, 0, n-1)
}
