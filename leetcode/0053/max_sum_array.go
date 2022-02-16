package leetcode

import "math"

// MaxSubArray uses Dynamic Programming, Kadane's Algorithm
func MaxSubArray(nums []int) int {
	currentSubArray := float64(nums[0])
	maxSubArray := float64(nums[0])

	for i := 1; i < len(nums); i++ {
		num := float64(nums[i])
		currentSubArray = math.Max(num, currentSubArray+num)
		maxSubArray = math.Max(maxSubArray, currentSubArray)
	}

	return int(maxSubArray)
}

// MaxSubArray uses recursion
func MaxSubArrayRecursion(nums []int) int {
	numsArray := nums
	return findBestSubArray(numsArray, 0, len(numsArray)-1)
}

func findBestSubArray(nums []int, left, right int) int {
	// Base case - empty array.
	if left > right {
		return int(math.Inf(-1))
	}

	mid := (left + right) / 2
	curr, bestLeftSum, bestRightSum := float64(0), float64(0), float64(0)

	// Iterate from the middle to the beginning.
	for i := mid - 1; i >= left; i-- {
		curr += float64(nums[i])
		bestLeftSum = math.Max(bestLeftSum, curr)
	}

	// Reset curr and iterate from the middle to the end.
	curr = 0
	for i := mid + 1; i <= right; i++ {
		curr += float64(nums[i])
		bestRightSum = math.Max(bestRightSum, curr)
	}

	// The bestCombinedSum uses the middle element and the best
	// possible sum from each half.
	bestCombinedSum := float64(nums[mid]) + bestLeftSum + bestRightSum

	// Find the best subarray possible from both halves.
	leftHalf := float64(findBestSubArray(nums, left, mid-1))
	rightHalf := float64(findBestSubArray(nums, mid+1, right))

	// The largest of the 3 is the answer for any given input array.
	return int(math.Max(math.Max(bestCombinedSum, leftHalf), rightHalf))
}
