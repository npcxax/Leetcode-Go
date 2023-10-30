package leetcode

func maxProfit(prices []int) int {
	var p = prices[0]
	var maxprofit = 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < p {
			p = prices[i]
		} else if prices[i]-p > maxprofit {
			maxprofit = prices[i] - p
		}
	}

	return maxprofit
}
