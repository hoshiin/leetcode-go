package leetcode

import "math"

func MaxProfit(prices []int) int {
	minprice := math.MaxInt64
	maxprofit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		}
		if prices[i]-minprice > maxprofit {
			maxprofit = prices[i] - minprice
		}
	}
	return maxprofit
}

func MaxProfitBruteForce(prices []int) int {
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > max {
				max = profit
			}
		}
	}
	return max
}
