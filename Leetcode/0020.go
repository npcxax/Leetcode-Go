package leetcode

func isValid(s string) bool {
	var hashmap = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stk = make([]rune, 0)

	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stk = append(stk, ch)
		} else {
			n := len(stk)
			if n == 0 {
				return false
			} else if hashmap[ch] != stk[n-1] {
				return false
			}
			stk = stk[:n-1]
		}
	}

	return len(stk) == 0
}
