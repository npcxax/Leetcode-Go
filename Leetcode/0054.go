package leetcode

import "fmt"

func spiralOrder(matrix [][]int) []int {
	var m, n = len(matrix), len(matrix[0])
	var result = make([]int, 0)

	var top, bottom, left, right = 0, m - 1, 0, n - 1

	for left <= right && top <= bottom {
		fmt.Println(left)
		for j := left; j <= right; j++ {
			result = append(result, matrix[top][j])
		}
		for i := top + 1; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		if left < right && top < bottom {
			for j := right - 1; j >= left; j-- {
				result = append(result, matrix[bottom][j])
			}
			for i := bottom - 1; i > top; i-- {
				result = append(result, matrix[i][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}

	return result
}
