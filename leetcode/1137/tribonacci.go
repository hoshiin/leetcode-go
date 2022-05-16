package leetcode

import "math"

func Tribonacci(n int) int {
	return tribonacciInternal(n, map[int]int{0: 0, 1: 1})
}

func tribonacciInternal(n int, memo map[int]int) int {
	if n < 0 {
		return 0
	}
	if num, ok := memo[n]; ok {
		return num
	}
	memo[n] = tribonacciInternal(n-1, memo) + tribonacciInternal(n-2, memo) + tribonacciInternal(n-3, memo)
	return memo[n]
}

func TribonacciDP(n int) int {
	nums := make([]int, int(math.Max(float64(n+1), 3.0)))
	nums[1] = 1
	nums[2] = 1

	for i := 3; i < n+1; i++ {
		nums[i] = nums[i-1] + nums[i-2] + nums[i-3]
	}

	return nums[n]
}
