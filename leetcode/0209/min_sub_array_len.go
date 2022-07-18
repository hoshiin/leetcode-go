package leetcode

import "math"

func MinSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans := math.MaxInt
	left, sum := 0, 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		for sum >= target {
			ans = min(ans, i+1-left)
			sum -= nums[left]
			left++
		}
	}
	if ans == math.MaxInt {
		return 0
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
