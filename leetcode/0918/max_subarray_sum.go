package leetcode

import "math"

func MaxSubarraySumCircular(nums []int) int {
	result := kadane(nums)
	sum := 0
	for i, v := range nums {
		sum += v
		nums[i] = -v
	}
	if result < 0 {
		return result
	}
	return max(result, sum+kadane(nums))
}

func kadane(nums []int) int {
	sum := 0
	result := math.MinInt32

	for _, v := range nums {
		sum += v
		result = max(result, sum)
		sum = max(sum, 0)
	}
	return result
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
