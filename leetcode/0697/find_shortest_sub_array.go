package leetcode

import "math"

func FindShortestSubArray(nums []int) int {
	list := make(map[int]int)
	degree := 0
	firstSeen := make(map[int]int)
	minLength := len(nums)
	for i, n := range nums {
		if _, ok := list[n]; ok {
			list[n]++
		} else {
			list[n] = 1
			firstSeen[n] = i
		}
		if list[n] > degree {
			degree = list[n]
			minLength = i - firstSeen[n] + 1
		} else if list[n] == degree {
			minLength = int(math.Min(float64(minLength), float64(i-firstSeen[n]+1)))
		}
	}
	return minLength
}
