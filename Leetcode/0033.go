package leetcode

func search(nums []int, target int) int {
	var (
		n           = len(nums)
		left, right = 0, n - 1
	)
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else {
			if nums[0] <= nums[mid] {
				if nums[0] <= target && target < nums[mid] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			} else {
				if nums[mid] < target && target <= nums[right] {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
		}
	}

	return -1
}
