package leetcode

// my solution
func lengthOfLongestSubstring(s string) int {
	var m = make(map[byte]int, 0)
	var n = len(s)
	var start, end = 0, 0
	var result = 0
	var curLen = 0

	for end < n {
		if _, ok := m[s[end]]; ok {
			for start <= m[s[end]] {
				delete(m, s[start])
				start++
			}
		}
		m[s[end]] = end
		curLen = end - start + 1
		end++
		if curLen > result {
			result = curLen
		}
	}

	return result
}

// official
func lengthOfLongestSubstring(s string) int {
	var m = make(map[byte]int, 0)
	var n = len(s)
	var rk, ans = -1, 0

	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		if rk-i+1 > ans {
			ans = rk - i + 1
		}
	}

	return ans
}
