package leetcode

func isSubsequence(s string, t string) bool {
	// preprocessing, solve t
	var m = len(t)
	var dp = make([][]int, 0)

	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, 26))
	}

	for i := 0; i < 26; i++ {
		dp[m][i] = m
	}

	for i := m - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if t[i]-'a' == byte(j) {
				dp[i][j] = i
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}

	// solve s
	var n = len(s)

	var i, j = 0, 0
	for i < n {
		j = dp[j][s[i]-'a']
		if j == m {
			return false
		}
		j += 1
		i++
	}
	return true
}
