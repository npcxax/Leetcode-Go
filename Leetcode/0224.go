package leetcode

import (
	"strconv"
	"strings"
)

var prior = map[byte]int{
	'+': 1,
	'-': 1,
}

func cal(numStk *[]int, opStk *[]byte) {
	n, m := len(*numStk), len(*opStk)
	if n == 0 || n < 2 {
		return
	}
	if m == 0 {
		return
	}

	n1, n2 := (*numStk)[n-1], (*numStk)[n-2]
	switch (*opStk)[m-1] {
	case '+':
		n2 = n2 + n1
	case '-':
		n2 = n2 - n1
	}
	*numStk = (*numStk)[:n-2]
	*numStk = append(*numStk, n2)
	*opStk = (*opStk)[:m-1]
}

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")

	var numStk = make([]int, 0)
	numStk = append(numStk, 0)

	var opStk = make([]byte, 0)
	var n = len(s)

	var isNum = func(b byte) bool {
		return b >= '0' && b <= '9'
	}

	for i := 0; i < n; i++ {
		if s[i] == '(' {
			opStk = append(opStk, s[i])
		} else if s[i] == ')' {
			for len(opStk) > 0 && opStk[len(opStk)-1] != '(' {
				cal(&numStk, &opStk)
			}
			opStk = opStk[:len(opStk)-1]
		} else {
			if isNum(s[i]) { // num
				j := i
				for j < n && isNum(s[j]) {
					j++
				}
				num, _ := strconv.Atoi(s[i:j])
				i = j - 1
				numStk = append(numStk, num)
			} else { // operator
				if i > 0 && (s[i-1] == '(' || s[i-1] == '+' || s[i-1] == '-') {
					numStk = append(numStk, 0)
				}
				for len(opStk) > 0 && opStk[len(opStk)-1] != '(' {
					cal(&numStk, &opStk)
				}
				opStk = append(opStk, s[i])
			}
		}
	}
	for len(opStk) != 0 {
		cal(&numStk, &opStk)
	}
	return numStk[len(numStk)-1]
}
