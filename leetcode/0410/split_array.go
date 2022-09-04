package leetcode

import "math"

func SplitArray(nums []int, m int) int {
	sum := 0
	maxElem := math.MinInt
	for _, num := range nums {
		sum += num
		maxElem = max(maxElem, num)
	}

	// Define the left and right boundary of binary search
	left, right := maxElem, sum
	ans := right
	for left <= right {
		// Find the mid value
		mid := left + (right-left)/2
		// Find the minimum splits. If splitsRequired is less than
		// or equal to m move towards left i.e., smaller values
		if canSplit(nums, mid, m) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func canSplit(nums []int, largest, m int) bool {
	subArray, curSum := 0, 0
	for _, n := range nums {
		curSum += n
		if curSum > largest {
			subArray++
			curSum = n
		}
	}
	return subArray+1 <= m
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
