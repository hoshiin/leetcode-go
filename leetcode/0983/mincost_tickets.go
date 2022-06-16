package leetcode

func MincostTickets(days []int, costs []int) int {
	dp := make([]int, days[len(days)-1]+1)
	for i := 0; i < len(days); i++ {
		dp[days[i]] = -1
	}
	for i := 1; i <= days[len(days)-1]; i++ {
		if dp[i] == 0 {
			dp[i] = dp[i-1]
		} else {
			dp[i] = costs[0] + dp[i-1]
		}
		sevenDay := costs[1]
		if i >= 7 {
			sevenDay = dp[i-7] + costs[1]
		}
		thirtyDay := costs[2]
		if i >= 30 {
			thirtyDay = dp[i-30] + costs[2]
		}
		dp[i] = min(dp[i], sevenDay, thirtyDay)
	}
	return dp[days[len(days)-1]]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}
