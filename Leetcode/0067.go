package leetcode

import (
	"fmt"
	"strconv"
)

// simulate, my solution
func addBinary(a string, b string) string {
	m, n := len(a), len(b)
	res := ""
	carry := 0
	for i, j := m-1, n-1; carry == 1 || i >= 0 || j >= 0; i,

		j = i-1, j-1 {
		total := carry
		if i >= 0 {
			total += int(a[i] - '0')
		}
		if j >= 0 {
			total += int(b[j] - '0')
		}
		fmt.Println(total)
		carry = total / 2
		res = strconv.Itoa(total%2) + res
	}
	return res
}
