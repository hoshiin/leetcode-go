package leetcode

func MaxProfit(prices []int) int {
	i := 0
	maxprofit := 0
	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		valley := prices[i]
		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak := prices[i]
		maxprofit += peak - valley
	}
	return maxprofit
}
