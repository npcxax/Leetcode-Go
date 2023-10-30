package leetcode

func lengthOfLastWord(s string) int {
	n := len(s)
	i := n - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	j := i
	for j >= 0 && s[j] != ' ' {
		j--
	}
	return i - j
}
