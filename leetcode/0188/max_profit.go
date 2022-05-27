package leetcode

var memo [][][]int

func MaxProfit(k int, prices []int) int {
	memo = make([][][]int, len(prices))
	for i := range memo {
		memo[i] = make([][]int, k+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, 2)
		}
	}
	return dp(0, k, 0, prices)
}

func dp(i, transactionsRemaining, holding int, prices []int) int {
	if transactionsRemaining == 0 || len(prices) == i {
		return 0
	}
	if memo[i][transactionsRemaining][holding] == 0 {
		doNothing := dp(i+1, transactionsRemaining, holding, prices)
		doSomething := 0

		if holding == 1 {
			// Sell Stock
			doSomething = prices[i] + dp(i+1, transactionsRemaining-1, 0, prices)
		} else {
			// Buy Stock
			doSomething = -prices[i] + dp(i+1, transactionsRemaining, 1, prices)
		}
		// Recurrence relation. Choose the most profitable option.
		memo[i][transactionsRemaining][holding] = max(doNothing, doSomething)
	}
	return memo[i][transactionsRemaining][holding]
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
