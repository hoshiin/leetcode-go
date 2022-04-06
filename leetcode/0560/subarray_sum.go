package leetcode

func SubarraySum(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}

func subarraySumOptimized(nums []int, k int) int {
	count, sum := 0, 0
	mp := map[int]int{0: 1}
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if c, ok := mp[sum-k]; ok {
			count += c
		}
		mp[sum]++
	}
	return count
}
