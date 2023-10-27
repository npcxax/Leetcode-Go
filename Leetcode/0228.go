package leetcode

import "fmt"

func summaryRanges(nums []int) []string {
	var result = make([]string, 0)
	var start, end = 0, 0
	var n = len(nums)
	for end < n {
		for end+1 < n && nums[end]+1 == nums[end+1] {
			end++
		}
		if start != end {
			result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[end]))
		} else {
			result = append(result, fmt.Sprintf("%d", nums[start]))
		}
		end++
		start = end
	}

	return result
}
