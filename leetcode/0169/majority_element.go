package leetcode

import "sort"

func MajorityElement(nums []int) int {
	majority := len(nums) / 2
	index := make(map[int]int)
	for _, n := range nums {
		if count := index[n]; count >= majority {
			return n
		}
		index[n]++
	}
	return -1
}

func MajorityElementSort(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums[len(nums)/2]
}

func MajorityElementBoyceMooreAlgorithms(nums []int) int {
	count := 0
	candidate := -1
	for _, n := range nums {
		if count == 0 {
			candidate = n
		}
		if n == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
