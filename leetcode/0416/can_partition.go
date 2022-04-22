package leetcode

func CanPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 {
		return false
	}
	return recursion(nums, 0, sum/2)
}

func recursion(nums []int, i, sum int) bool {
	if sum == 0 {
		return true
	}
	if sum < 0 || i == len(nums) {
		return false
	}
	return recursion(nums, i+1, sum-nums[i]) || recursion(nums, i+1, sum)
}

func CanPartitionUseMemo(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2
	memo := make([][]bool, len(nums)+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]bool, sum+1)
	}

	return recursionUseMemo(nums, 0, sum, memo)
}

func recursionUseMemo(nums []int, i, sum int, memo [][]bool) bool {
	if sum == 0 {
		return true
	}
	if sum < 0 || i == len(nums) {
		return false
	}
	if memo[i][sum] {
		return memo[i][sum]
	}
	res := recursionUseMemo(nums, i+1, sum-nums[i], memo) || recursionUseMemo(nums, i+1, sum, memo)
	memo[i][sum] = res
	return res
}

func CanPartitionDP(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)
	for i := 0; i <= target; i++ {
		dp[i] = (nums[0] == i)
	}
	for i := 1; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[target]
}
