package leetcode

// my solution
func ladderLength(beginWord string, endWord string, wordList []string) int {
	var (
		n        = len(wordList)
		graph    = make([][]int, n+1)
		isDif1   func(a, b string) bool
		endIndex = -1
		visited  = make([]bool, n+1)
		stk      = make([]int, 0)
		count    = 1
	)
	isDif1 = func(a, b string) bool {
		count := 0
		for i := range a {
			if a[i] != b[i] {
				count++
			}
		}
		return count == 1
	}

	// build map
	for i := 0; i < n; i++ {
		if wordList[i] == endWord {
			endIndex = i
		}
		for j := i + 1; j < n; j++ {
			if isDif1(wordList[i], wordList[j]) {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
		if isDif1(beginWord, wordList[i]) {
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
			for _, nxt := range graph[s] {
				if nxt == endIndex {
					return count + 1
				}
				if !visited[nxt] {
					visited[nxt] = true
					stk = append(stk, nxt)
				}
			}
		}
		stk = stk[sz:]
		count++
	}

	return 0
}

// graph optimization
func ladderLength(beginWord string, endWord string, wordList []string) int {
	var (
		hashmap  = make(map[string]int)
		graph    = make([][]int, 0)
		endIndex = -1
		addEdge  func(s string)
		addWord  func(s string) int
		stk      = make([]int, 0)
		count    = 1
		visited  []bool
	)

	addWord = func(s string) int {
		id, ok := hashmap[s]
		if !ok {
			hashmap[s] = len(graph)
			id = len(graph)
			graph = append(graph, make([]int, 0))
		}
		return id
	}

	addEdge = func(s string) {
		id1 := addWord(s)
		if s == endWord {
			endIndex = id1
		}
		b := []byte(s)
		for i, c := range b {
			b[i] = '*'
			id2 := addWord(string(b))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			b[i] = c
		}
	}

	for _, word := range wordList {
		addEdge(word)
	}

	if endIndex == -1 {
		return 0
	}

	addEdge(beginWord)
	stk = append(stk, hashmap[beginWord])

	visited = make([]bool, len(graph))
	visited[hashmap[beginWord]] = true

	for len(stk) > 0 {
		sz := len(stk)
		for i := 0; i < sz; i++ {
			v := stk[i]
			if v == endIndex {
				return count>>1 + 1
			}
			for _, nxt := range graph[v] {
				if !visited[nxt] {
					visited[nxt] = true
					stk = append(stk, nxt)
				}
			}
		}
		stk = stk[sz:]
		count++
	}

	return 0
}
