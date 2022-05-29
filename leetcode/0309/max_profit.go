package leetcode

import "math"

func MaxProfit(prices []int) int {
	sold, held, reset := math.MinInt64, math.MinInt64, 0
	for _, price := range prices {
		preSold := sold
		sold = held + price
		held = max(held, reset-price)
		reset = max(reset, preSold)
	}
	return max(sold, reset)
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
