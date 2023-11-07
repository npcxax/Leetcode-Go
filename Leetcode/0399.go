package leetcode

// 1.bfs
type Edge struct {
	To     string
	Weight float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	var (
		hashmap = make(map[string]int)
		result  = make([]float64, len(queries))
		graph   [][]*Edge
		bfs     func(start, end int) float64
		id      = 0
	)

	for _, eq := range equations {
		if _, ok := hashmap[eq[0]]; !ok {
			hashmap[eq[0]] = id
			id++
		}
		if _, ok := hashmap[eq[1]]; !ok {
			hashmap[eq[1]] = id
			id++
		}
	}

	bfs = func(start, end int) float64 {
		var stk = make([]int, 0)
		var ratio = make([]float64, id)
		ratio[start] = 1
		stk = append(stk, start)

		for len(stk) > 0 {
			start = stk[0]
			stk = stk[1:]
			for _, edge := range graph[start] {
				if hashmap[edge.To] == end {
					return edge.Weight * ratio[start]
				}
				if w := hashmap[edge.To]; ratio[w] == 0 {
					ratio[w] = ratio[start] * edge.Weight
					stk = append(stk, w)
				}
			}
		}
		return -1
	}

	graph = make([][]*Edge, id)

	for idx, eq := range equations {
		graph[hashmap[eq[0]]] = append(graph[hashmap[eq[0]]], &Edge{
			To:     eq[1],
			Weight: values[idx],
		})
		graph[hashmap[eq[1]]] = append(graph[hashmap[eq[1]]], &Edge{
			To:     eq[0],
			Weight: 1 / values[idx],
		})
	}

	for idx, q := range queries {
		if _, ok := hashmap[q[0]]; !ok {
			result[idx] = -1
			continue
		}
		if _, ok := hashmap[q[1]]; !ok {
			result[idx] = -1
			continue
		}
		result[idx] = bfs(hashmap[q[0]], hashmap[q[1]])
	}

	return result
}

// Floyd
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	var (
		result  = make([]float64, len(queries))
		hashmap = make(map[string]int)
		graph   [][]float64
		id      = 0
	)

	for _, eq := range equations {
		if _, ok := hashmap[eq[0]]; !ok {
			hashmap[eq[0]] = id
			id++
		}
		if _, ok := hashmap[eq[1]]; !ok {
			hashmap[eq[1]] = id
			id++
		}
	}

	for i := 0; i < id; i++ {
		graph = append(graph, make([]float64, id))
	}

	for idx, eq := range equations {
		graph[hashmap[eq[0]]][hashmap[eq[1]]] = values[idx]
		graph[hashmap[eq[1]]][hashmap[eq[0]]] = 1 / values[idx]
	}

	for i := range graph {
		for j := range graph {
			for k := range graph {
				if graph[i][j] != 0 && graph[j][k] != 0 {
					graph[i][k] = graph[i][j] * graph[j][k]
					graph[k][i] = 1 / graph[i][k]
				}
			}
		}
	}

	for idx, q := range queries {
		if _, ok := hashmap[q[0]]; !ok {
			result[idx] = -1
			continue
		}
		if _, ok := hashmap[q[1]]; !ok {
			result[idx] = -1
			continue
		}
		if graph[hashmap[q[0]]][hashmap[q[1]]] == 0 {
			result[idx] = -1
		} else {
			result[idx] = graph[hashmap[q[0]]][hashmap[q[1]]]
		}

	}

	return result
}

// unionfind with weight
type UnionFind struct {
	parent []int
	weight []float64
}

func createUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		weight: make([]float64, n),
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.weight[i] = 1
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if x != uf.parent[x] {
		origin := uf.parent[x]
		uf.parent[x] = uf.Find(uf.parent[x])
		uf.weight[x] *= uf.weight[origin]
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int, value float64) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX == rootY {
		return
	}
	uf.parent[rootX] = rootY
	uf.weight[rootX] = uf.weight[y] * value / uf.weight[x]
}

func (uf *UnionFind) IsConnected(x, y int) float64 {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		return -1
	}
	return uf.weight[x] / uf.weight[y]
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	var (
		uf      *UnionFind = createUnionFind(2 * len(equations))
		hashmap            = make(map[string]int)
		id                 = 0
		result             = make([]float64, len(queries))
	)

	for idx, eq := range equations {
		if _, ok := hashmap[eq[0]]; !ok {
			hashmap[eq[0]] = id
			id++
		}
		if _, ok := hashmap[eq[1]]; !ok {
			hashmap[eq[1]] = id
			id++
		}
		uf.Union(hashmap[eq[0]], hashmap[eq[1]], values[idx])
	}

	for idx, q := range queries {
		if _, ok := hashmap[q[0]]; !ok {
			result[idx] = -1
			continue
		}
		if _, ok := hashmap[q[1]]; !ok {
			result[idx] = -1
			continue
		}
		result[idx] = uf.IsConnected(hashmap[q[0]], hashmap[q[1]])
	}

	return result
}
