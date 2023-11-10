package leetcode

type WordDictionary struct {
	isEnd    bool
	children [26]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	for _, ch := range word {
		ch -= 'a'
		if this.children[ch] == nil {
			this.children[ch] = &WordDictionary{}
		}
		this = this.children[ch]
	}
	this.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(w *WordDictionary, idx int) bool
	dfs = func(w *WordDictionary, idx int) bool {
		if idx == len(word) {
			return w.isEnd
		}
		flag := false
		if word[idx] == '.' {
			for i := 0; i < 26; i++ {
				if w.children[i] != nil {
					flag = flag || dfs(w.children[i], idx+1)
				}
			}
		} else {
			if w.children[word[idx]-'a'] == nil {
				return false
			}
			flag = flag || dfs(w.children[word[idx]-'a'], idx+1)
		}
		return flag
	}
	return dfs(this, 0)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
