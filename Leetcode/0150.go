package leetcode

import (
	"strconv"
)

// my solution
func evalRPN(tokens []string) int {
	var hashmap = map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	var stk = make([]int, 0)
	for _, item := range tokens {
		if hashmap[item] {
			n := len(stk)
			n1, n2 := stk[n-1], stk[n-2]
			switch item {
			case "+":
				n1 = n2 + n1
			case "-":
				n1 = n2 - n1
			case "*":
				n1 = n2 * n1
			case "/":
				n1 = n2 / n1
			}
			stk = stk[:n-2]
			stk = append(stk, n1)
		} else {
			num, _ := strconv.Atoi(item)
			stk = append(stk, num)
		}
	}

	return stk[0]
}

func evalRPN(tokens []string) int {
	var stk = make([]int, 0)
	for _, item := range tokens {
		num, err := strconv.Atoi(item)
		if err == nil {
			stk = append(stk, num)
		} else {
			n := len(stk)
			n1, n2 := stk[n-1], stk[n-2]
			switch item {
			case "+":
				n1 = n2 + n1
			case "-":
				n1 = n2 - n1
			case "*":
				n1 = n2 * n1
			case "/":
				n1 = n2 / n1
			}
			stk = stk[:n-2]
			stk = append(stk, n1)
		}
	}

	return stk[0]
}
