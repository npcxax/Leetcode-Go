package leetcode

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

// my solution
func construct(grid [][]int) *Node {
	var (
		n     = len(grid)
		check func(x, y, length int) bool
		build func(x, y, length int) *Node
	)

	check = func(x, y, length int) bool {
		flag := (grid[x][y] == 1)
		for i := x; i < x+length; i++ {
			for j := y; j < y+length; j++ {
				if flag != (grid[i][j] == 1) {
					return false
				}
			}
		}
		return true
	}

	build = func(x, y, length int) *Node {
		if length < 1 {
			return nil
		}
		newNode := &Node{}
		if length >= 2 {
			flag := true
			if !check(x, y, length) {
				flag = false
				length /= 2
				newNode.TopLeft = build(x, y, length)
				newNode.TopRight = build(x, y+length, length)
				newNode.BottomLeft = build(x+length, y, length)
				newNode.BottomRight = build(x+length, y+length, length)
			}
			if !flag {
				newNode.IsLeaf = false
				newNode.Val = false
			} else {
				newNode.IsLeaf = true
				newNode.Val = (grid[x][y] == 1)
			}
		} else {
			newNode.IsLeaf = true
			newNode.Val = (grid[x][y] == 1)
		}
		return newNode
	}

	return build(0, 0, n)
}

// official
func construct(grid [][]int) *Node {
	var dfs func(rows [][]int, c0, c1 int) *Node
	dfs = func(rows [][]int, c0, c1 int) *Node {
		for _, row := range rows {
			for _, v := range row[c0:c1] {
				if v != rows[0][c0] {
					rMid, cMid := len(rows)>>1, (c0+c1)>>1
					return &Node{
						IsLeaf:      false,
						TopLeft:     dfs(rows[:rMid], c0, cMid),
						TopRight:    dfs(rows[:rMid], cMid, c1),
						BottomLeft:  dfs(rows[rMid:], c0, cMid),
						BottomRight: dfs(rows[rMid:], cMid, c1),
					}
				}
			}
		}
		return &Node{Val: rows[0][c0] == 1, IsLeaf: true}
	}

	return dfs(grid, 0, len(grid))
}

// prefix optimize
func construct(grid [][]int) *Node {
	var (
		n   = len(grid)
		pre = make([][]int, n+1)
		dfs func(r0, r1, c0, c1 int) *Node
	)
	pre[0] = make([]int, n+1)
	for i, row := range grid {
		pre[i+1] = make([]int, n+1)
		for j, v := range row {
			pre[i+1][j+1] = pre[i+1][j] + pre[i][j+1] - pre[i][j] + v
		}
	}
	dfs = func(r0, r1, c0, c1 int) *Node {
		total := pre[r1][c1] - pre[r1][c0] - pre[r0][c1] + pre[r0][c0]
		if total == 0 {
			return &Node{IsLeaf: true, Val: false}
		}
		if total == (r1-r0)*(c1-c0) {
			return &Node{IsLeaf: true, Val: true}
		}
		rMid, cMid := (r0+r1)>>1, (c0+c1)>>1
		return &Node{
			IsLeaf:      false,
			Val:         true,
			TopLeft:     dfs(r0, rMid, c0, cMid),
			TopRight:    dfs(r0, rMid, cMid, c1),
			BottomLeft:  dfs(rMid, r1, c0, cMid),
			BottomRight: dfs(rMid, r1, cMid, c1),
		}
	}

	return dfs(0, n, 0, n)
}
