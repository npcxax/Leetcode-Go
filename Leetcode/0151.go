package leetcoode

func reverseWords(s string) string {
	// 1. erase space in s
	slowIndex, fastIndex := 0, 0

	for len(s) > 0 && fastIndex < len(s) && s[fastIndex] == ' ' {
		fastIndex++
	}
	b := []byte(s)
	for ; fastIndex < len(s); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	}
	// erase tail's space
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' {
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}
	// 2. reverse s
	slowIndex = len(b)
	reverse(&b, 0, slowIndex-1)

	// 3. reverse words
	i := 0
	for i < slowIndex {
		j := i
		for j < slowIndex && b[j] != ' ' {
			j++
		}
		reverse(&b, i, j-1)
		i = j
		i++
	}
	return string(b)
}

func reverse(b *[]byte, left int, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}
