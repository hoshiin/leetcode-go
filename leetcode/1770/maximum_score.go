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
