package leetcode

import "sort"

func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for i, num := range nums {
		if num > 0 {
			break
		}
		if i > 0 && num == nums[i-1] {
			continue
		}
		low, high := i+1, len(nums)-1
		for low < high {
			if nums[i]+nums[low]+nums[high] < 0 {
				low++
				continue
			}
			if nums[i]+nums[low]+nums[high] > 0 {
				high--
				continue
			}
			res = append(res, []int{nums[i], nums[low], nums[high]})
			low++
			high--
			for low < high && nums[low] == nums[low-1] {
				low++
			}
		}
	}
	return res
}

func ThreeSumHash(nums []int) [][]int {
	ans := [][]int{}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for i, num := range nums {
		seen := map[int]struct{}{}
		if num > 0 {
			break
		}
		if i > 0 && nums[i-1] == num {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			complement := -nums[i] - nums[j]
			if _, ok := seen[complement]; ok {
				ans = append(ans, []int{nums[i], nums[j], complement})
				for j < len(nums)-1 && nums[j] == nums[j+1] {
					j++
				}
			}
			seen[nums[j]] = struct{}{}
		}
	}
	return ans
}
