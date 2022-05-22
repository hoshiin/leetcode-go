package leetcode

func MaximumScore(nums []int, multipliers []int) int {
	memo := make([][]int, len(multipliers))
	for i, _ := range memo {
		memo[i] = make([]int, len(multipliers))
	}
	return dp(nums, multipliers, memo, 0, 0)
}

func dp(nums []int, multipliers []int, memo [][]int, i, left int) int {
	if i == len(multipliers) {
		return 0
	}
	mult := multipliers[i]
	right := len(nums) - 1 - (i - left)
	if memo[i][left] == 0 {
		memo[i][left] = max(mult*nums[left]+dp(nums, multipliers, memo, i+1, left+1), mult*nums[right]+dp(nums, multipliers, memo, i+1, left))
	}
	return memo[i][left]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaximumScoreBottomUp(nums []int, multipliers []int) int {
	n := len(nums)
	m := len(multipliers)
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := m - 1; i >= 0; i-- {
		for left := i; left >= 0; left-- {
			mult := multipliers[i]
			right := n - 1 - (i - left)
			dp[i][left] = max(mult*nums[left]+dp[i+1][left+1], mult*nums[right]+dp[i+1][left])
		}
	}
	return dp[0][0]
}
