package leetcode

// my solution
func setZeroes(matrix [][]int) {
	var mRow, mCol = make(map[int]int), make(map[int]int)
	var m, n = len(matrix), len(matrix[0])

	for i, row := range matrix {
		for j, num := range row {
			if num == 0 {
				mRow[i]++
				mCol[j]++
			}
		}
	}
	for k := range mRow {
		for j := 0; j < n; j++ {
			matrix[k][j] = 0
		}
	}
	for k := range mCol {
		for i := 0; i < m; i++ {
			matrix[i][k] = 0
		}
	}
}

// official
func setZeroes(matrix [][]int) {
	var m, n = len(matrix), len(matrix[0])

	var flagOfRow, flagOfCol = false, false
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			flagOfRow = true
		}
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			flagOfCol = true
		}
	}

	// use the first row and column to mark
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
	if flagOfCol {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
	if flagOfRow {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
}

// official
func setZeroes(matrix [][]int) {
	var m, n = len(matrix), len(matrix[0])
	var flagOfCol = false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			flagOfCol = true
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if flagOfCol {
			matrix[i][0] = 0
		}
	}
}
