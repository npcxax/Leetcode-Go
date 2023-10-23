package leetcode

// my solution
func findSubstring(s string, words []string) []int {
	var m = make(map[string]int, 0)

	for _, word := range words { // 1 or more represent not used, 0 represent used
		m[word]++
	}

	var result = make([]int, 0)
	var wordLen = len(words[0])
	var numOfWords = len(words)
	var start, end = 0, numOfWords * wordLen
	var n = len(s)

	for end <= n {
		var tempStart = start
		var tempEnd = start
		var flag = true
		for tempEnd < end {
			tempEnd++
			if (tempEnd-tempStart)%wordLen == 0 {
				if _, ok := m[s[tempStart:tempEnd]]; !ok {
					flag = false
					break
				}
				if m[s[tempStart:tempEnd]] == 0 {
					flag = false
					break
				}
				m[s[tempStart:tempEnd]]--
				tempStart = tempEnd
			}
		}
		// fmt.Printf("start=%d,end=%d,tempEnd=%d\n", start, end, tempEnd)

		for _, word := range words {
			m[word] = 0
		}
		for _, word := range words {
			m[word]++
		}
		if flag && tempEnd-start == end-start {
			result = append(result, start)
		}
		start++
		end++
	}

	return result
}

// official
func findSubstring(s string, words []string) []int {
	var ls, m, n = len(s), len(words), len(words[0])

	var result = make([]int, 0)

	for i := 0; i < n && i+m*n <= ls; i++ {
		var differ = make(map[string]int, 0)
		for j := 0; j < m; j++ {
			differ[s[i+j*n:i+(j+1)*n]]++
		}
		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}
		for start := i; start <= ls-m*n; start += n {
			if start != i {
				word := s[start+(m-1)*n : start+m*n]
				differ[word]++
				if differ[word] == 0 {
					delete(differ, word)
				}
				word = s[start-n : start]
				differ[word]--
				if differ[word] == 0 {
					delete(differ, word)
				}
			}
			if len(differ) == 0 {
				result = append(result, start)
			}
		}
	}
	return result
}
