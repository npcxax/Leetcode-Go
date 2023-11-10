package leetcode

func letterCombinations(digits string) []string {
	var (
		num2ch = [][]rune{{}, {}, {'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'},
			{'j', 'k', 'l'}, {'m', 'n', 'o'}, {'p', 'q', 'r', 's'}, {'t', 'u', 'v'},
			{'w', 'x', 'y', 'z'},
		}
		result = make([]string, 0)
		dfs    func(cur int, b []rune)
		m      = len(digits)
	)
	if m == 0 {
		return result
	}
	dfs = func(cur int, b []rune) {
		if cur == m {
			result = append(result, string(b))
			return
		}
		for _, ch := range num2ch[digits[cur]-'0'] {
			dfs(cur+1, append(b, ch))
		}
	}

	dfs(0, []rune{})

	return result
}
