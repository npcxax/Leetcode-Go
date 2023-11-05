package leetcode

func numIslands(grid [][]byte) int {
	var dfs func(x, y int)
	var m, n = len(grid), len(grid[0])
	var dir = [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	dfs = func(x, y int) {
		grid[x][y] = '2'
		for _, d := range dir {
			tx, ty := x+d[0], y+d[1]
			if tx < 0 || tx >= m || ty < 0 || ty >= n {
				continue
			}
			if grid[tx][ty] == '1' {
				dfs(tx, ty)
			}
		}
	}
	var count = 0
	for i, row := range grid {
		for j, c := range row {
			if c == '1' {
				dfs(i, j)
				count++
			}
		}
	}

	return count
}

// UnionFind
type UnionFind struct {
	parent []int
	count  int
}

func CreateUnion(grid [][]byte) *UnionFind {
	var m, n = len(grid), len(grid[0])
	var uf = &UnionFind{
		parent: make([]int, m*n),
		count:  0,
	}

	for i, row := range grid {
		for j, c := range row {
			if c == '1' {
				uf.parent[i*n+j] = i*n + j
				uf.count++
			} else {
				uf.parent[i*n+j] = -1
			}
		}
	}
	return uf
}

func (uf *UnionFind) find(x int) (int, int) {
	var count = 0
	for uf.parent[x] != -1 && uf.parent[x] != x {
		x = uf.parent[x]
		count += 1
	}
	return x, count
}

func (uf *UnionFind) Union(x, y int) {
	fx, cx := uf.find(x)
	fy, cy := uf.find(y)
	if fx == fy {
		return
	}
	if cx > cy {
		uf.parent[fy] = fx
	} else {
		uf.parent[fx] = fy
	}
	uf.count--
}

func numIslands(grid [][]byte) int {
	var uf = CreateUnion(grid)
	var m, n = len(grid), len(grid[0])
	for i, row := range grid {
		for j, c := range row {
			if c == '1' {
				c = '0'
				if i-1 >= 0 && grid[i-1][j] == '1' {
					uf.Union((i-1)*n+j, i*n+j)
				}
				if i+1 < m && grid[i+1][j] == '1' {
					uf.Union((i+1)*n+j, i*n+j)
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					uf.Union(i*n+j-1, i*n+j)
				}
				if j+1 < n && grid[i][j+1] == '1' {
					uf.Union(i*n+j+1, i*n+j)
				}
			}
		}
	}
	return uf.count
}
