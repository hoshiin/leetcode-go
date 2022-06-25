package leetcode

import (
	"math"
	"sort"
)

func SortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	for i, num := range nums {
		ans[i] = num * num
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	return ans
}

func SortedSquaresUseAbs(nums []int) []int {
	ans := make([]int, len(nums))
	left, right := 0, len(nums)-1

	for i := len(nums) - 1; i >= 0; i-- {
		if math.Abs(float64(nums[left])) < math.Abs(float64(nums[right])) {
			ans[i] = nums[right] * nums[right]
			right--
		} else {
			ans[i] = nums[left] * nums[left]
			left++
		}
	}
	return ans
}
