package leetcode

func LengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	longest := 0
	for _, num := range dp {
		longest = max(longest, num)
	}
	return longest
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
