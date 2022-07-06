package leetcode

import "math"

func ThirdMax(nums []int) int {
	maxArr := [3]int{math.MinInt, math.MinInt, math.MinInt}
	for _, num := range nums {
		if maxArr[0] == num || maxArr[1] == num || maxArr[2] == num {
			continue
		}
		if maxArr[0] < num {
			maxArr[0], maxArr[1], maxArr[2] = num, maxArr[0], maxArr[1]
		} else if maxArr[1] < num {
			maxArr[2], maxArr[1] = maxArr[1], num
		} else if maxArr[2] < num {
			maxArr[2] = num
		}
	}
	if maxArr[2] == math.MinInt {
		return maxArr[0]
	}
	return maxArr[2]
}
