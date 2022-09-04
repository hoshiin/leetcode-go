package leetcode

import "sort"

func SmallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	left, right := 0, nums[len(nums)-1]-nums[0]
	for left < right {
		mid := (left + right) / 2
		if isSmallPairs(nums, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func isSmallPairs(nums []int, k, mid int) bool {
	count, left := 0, 0
	for right := 1; right < len(nums); right++ {
		for nums[right]-nums[left] > mid {
			left++
		}
		count += right - left
	}
	return count >= k
}
