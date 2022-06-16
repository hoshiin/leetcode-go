package leetcode

var mod = 1000000000 + 7

func NumRollsToTarget(n int, k int, target int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return helper(n, k, target, dp)
}

func helper(n, k, target int, dp [][]int) int {
	if n == 0 {
		if target == 0 {
			return 1
		} else {
			return 0
		}
	}
	if dp[n][target] >= 0 {
		return dp[n][target]
	}
	numWays := 0
	for i := 1; i <= k; i++ {
		rem := target - i
		if rem < 0 {
			break
		}
		numWays = (numWays + helper(n-1, k, rem, dp)) % mod
	}
	dp[n][target] = numWays
	return numWays
}
