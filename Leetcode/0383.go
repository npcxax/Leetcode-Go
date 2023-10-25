package leetcode

// my solution
func canConstruct(ransomNote string, magazine string) bool {
	var hashmap = make(map[rune]int, 0)
	for _, ch := range ransomNote {
		hashmap[ch]++
	}
	for _, ch := range magazine {
		if _, ok := hashmap[ch]; ok {
			hashmap[ch]--
			if hashmap[ch] == 0 {
				delete(hashmap, ch)
			}
		}
		if len(hashmap) == 0 {
			return true
		}
	}
	return false
}

// official
func canConstruct(ransomNote string, magazine string) bool {
	var hashmap = make([]int, 26)
	for _, ch := range magazine {
		hashmap[ch-'0']++
	}
	for _, ch := range ransomNote {
		hashmap[ch-'0']--
		if hashmap[ch-'0'] < 0 {
			return false
		}
	}
	return true
}
