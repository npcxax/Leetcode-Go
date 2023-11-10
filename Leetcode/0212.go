package leetcode

// my solution
// type Trie struct {
// 	children [26]*Trie
// }

// func CreateTrie() *Trie {
// 	return &Trie{}
// }

// func (this *Trie) Insert(word string) {
// 	for _, ch := range word {
// 		ch -= 'a'
// 		if this.children[ch] == nil {
// 			this.children[ch] = &Trie{}
// 		}
// 		this = this.children[ch]
// 	}
// }

// func (this *Trie) Search(word string) bool {
// 	for _, ch := range word {
// 		ch -= 'a'
// 		if this.children[ch] == nil {
// 			return false
// 		}
// 		this = this.children[ch]
// 	}
// 	return true
// }

// func findWords(board [][]byte, words []string) []string {
// 	var (
// 		m, n     = len(board), len(board[0])
// 		dir      = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
// 		trie     = CreateTrie()
// 		maxDepth = 0
// 		result   = make([]string, 0)
// 		dfs      func(x, y, cur int, b []byte)
// 	)
// 	// dfs
// 	dfs = func(x, y, cur int, b []byte) {
// 		if cur == maxDepth {
// 			return
// 		}
// 		b = append(b, board[x][y])
// 		c := board[x][y]
// 		board[x][y] = '#'
// 		for _, d := range dir {
// 			tx, ty := x+d[0], y+d[1]
// 			if tx < 0 || tx >= m || ty < 0 || ty >= n || board[tx][ty] == '#' {
// 				continue
// 			}
// 			dfs(tx, ty, cur+1, b)
// 		}
// 		trie.Insert(string(b))
// 		board[x][y] = c
// 	}
// 	// find the longest word
// 	for _, word := range words {
// 		if len(word) > maxDepth {
// 			maxDepth = len(word)
// 		}
// 	}

// 	// build trie
// 	for i, row := range board {
// 		for j := range row {
// 			dfs(i, j, 0, []byte{})
// 		}
// 	}

// 	// search word
// 	for _, word := range words {
// 		if trie.Search(word) {
// 			result = append(result, word)
// 		}
// 	}

// 	return result
// }

// optimize
type Trie struct {
	word     string
	children [26]*Trie
}

func CreateTrie() *Trie {
	return &Trie{}
}

func (this *Trie) Insert(word string) {
	t := this
	for _, ch := range word {
		ch -= 'a'
		if t.children[ch] == nil {
			t.children[ch] = &Trie{}
		}
		t = t.children[ch]
	}
	t.word = word
}

func findWords(board [][]byte, words []string) []string {
	var (
		m, n   = len(board), len(board[0])
		dir    = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		result = make([]string, 0)
		dfs    func(x, y int, trie *Trie)
		trie   = CreateTrie()
	)
	// dfs
	dfs = func(x, y int, trie *Trie) {
		ch := board[x][y] - 'a'
		trie = trie.children[ch]
		if trie == nil {
			return
		}
		if trie.word != "" {
			result = append(result, trie.word)
			trie.word = ""
		}

		board[x][y] = '#'
		for _, d := range dir {
			tx, ty := x+d[0], y+d[1]
			if tx < 0 || tx >= m || ty < 0 || ty >= n ||
				board[tx][ty] == '#' {
				continue
			}
			dfs(tx, ty, trie)
		}
		board[x][y] = ch + 'a'
	}

	// build trie
	for _, word := range words {
		trie.Insert(word)
	}

	// search
	for i, row := range board {
		for j := range row {
			dfs(i, j, trie)
		}
	}

	return result
}
