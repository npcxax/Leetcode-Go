package leetcode

import "sort"

type twoDSlice [][]int

// my solution
func (t twoDSlice) Len() int {
	return len(t)
}

func (t twoDSlice) Less(x, y int) bool {
	if t[x][0] != t[y][0] {
		return t[x][0] < t[y][0]
	}
	return t[x][1] > t[y][1]
}

func (t twoDSlice) Swap(x, y int) {
	t[x], t[y] = t[y], t[x]
}

func merge(intervals [][]int) [][]int {
	sort.Sort(twoDSlice(intervals))
	start, end := intervals[0][0], intervals[0][1]
	result := make([][]int, 0)

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			// new intervals
			result = append(result, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		} else {
			// same interval
			if intervals[i][1] < end {
				continue
			}
			end = intervals[i][1]
		}
	}

	result = append(result, []int{start, end})

	return result
}

// others
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(x, y int) bool {
		if intervals[x][0] != intervals[y][0] {
			return intervals[x][0] < intervals[y][0]
		}
		return intervals[x][1] > intervals[y][1]
	})
	var result = make([][]int, 0)

	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			result = append(result, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		} else {
			if intervals[i][1] < end {
				continue
			}
			end = intervals[i][1]
		}
	}
	result = append(result, []int{start, end})
	return result
}
