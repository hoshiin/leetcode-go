package leetcode

import "math"

func Rob(nums []int) int {
	return dp(nums, len(nums)-1, make(map[int]int))
}

func dp(nums []int, i int, memo map[int]int) int {
	if i == 0 {
		return nums[0]
	}
	if i == 1 {
		return int(math.Max(float64(nums[0]), float64(nums[1])))
	}
	if _, ok := memo[i]; !ok {
		memo[i] = int(math.Max(float64(dp(nums, i-1, memo)), float64(dp(nums, i-2, memo)+nums[i])))
	}
	return memo[i]
}

func RobBottomUp(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = int(math.Max(float64(nums[0]), float64(nums[1])))

	for i := 2; i < len(nums); i++ {
		dp[i] = int(math.Max(float64(dp[i-1]), float64(dp[i-2]+nums[i])))
	}
	return dp[len(nums)-1]
}
