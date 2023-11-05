package leetcode

type Node struct {
	Val       int
	Neighbors []*Node
}

// dfs
func cloneGraph(node *Node) *Node {
	var hashmap = make(map[*Node]*Node, 0)
	var dfs func(*Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if v, ok := hashmap[node]; ok {
			return v
		}
		var newNode = &Node{
			Val:       node.Val,
			Neighbors: nil,
		}
		hashmap[node] = newNode
		for _, neighbour := range node.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, dfs(neighbour))
		}
		return newNode
	}
	return dfs(node)
}

// bfs
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	var visited = make(map[*Node]*Node)
	var stk = make([]*Node, 0)
	stk = append(stk, node)
	var root = &Node{
		Val: node.Val,
	}
	visited[node] = root

	for len(stk) > 0 {
		sz := len(stk)
		for i := 0; i < sz; i++ {
			node = stk[i]
			for _, neighbour := range node.Neighbors {
				if v, ok := visited[neighbour]; ok {
					visited[node].Neighbors = append(visited[node].Neighbors, v)
				} else {
					var tnode = &Node{
						Val: neighbour.Val,
					}
					visited[neighbour] = tnode
					stk = append(stk, neighbour)
					visited[node].Neighbors = append(visited[node].Neighbors, tnode)
				}
			}
		}
		stk = stk[sz:]
	}

	return root
}

// official
func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	var visited = make(map[*Node]*Node)
	var stk = make([]*Node, 0)
	visited[node] = &Node{
		Val: node.Val,
	}
	stk = append(stk, node)

	for len(stk) > 0 {
		n := stk[0]
		stk = stk[1:]
		for _, neighbour := range n.Neighbors {
			if _, ok := visited[neighbour]; !ok {
				visited[neighbour] = &Node{
					Val: neighbour.Val,
				}
				stk = append(stk, neighbour)
			}
			visited[n].Neighbors = append(visited[n].Neighbors, visited[neighbour])
		}
	}

	return visited[node]
}
