package leetcode

// same row, same column, same diagonal
func totalNQueens(n int) int {
	var (
		result               = 0
		column               = make(map[int]bool)
		diagonal1, diagonal2 = make(map[int]bool), make(map[int]bool)
		backtrace            func(row int)
	)

	backtrace = func(row int) {
		if row == n {
			result++
			return
		}
		for i := 0; i < n; i++ {
			if column[i] {
				continue
			}
			if diagonal1[row-i] {
				continue
			}
			if diagonal2[row+i] {
				continue
			}
			column[i] = true
			diagonal1[row-i] = true
			diagonal2[row+i] = true
			backtrace(row + 1)
			column[i] = false
			diagonal1[row-i] = false
			diagonal2[row+i] = false
		}
	}
	backtrace(0)

	return result
}
