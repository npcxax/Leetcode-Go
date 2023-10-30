package leetcode

var rToI = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// func romanToInt(s string) int {
// 	// I,X,C可以放在左边
// 	var curR byte = 'I'
// 	var result = 0
// 	for i := len(s) - 1; i >= 0; i-- {
// 		if rToI[s[i]] > rToI[curR] {
// 			curR = s[i]
// 		}
// 		if rToI[s[i]] < rToI[curR] {
// 			result -= rToI[s[i]]
// 		} else {
// 			result += rToI[s[i]]
// 		}
// 	}

// 	return result
// }

func romanToInt(s string) int {
	n := len(s)
	result := 0
	for i := range s {
		value := rToI[s[i]]
		if i < n-1 && value < rToI[s[i+1]] {
			result -= value
		} else {
			result += value
		}
	}
	return result
}
