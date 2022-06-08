package leetcode

func MaxProfit(prices []int, fee int) int {
	cash := 0
	hold := -prices[0]
	for i := 1; i < len(prices); i++ {
		cash = max(cash, hold+prices[i]-fee)
		hold = max(hold, cash-prices[i])
	}
	return cash
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
