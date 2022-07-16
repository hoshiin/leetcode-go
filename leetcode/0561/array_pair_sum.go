package leetcode

import "sort"

func ArrayPairSum(nums []int) int {
	sum := 0
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}
