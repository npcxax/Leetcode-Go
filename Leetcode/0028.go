package leetcode

// violent
func strStr(haystack string, needle string) int {
	var m, n = len(haystack), len(needle)
	var flag = true
	for i := 0; i <= m-n; i++ {
		flag = true
		for j := 0; j < n; j++ {
			if haystack[i+j] != needle[j] {
				flag = false
				break
			}
		}
		if flag {
			return i
		}
	}
	return -1
}

// KMP
func buildMatch(needle string, match []int) {
	var m = len(needle)
	match[0] = -1
	for i := 1; i < m; i++ {
		j := match[i-1]
		for j >= 0 && needle[j+1] != needle[i] {
			j = match[j]
		}
		if needle[j+1] == needle[i] {
			match[i] = j + 1
		} else {
			match[i] = -1
		}
	}
}

func strStr(haystack string, needle string) int {
	var m, n = len(haystack), len(needle)
	var match = make([]int, n)

	buildMatch(needle, match)

	var s, p = 0, 0
	for s < m && p < n {
		if haystack[s] == needle[p] {
			s++
			p++
		} else if p > 0 {
			p = match[p-1] + 1
		} else {
			s++
		}
	}
	if p == n {
		return s - n
	}
	return -1
}
