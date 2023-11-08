package leetcode

func minMutation(startGene string, endGene string, bank []string) int {
	var (
		n        = len(bank)
		graph    = make([][]int, n+1)
		endIndex = -1
		isDif1   func(a, b string) bool
		stk      = make([]int, 0)
		visited  = make([]bool, n+1)
		count    = 0
	)
	if startGene == endGene {
		return 0
	}

	isDif1 = func(a, b string) bool {
		count := 0
		for i := range a {
			if a[i] != b[i] {
				count++
			}
		}
		return count == 1
	}

	// build graph
	for i, a := range bank {
		if a == endGene {
			endIndex = i
		}
		for j := i + 1; j < n; j++ {
			if isDif1(a, bank[j]) {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
		if isDif1(startGene, a) {
			graph[n] = append(graph[n], i)
			graph[i] = append(graph[i], n)
		}
	}

	// bfs
	stk = append(stk, n)
	visited[n] = true
	for len(stk) > 0 {
		sz := len(stk)
		for i := 0; i < sz; i++ {
			s := stk[i]
			for _, edge := range graph[s] {
				if edge == endIndex {
					return count + 1
				}
				if !visited[edge] {
					visited[edge] = true
					stk = append(stk, edge)
				}
			}
		}
		stk = stk[sz:]
		count++
	}

	return -1
}
