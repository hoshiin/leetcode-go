package leetcode

import "math"

func MinCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = cost[i] + int(math.Min(float64(dp[i-1]), float64(dp[i-2])))
	}
	return int(math.Min(float64(dp[len(cost)-1]), float64(dp[len(cost)-2])))
}

func MinCostClimbingStairsRecursive(cost []int) int {
	return min(minCost(cost, len(cost)-1, map[int]int{}), minCost(cost, len(cost)-2, map[int]int{}))
}

func minCost(cost []int, n int, dp map[int]int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return cost[n]
	}
	if c, ok := dp[n]; ok {
		return c
	}
	dp[n] = cost[n] + min(minCost(cost, n-1, dp), minCost(cost, n-2, dp))
	return dp[n]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
