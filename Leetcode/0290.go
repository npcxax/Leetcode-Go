package leetcode

import "strings"

// my solution
func wordPattern(pattern string, s string) bool {
	pToS := make(map[byte]string)
	sToP := make(map[string]byte)
	count := 0
	start, end := 0, 0
	sLen := len(s)
	pLen := len(pattern)
	for end < sLen && count < pLen {
		for end < sLen && s[end] != ' ' {
			end++
		}
		sValue := s[start:end]
		if v, ok := pToS[pattern[count]]; ok {
			if v != sValue {
				return false
			}
		} else {
			pToS[pattern[count]] = sValue
		}
		if v, ok := sToP[sValue]; ok {
			if v != pattern[count] {
				return false
			}
		} else {
			sToP[sValue] = pattern[count]
		}
		start = end + 1
		end = start
		count++
	}
	return end == sLen+1 && count == pLen
}

// official
func wordPattern(pattern string, s string) bool {
	word2ch := make(map[string]byte)
	ch2word := make(map[byte]string)
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	for i, word := range words {
		ch := pattern[i]
		if (word2ch[word] > 0 && word2ch[word] != ch) || (ch2word[ch] != "" && ch2word[ch] != word) {
			return false
		}
		word2ch[word] = ch
		ch2word[ch] = word
	}
	return true
}
