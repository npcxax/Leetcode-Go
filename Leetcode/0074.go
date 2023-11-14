package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	var (
		m, n        = len(matrix), len(matrix[0])
		left, right = 0, m*n - 1
	)

	for left <= right {
		mid := left + (right-left)>>1
		tx, ty := mid/n, mid%n
		if matrix[tx][ty] == target {
			return true
		} else if matrix[tx][ty] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
