package leetcode

var dir = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func dfs(board [][]byte, x int, y int) {
	if board[x][y] == 'T' || board[x][y] == 'X' {
		return
	}
	board[x][y] = 'T'
	var m, n = len(board), len(board[0])
	for _, d := range dir {
		tx, ty := x+d[0], y+d[1]
		if tx < 0 || tx >= m || ty < 0 || ty >= n {
			continue
		}
		dfs(board, tx, ty)
	}
}

func solve(board [][]byte) {
	var m, n = len(board), len(board[0])
	for i := 0; i < m; i++ {
		dfs(board, i, 0)
		dfs(board, i, n-1)
	}
	for j := 1; j < n-1; j++ {
		dfs(board, 0, j)
		dfs(board, m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'T' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
