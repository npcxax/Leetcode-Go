package leetcode

func permute(nums []int) [][]int {
	var (
		m         = len(nums)
		result    = make([][]int, 0)
		backtrace func(num int, arr []int)
		seen      = make(map[int]bool)
	)
	backtrace = func(num int, arr []int) {
		arr = append(arr, num)
		seen[num] = true
		if len(arr) == m {
			result = append(result, arr)
			seen[num] = false
			return
		}
		for _, v := range nums {
			if !seen[v] {
				backtrace(v, arr)
			}
		}
		seen[num] = false
	}

	for _, num := range nums {
		backtrace(num, []int{})
	}

	return result
}
