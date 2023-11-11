package leetcode

func isToeplitzMatrix(matrix [][]int) bool {
	var (
		m, n = len(matrix), len(matrix[0])
	)
	for idx := -m + 1; idx < n; idx++ {
		num := -1
		for j := 0; j < m; j++ {
			if idx+j < 0 || idx+j >= n {
				continue
			}
			if num == -1 {
				num = matrix[j][idx+j]
			}
			if num != matrix[j][idx+j] {
				return false
			}
		}
	}

	return true
}
