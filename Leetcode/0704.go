package leetcode

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	mid := 0
	for l <= r {
		mid = (l + r) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
