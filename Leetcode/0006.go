package leetcode

import "bytes"

func convert(s string, numRows int) string {
	var n = len(s)
	if numRows == 1 || numRows > n {
		return s
	}
	// r = numRows, row = r, t = r+r-2 = 2r-2
	// n/t round up periods, and every period have r-1 column
	var matrix = make([][]byte, numRows)
	// (n+2r-1)/(2r-2) /2
	var x = 0
	var count = 0
	for count < n {
		matrix[x] = append(matrix[x], s[count])
		if count%(2*numRows-2) < numRows-1 {
			x++
		} else {
			x--
		}
		count++
	}
	return string(bytes.Join(matrix, nil))
}
