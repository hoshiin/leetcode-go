package leetcode

func Rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	max1 := dp(nums[1:])
	max2 := dp(nums[:len(nums)-1])
	return max(max1, max2)
}

func dp(nums []int) int {
	prevMax, currMax := 0, 0
	for _, num := range nums {
		tmp := currMax
		currMax = max(prevMax+num, currMax)
		prevMax = tmp
	}
	return currMax
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
