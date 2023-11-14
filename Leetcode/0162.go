package leetcode

import "math"

// violence
func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			if nums[i] > nums[i+1] {
				return i
			}
		} else if i == n-1 {
			if nums[i] > nums[i-1] {
				return i
			}
		} else {
			if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
				return i
			}
		}
	}
	return -1
}

// binary search
func findPeakElement(nums []int) int {
	n := len(nums)
	left, right := 0, n-1

	var get = func(i int) int {
		if i == -1 || i == n {
			return math.MinInt
		}
		return nums[i]
	}

	for left <= right {
		mid := left + (right-left)>>1
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
