package leetcode

func exist(board [][]byte, word string) bool {
	var (
		m, n   = len(board), len(board[0])
		result = false
		dfs    func(x, y, idx int) bool
		dir    = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	)

	dfs = func(x, y, idx int) bool {
		if board[x][y] != word[idx] {
			return false
		}
		if idx == len(word)-1 {
			return true
		}
		c := board[x][y]
		board[x][y] = '.'
		t := false
		for _, d := range dir {
			tx, ty := x+d[0], y+d[1]
			if tx < 0 || tx >= m || ty < 0 || ty >= n {
				continue
			}
			if board[tx][ty] == '.' {
				continue
			}
			t = t || dfs(tx, ty, idx+1)
		}
		board[x][y] = c
		return t
	}
	for i, row := range board {
		for j, ch := range row {
			if ch == word[0] {
				result = result || dfs(i, j, 0)
			}
		}
	}
	return result
}
