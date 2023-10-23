package leetcode

func minWindow(s string, t string) string {
	var ori, cnt = make(map[byte]int, 0), make(map[byte]int, 0)
	for _, ch := range t {
		ori[byte(ch)]++
	}
	var ansL, ansR = -1, -1
	var sLen = len(s)
	var maxLen = 0x7fffffff

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < maxLen {
				maxLen = r - l + 1
				ansL = l
				ansR = l + maxLen
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]]--
			}
			l++
		}
	}

	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
