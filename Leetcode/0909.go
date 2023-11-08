package leetcode

type Point struct {
	x, y int
}

type Status struct {
	id   int
	step int
}

func snakesAndLadders(board [][]int) int {
	var (
		n       = len(board)
		hashmap = make(map[int]*Point)
		id      = 1
		flag    = true
		stk     = make([]*Status, 0)
		visited = make([]bool, n*n)
	)

	// create relationship about id to (x,y)
	for i := n - 1; i >= 0; i-- {
		if flag {
			for j := 0; j < n; j++ {
				hashmap[id] = &Point{x: i, y: j}
				id++
			}
		} else {
			for j := n - 1; j >= 0; j-- {
				hashmap[id] = &Point{x: i, y: j}
				id++
			}
		}
		flag = !flag
	}

	// bfs
	stk = append(stk, &Status{id: 1, step: 0})
	for len(stk) > 0 {
		s := stk[0]
		stk = stk[1:]
		for i := 1; i <= 6; i++ {
			nxt := s.id + i
			if nxt > n*n {
				break
			}
			x, y := hashmap[nxt].x, hashmap[nxt].y
			if board[x][y] > 0 {
				nxt = board[x][y]
			}
			if nxt == n*n {
				return s.step + 1
			}
			if !visited[nxt] {
				visited[nxt] = true
				stk = append(stk, &Status{id: nxt, step: s.step + 1})
			}
		}
	}

	// not found
	return -1
}
