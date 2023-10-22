package leetcode

func minSubArrayLen(target int, nums []int) int {
	var n = len(nums)
	var left, right = 0, 0
	var curTotal = 0
	var result = 0x7fffffff

	for right < n {
		curTotal += nums[right]
		for curTotal >= target {
			if right-left+1 < result {
				result = right - left + 1
			}
			curTotal -= nums[left]
			left++
		}
		right++

	}
	if result == 0x7fffffff {
		return 0
	}
	return result
}
