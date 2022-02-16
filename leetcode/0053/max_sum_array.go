package leetcode

import "math"

// MaxSubArray uses Dynamic Programming, Kadane's Algorithm
func MaxSubArray(nums []int) int {
	currentSubArray := float64(nums[0])
	maxSubArray := float64(nums[0])

	for i := 1; i < len(nums); i++ {
		num := float64(nums[i])
		currentSubArray = math.Max(float64(num), currentSubArray+num)
		maxSubArray = math.Max(maxSubArray, currentSubArray)
	}

	return int(maxSubArray)
}

// TODO: Recursion Solution
