package leetcode

func isIsomorphic(s string, t string) bool {
	var m, n = len(s), len(t)
	if m != n {
		return false
	}
	var countS = make([]byte, 26)
	var countT = make([]byte, 26)

	for i := 0; i < 26; i++ {
		countS[i] = 0
		countT[i] = 0
	}

	for i, ch := range s {
		if countS[ch-'a'] != 0 && countT[t[i]-'a'] != 0 {
			if countS[ch-'a'] != t[i] {
				return false
			}
			if countT[t[i]-'a'] != s[i] {
				return false
			}
		} else if countS[ch-'a'] == 0 && countT[t[i]-'a'] == 0 {
			countS[ch-'a'] = t[i]
			countT[t[i]-'a'] = s[i]
		} else {
			return false
		}
	}
	return true
}
