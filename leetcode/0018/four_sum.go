package leetcode

import "sort"

func FourSum(nums []int, target int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if nums[i]+nums[len(nums)-1]+nums[len(nums)-2]+nums[len(nums)-3] < target {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			if nums[i]+nums[j]+nums[len(nums)-2]+nums[len(nums)-1] < target {
				continue
			}
			head, end := j+1, len(nums)-1
			for head < end {
				cur := nums[i] + nums[j] + nums[head] + nums[end]
				if cur == target {
					ans = append(ans, []int{nums[i], nums[j], nums[head], nums[end]})
					head++
					for head < end && nums[head] == nums[head-1] {
						head++
					}
				} else if cur > target {
					end--
				} else {
					head++
				}
			}
		}
	}
	return ans
}
