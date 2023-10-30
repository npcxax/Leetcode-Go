package leetcode

func candy(ratings []int) int {
	var n = len(ratings)
	var left = make([]int, n)

	left[0] = 1
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	var total = left[n-1]
	var right = 1
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		total += max(left[i], right)
	}

	return total
}
