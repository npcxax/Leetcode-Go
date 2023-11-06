package leetcode

// bfs
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges  = make([][]int, numCourses)
		result = make([]int, 0)
		count  = make([]int, numCourses)
		stk    = make([]int, 0)
	)

	for _, pr := range prerequisites {
		edges[pr[1]] = append(edges[pr[1]], pr[0])
		count[pr[0]]++
	}

	for i := 0; i < numCourses; i++ {
		if count[i] == 0 {
			stk = append(stk, i)
		}
	}

	for len(stk) > 0 {
		n := stk[0]
		stk = stk[1:]
		result = append(result, n)
		for _, edge := range edges[n] {
			count[edge]--
			if count[edge] == 0 {
				stk = append(stk, edge)
			}
		}
	}
	if len(result) != numCourses {
		return []int{}
	}

	return result
}

// dfs
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		result  = make([]int, 0)
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		valid   = true
		dfs     func(n int)
	)

	dfs = func(n int) {
		visited[n] = 1
		for _, edge := range edges[n] {
			if visited[edge] == 0 {
				dfs(edge)
				if !valid {
					return
				}
			} else if visited[edge] == 1 {
				valid = false
				return
			}
		}
		visited[n] = 2
		result = append(result, n)
	}

	for _, pr := range prerequisites {
		edges[pr[0]] = append(edges[pr[0]], pr[1])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}

	if !valid {
		return []int{}
	}

	return result
}
