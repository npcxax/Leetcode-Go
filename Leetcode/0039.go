package leetcode

func combinationSum(candidates []int, target int) [][]int {
	var (
		result    = make([][]int, 0)
		backtrace func(idx, total int, nums []int)
		m         = len(candidates)
	)

	backtrace = func(idx, total int, nums []int) {
		if total == target {
			result = append(result, append([]int(nil), nums...))
			return
		}
		if idx == m || total > target {
			return
		}

		total += candidates[idx]
		nums = append(nums, candidates[idx])

		backtrace(idx, total, nums)

		total -= candidates[idx]
		nums = nums[:len(nums)-1]
		backtrace(idx+1, total, nums)
	}

	backtrace(0, 0, []int{})

	return result
}
