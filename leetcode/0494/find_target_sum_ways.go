package leetcode

func FindTargetSumWays(nums []int, target int) int {
	c := 0
	calculate(nums, &c, 0, 0, target)
	return c
}

func calculate(nums []int, count *int, i, sum, target int) {
	if i == len(nums) {
		if sum == target {
			*count++
		}
	} else {
		calculate(nums, count, i+1, sum+nums[i], target)
		calculate(nums, count, i+1, sum-nums[i], target)
	}
}

// TODO: Dynamic Programming and DFS Solutions
