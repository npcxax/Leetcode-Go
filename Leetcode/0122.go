package leetcode

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit(prices []int) int {
	var n = len(prices)
	var dp0, dp1 = -prices[0], 0

	for i := 1; i < n; i++ {
		dp0, dp1 = max(dp0, dp1-prices[i]), max(dp1, dp0+prices[i])
	}

	return dp1
}
