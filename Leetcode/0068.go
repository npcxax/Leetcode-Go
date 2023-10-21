package leetcode

import "strings"

// my solution
func addSpace(str string, maxWidth int, wordCount int) string {
	if wordCount == 1 {
		for len(str) < maxWidth {
			str += " "
		}
	} else {
		var totalSpace = maxWidth - len(str)
		var count = 0
		b := []byte(str)
		for i := 0; (i < (totalSpace+wordCount-2)/(wordCount-1)) && count != totalSpace; i++ { // bound up
			var countWord = 0
			j := 0
			for countWord < wordCount-1 && count != totalSpace {
				for b[j] != ' ' {
					j++
				}
				t := b[:j]
				t = append(t, ' ')
				t = append(t, b[j:]...)
				b = t
				for b[j] == ' ' {
					j++
				}
				countWord++
				count++
			}
		}
		str = string(b)
	}
	return str
}

func fullJustify(words []string, maxWidth int) []string {
	var n = len(words)
	var result = make([]string, 0)
	var currentString = ""
	var wordCount = 0
	for i := 0; i < n; {
		if wordCount == 0 && len(words[i]) == maxWidth {
			result = append(result, words[i])
			i++
			continue
		}
		if len(words[i])+len(currentString)+1 > maxWidth {
			// fmt.Printf("i=%d,word = %s\n", i, words[i])
			result = append(result, addSpace(currentString, maxWidth, wordCount))
			currentString = ""
			wordCount = 0
			continue
		}
		if len(currentString) != 0 {
			currentString += " "
		}
		currentString += words[i]
		i++
		wordCount++
	}

	if len(currentString) != 0 {
		result = append(result, addSpace(currentString, maxWidth, 1))
	}

	return result
}

// official
func blank(n int) string {
	return strings.Repeat(" ", n)
}

func fullJustify(words []string, maxWidth int) []string {
	var left, m = 0, len(words)
	var result = make([]string, 0)
	for {
		right := left
		sumLen := 0
		for right < m && sumLen+len(words[right])+right-left <= maxWidth {
			sumLen += len(words[right])
			right++
		}
		totalWords := right - left
		totalSpaces := maxWidth - sumLen
		// last row
		if right == m {
			result = append(result, strings.Join(words[left:right], " ")+blank(totalSpaces-totalWords+1))
			return result
		}
		if totalWords == 1 {
			result = append(result, words[left]+blank(totalSpaces))
		} else {
			avgSpaces := totalSpaces / (totalWords - 1)
			extraSpaces := totalSpaces % (totalWords - 1)
			s1 := strings.Join(words[left:left+extraSpaces+1], blank(avgSpaces+1))
			s2 := strings.Join(words[left+extraSpaces+1:right], blank(avgSpaces))
			result = append(result, s1+blank(avgSpaces)+s2)
		}

		left = right
	}
}
