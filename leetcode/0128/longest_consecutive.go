package leetcode

import "sort"

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	maxCount, cur := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			continue
		}
		if nums[i-1]+1 == nums[i] {
			cur++
		} else {
			maxCount = max(maxCount, cur)
			cur = 1
		}
	}
	return max(maxCount, cur)
}

func LongestConsecutiveN(nums []int) int {
	index := make(map[int]struct{})
	for _, n := range nums {
		index[n] = struct{}{}
	}
	longestStreak := 0
	for n := range index {
		if _, ok := index[n-1]; !ok {
			curr := n
			currStreak := 1
			_, ok := index[curr+1]
			for ok {
				curr++
				currStreak++
				_, ok = index[curr+1]
			}
			longestStreak = max(longestStreak, currStreak)
		}
	}
	return longestStreak
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
