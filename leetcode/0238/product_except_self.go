package leetcode

func ProductExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range nums {
		ans[i] = 1
	}
	prefix := 1
	for i := 0; i < len(nums); i++ {
		ans[i] *= prefix
		prefix *= nums[i]
	}
	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] *= suffix
		suffix *= nums[i]
	}
	return ans
}
