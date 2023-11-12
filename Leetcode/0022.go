package leetcode

// violence
func generateParenthesis(n int) []string {
	var (
		result  = make([]string, 0)
		recur   func(idx int, b []byte)
		isValid func(s string) bool
	)
	isValid = func(s string) bool {
		balance := 0
		for _, ch := range s {
			if balance < 0 {
				return false
			}
			if ch == '(' {
				balance++
			} else {
				balance--
			}
		}
		return balance == 0
	}

	recur = func(idx int, b []byte) {
		if idx == 2*n {
			if isValid(string(b)) {
				result = append(result, string(b))
			}
			return
		}
		b[idx] = '('
		recur(idx+1, b)
		b[idx] = ')'
		recur(idx+1, b)
	}

	recur(0, make([]byte, 2*n))

	return result
}

// backtrace
func generateParenthesis(n int) []string {
	var (
		result = make([]string, 0)
		recur  func(nLeft, nRight int, b []byte)
	)

	recur = func(nLeft, nRight int, b []byte) {
		if len(b) == 2*n {
			result = append(result, string(b))
			return
		}
		if nLeft < n {
			b = append(b, '(')
			recur(nLeft+1, nRight, b)
			b = b[:len(b)-1]
		}
		if nRight < nLeft {
			b = append(b, ')')
			recur(nLeft, nRight+1, b)
			b = b[:len(b)-1]
		}
	}

	recur(0, 0, make([]byte, 0, 2*n))

	return result
}
