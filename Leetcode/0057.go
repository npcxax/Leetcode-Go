package leetcode

// my solution
func insert(intervals [][]int, newInterval []int) [][]int {
	var n = len(intervals)
	if n == 0 {
		return [][]int{newInterval}
	}

	for i := 0; i <= n; i++ {
		if i == n && newInterval[0] < intervals[i][0] {
			intervals = append(intervals[:i],
				append([][]int{newInterval}, intervals[i:]...)...)
			break
		}
	}

	var start, end = intervals[0][0], intervals[0][1]
	var result = make([][]int, 0)
	for idx, interval := range intervals {
		if idx == 0 {
			continue
		}
		if end < interval[0] {
			result = append(result, []int{start, end})
			start = interval[0]
			end = interval[1]
		} else {
			if end < interval[1] {
				end = interval[1]
			}
		}
	}

	result = append(result, []int{start, end})

	return result
}

// official
func insert(intervals [][]int, newInterval []int) [][]int {
	var left, right = newInterval[0], newInterval[1]
	var result = make([][]int, 0)
	var merge = false

	for _, interval := range intervals {
		if interval[0] > right {
			// no intersection
			if !merge {
				result = append(result, []int{left, right})
				merge = true
			}
			result = append(result, interval)
		} else if interval[1] < left {
			result = append(result, interval)
		} else {
			if left > interval[0] {
				left = interval[0]
			}
			if right < interval[1] {
				right = interval[1]
			}
		}
	}
	if !merge {
		result = append(result, []int{left, right})
		merge = true
	}

	return result
}
