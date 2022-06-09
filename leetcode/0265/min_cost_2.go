package leetcode

import "math"

func MinCostII(costs [][]int) int {
	n, k := len(costs), len(costs[0])
	for i := 1; i < n; i++ {
		smallest, secondSmallest := math.MaxInt32, math.MaxInt32
		for j := 0; j < k; j++ {
			if costs[i-1][j] <= smallest {
				secondSmallest = smallest
				smallest = costs[i-1][j]
			} else if costs[i-1][j] < secondSmallest {
				secondSmallest = costs[i-1][j]
			}
		}
		for j := 0; j < k; j++ {
			minToAdd := smallest
			if minToAdd == costs[i-1][j] {
				minToAdd = secondSmallest
			}
			costs[i][j] += minToAdd
		}
	}
	res := costs[n-1][0]
	for j := 1; j < k; j++ {
		res = min(res, costs[n-1][j])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
