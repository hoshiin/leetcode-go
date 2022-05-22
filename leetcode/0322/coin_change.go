package leetcode

import "math"

func CoinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}
	return dp(coins, amount, make([]int, amount))
}

func dp(coins []int, rem int, count []int) int {
	if rem < 0 {
		return -1
	}
	if rem == 0 {
		return 0
	}
	if count[rem-1] != 0 {
		return count[rem-1]
	}
	min := math.MaxInt64
	for _, coin := range coins {
		res := dp(coins, rem-coin, count)
		if res >= 0 && res < min {
			min = 1 + res
		}
	}
	if min == math.MaxInt64 {
		count[rem-1] = -1
	} else {
		count[rem-1] = min
	}
	return count[rem-1]
}
