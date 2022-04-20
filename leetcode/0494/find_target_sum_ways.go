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

func FindTargetSumWaysWithDP(nums []int, S int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if S > sum || (sum+S)%2 == 1 || S+sum < 0 {
		return 0
	}
	target := (S + sum) / 2
	dp := make([]int, target+1)
	dp[0] = 1
	for _, n := range nums {
		for i := target; i >= n; i-- {
			dp[i] += dp[i-n]
		}
	}
	return dp[target]
}

func FindTargetSumWaysDFS(nums []int, S int) int {
	sums := make([]int, len(nums))
	sums[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i > -1; i-- {
		sums[i] = sums[i+1] + nums[i]
	}
	res := 0
	dfsFindTargetSumWays(nums, 0, 0, S, &res, sums)
	return res
}

func dfsFindTargetSumWays(nums []int, i, curSum, S int, res *int, sums []int) {
	if i == len(nums) {
		if curSum == S {
			*res = *res + 1
		}
		return
	}
	if S-curSum > sums[i] {
		return
	}
	dfsFindTargetSumWays(nums, i+1, curSum+nums[i], S, res, sums)
	dfsFindTargetSumWays(nums, i+1, curSum-nums[i], S, res, sums)
}
