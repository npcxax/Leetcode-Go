package leetcode

// my solution
func searchRange(nums []int, target int) []int {
	var (
		n           = len(nums)
		left, right = 0, n - 1
	)

	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			l, r := mid, mid
			for l-1 >= 0 && nums[l-1] == nums[mid] {
				l--
			}
			for r+1 < n && nums[r+1] == nums[mid] {
				r++
			}
			return []int{l, r}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return []int{-1, -1}
}

// official
func searchRange(nums []int, target int) []int {
	var (
		n                 = len(nums)
		binarySearch      func(lower bool) int
		leftIdx, rightIdx = 0, 0
	)
	binarySearch = func(lower bool) int {
		left, right := 0, n-1
		ans := n
		for left <= right {
			mid := left + (right-left)>>1
			if nums[mid] > target || (lower && nums[mid] >= target) {
				right = mid - 1
				ans = mid
			} else {
				left = mid + 1
			}
		}
		return ans
	}
	leftIdx, rightIdx = binarySearch(true), binarySearch(false)-1
	if leftIdx <= rightIdx && rightIdx < n && nums[leftIdx] == nums[rightIdx] {
		return []int{leftIdx, rightIdx}
	}

	return []int{-1, -1}
}
