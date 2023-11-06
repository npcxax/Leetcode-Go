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
