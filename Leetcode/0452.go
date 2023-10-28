package leetcode

import "sort"

func findMinArrowShots(points [][]int) int {
	var count = 1
	sort.Slice(points, func(x, y int) bool {
		if points[x][0] != points[y][0] {
			return points[x][0] < points[y][0]
		}
		return points[x][1] < points[y][1]
	})
	// maximum overlap interval
	var left, right = points[0][1], points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > right {
			count++
			left = points[i][0]
			right = points[i][1]
		} else {
			if left < points[i][0] {
				left = points[i][0]
			}
			if right > points[i][1] {
				right = points[i][1]
			}
		}
	}

	return count
}

// official
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(x, y int) bool {
		return points[x][1] < points[y][1]
	})
	var count = 1
	var right = points[0][1]

	for _, point := range points {
		if point[0] > right {
			right = point[1]
			count++
		}
	}

	return count
}
