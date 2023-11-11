package leetcode

// my solution
func solveNQueens(n int) [][]string {
	var (
		result               = make([][]string, 0)
		graph                = make([][]byte, n)
		backtrace            func(x, y int)
		isValid              func(x, y int) bool
		columns              = make(map[int]bool)
		diagonal1, diagonal2 = make(map[int]bool), make(map[int]bool)
	)

	backtrace = func(x, y int) {
		if !isValid(x, y) {
			return
		}

		if x == n-1 {
			t := make([]string, n)
			for i := 0; i < n; i++ {
				t[i] = string(graph[i])
			}
			result = append(result, t)
			graph[x][y] = '.'
			columns[y] = false
			diagonal1[x-y] = false
			diagonal2[x+y] = false
			return
		}

		for k := 0; k < n; k++ {
			if k == y || k == y-1 || k == y+1 {
				continue
			}
			backtrace(x+1, k)
		}
		graph[x][y] = '.'
		columns[y] = false
		diagonal1[x-y] = false
		diagonal2[x+y] = false
	}

	isValid = func(x, y int) bool {
		if columns[y] || diagonal1[x-y] || diagonal2[x+y] {
			return false
		}
		columns[y] = true
		diagonal1[x-y] = true
		diagonal2[x+y] = true
		graph[x][y] = 'Q'
		return true
	}

	for i := 0; i < n; i++ {
		graph[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			graph[i][j] = '.'
		}
	}

	for i := 0; i < n; i++ {
		backtrace(0, i)
	}

	return result
}

// official
func solveNQueens(n int) [][]string {
	var (
		result               = make([][]string, 0)
		column               = make(map[int]bool)
		diagonal1, diagonal2 = make(map[int]bool), make(map[int]bool)
		generateBoard        func(arr []int)
		backtrace            func(arr []int, row int)
		arr                  = make([]int, n)
	)

	generateBoard = func(arr []int) {
		res := make([]string, n)
		for i := 0; i < n; i++ {
			b := []byte{}
			for j := 0; j < n; j++ {
				if j == arr[i] {
					b = append(b, 'Q')
				} else {
					b = append(b, '.')
				}
			}
			res[i] = string(b)
		}
		result = append(result, res)
	}

	backtrace = func(arr []int, row int) {
		if row == n {
			generateBoard(arr)
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
			arr[row] = i
			column[i] = true
			diagonal1[row-i] = true
			diagonal2[row+i] = true
			backtrace(arr, row+1)
			arr[row] = -1
			column[i] = false
			diagonal1[row-i] = false
			diagonal2[row+i] = false
		}
	}

	for i := 0; i < n; i++ {
		arr[i] = -1
	}

	backtrace(arr, 0)

	return result
}
