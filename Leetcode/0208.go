package leetcode

// my solution
type Node struct {
	word  string
	nodes [26]*Node
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{root: &Node{}}
}

func (this *Trie) Insert(word string) {
	p := this.root
	for _, ch := range word {
		if p.nodes[ch-'a'] == nil {
			p.nodes[ch-'a'] = &Node{}
		}
		p = p.nodes[ch-'a']
	}
	p.word = word
}

func (this *Trie) Search(word string) bool {
	p := this.root
	for _, ch := range word {
		if p == nil {
			return false
		}
		if p.nodes[ch-'a'] == nil {
			return false
		}
		p = p.nodes[ch-'a']
	}
	return p.word == word
}

func (this *Trie) StartsWith(prefix string) bool {
	p := this.root
	for _, ch := range prefix {
		if p == nil {
			return false
		}
		if p.nodes[ch-'a'] == nil {
			return false
		}
		p = p.nodes[ch-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

// official
type Trie struct {
	isEnd bool
	nodes [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) searchPrefix(prefix string) *Trie {
	for _, ch := range prefix {
		ch -= 'a'
		if this.nodes[ch] == nil {
			return nil
		}
		this = this.nodes[ch]
	}
	return this
}

func (this *Trie) Insert(word string) {
	for _, ch := range word {
		ch -= 'a'
		if this.nodes[ch] == nil {
			this.nodes[ch] = &Trie{}
		}
		this = this.nodes[ch]
	}
	this.isEnd = true
}

func (this *Trie) Search(word string) bool {
	t := this.searchPrefix(word)
	return t != nil && t.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.searchPrefix(prefix) != nil
}
