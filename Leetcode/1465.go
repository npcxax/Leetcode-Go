package leetcode

import "sort"

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	var maxCol, maxRow = 0, 0
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	maxCol = getMax(h, horizontalCuts)
	maxRow = getMax(w, verticalCuts)

	return (maxCol * maxRow) % (1e9 + 7)
}

func getMax(maxH int, nums []int) int {
	if len(nums) == 1 {
		return max(nums[0], maxH-nums[0])
	}
	var maxNum = nums[0]
	var n = len(nums)
	for i := 1; i < n; i++ {
		maxNum = max(maxNum, nums[i]-nums[i-1])
		if i == n-1 {
			maxNum = max(maxNum, maxH-nums[i])
		}
	}
	return maxNum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
