package leetcode

// TopSort
// bfs
func canFinish(numCourses int, prerequisites [][]int) bool {
	var graph = make([][]int, numCourses)
	var count = make([]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		count[p[0]]++
	}
	// Find the node with indegree 0
	var stk = make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if count[i] == 0 {
			stk = append(stk, i)
		}
	}

	var result = make([]int, 0)
	for len(stk) > 0 {
		n := stk[0]
		stk = stk[1:]
		result = append(result, n)
		for _, edge := range graph[n] {
			count[edge]--
			if count[edge] == 0 {
				stk = append(stk, edge)
			}
		}
	}
	return len(result) == numCourses
}

// dfs
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		valid   = true
		dfs     func(int)
		result  = make([]int, 0)
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

	for _, p := range prerequisites {
		edges[p[0]] = append(edges[p[0]], p[1])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}

	return valid
}
