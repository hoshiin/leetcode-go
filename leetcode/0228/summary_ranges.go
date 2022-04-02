package leetcode

import "fmt"

func SummaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	last := nums[0]
	from := nums[0]
	to := 0
	ranges := []string{}
	for i := 1; i < len(nums); i++ {
		if nums[i] != last+1 {
			to = last
			ranges = append(ranges, printRanges(from, to))
			from = nums[i]
		}
		last = nums[i]
	}
	to = last
	ranges = append(ranges, printRanges(from, to))
	return ranges
}

func printRanges(from, to int) string {
	if from == to {
		return fmt.Sprintf("%d", from)
	}
	return fmt.Sprintf("%d->%d", from, to)
}

func SummaryRangesOptimized(nums []int) []string {
	ranges := []string{}
	for i, j := 0, 0; j < len(nums); j++ {
		if j+1 < len(nums) && nums[j]+1 == nums[j+1] {
			continue
		}
		if i == j {
			ranges = append(ranges, fmt.Sprintf("%d", nums[i]))
		} else {
			ranges = append(ranges, fmt.Sprintf("%d->%d", nums[i], nums[j]))
		}
		i = j + 1
	}
	return ranges
}
