package leetcode

import "strings"

// my solution
func isPalindrome(s string) bool {
	s = removeNotCharacter(s)
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func removeNotCharacter(s string) string {
	var b = make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') {
			b = append(b, s[i])
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			b = append(b, s[i]-'A'+'a')
		}
	}
	return string(b)
}

// official
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var left, right = 0, len(s) - 1
	for left < right {
		for left < right && !isalnum(s[left]) {
			left++
		}
		for left < right && !isalnum(s[right]) {
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
