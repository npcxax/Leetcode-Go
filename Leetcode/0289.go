package leetcode

import "fmt"

// my solution
func gameOfLife(board [][]int) {
	var m, n = len(board), len(board[0])
	var copyBoard = make([][]int, m)
	for i := 0; i < m; i++ {
		copyBoard[i] = append(copyBoard[i], board[i]...)
	}

	var check = func(x, y int) int {
		var count = 0
		var dirs = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

		for _, dir := range dirs {
			tx, ty := dir[0]+x, dir[1]+y
			if tx < 0 || tx >= m || ty < 0 || ty >= n {
				continue
			}
			if copyBoard[tx][ty] == 1 {
				count++
			}
		}
		if count < 2 || count > 3 {
			return 0
		}
		if count == 3 {
			return 1
		}
		return copyBoard[x][y]
	}

	for i, row := range board {
		for j := range row {
			board[i][j] = check(i, j)
		}
	}
}

// official
func gameOfLife(board [][]int) {
	// -1 represent 1 to 0, 2 represent 0 to 1
	var m, n = len(board), len(board[0])

	var check = func(x, y int) int {
		var count = 0
		var dirs = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

		for _, dir := range dirs {
			tx, ty := x+dir[0], y+dir[1]
			if tx < 0 || tx >= m || ty < 0 || ty >= n {
				continue
			}
			if board[tx][ty] == 1 || board[tx][ty] == -1 {
				count++
			}
		}
		fmt.Println(board)
		if (count < 2 || count > 3) && board[x][y] == 1 {
			return -1
		}
		if count == 3 && board[x][y] == 0 {
			return 2
		}
		return board[x][y]
	}

	for i, row := range board {
		for j := range row {
			board[i][j] = check(i, j)
		}
	}

	for i, row := range board {
		for j := range row {
			if board[i][j] == -1 {
				board[i][j] = 0
			} else if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}
}
