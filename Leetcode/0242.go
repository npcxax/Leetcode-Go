package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sCount := make([]int, 26)
	for _, ch := range s {
		sCount[ch-'a']++
	}
	for _, ch := range t {
		sCount[ch-'a']--
		if sCount[ch-'a'] < 0 {
			return false
		}
	}
	return true
}

// unicode
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sCount := make(map[rune]int)
	for _, ch := range s {
		sCount[ch]++
	}
	for _, ch := range t {
		sCount[ch]--
		if sCount[ch] < 0 {
			return false
		}
	}
	return true
}
