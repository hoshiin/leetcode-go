package leetcode

func FindMaxConsecutiveOnes(nums []int) int {
	longestSequence := 0
	left, right := 0, 0
	numZeroes := 0

	for right < len(nums) {
		if nums[right] == 0 {
			numZeroes++
		}
		for numZeroes == 2 {
			if nums[left] == 0 {
				numZeroes--
			}
			left++
		}
		longestSequence = max(longestSequence, right-left+1)
		right++
	}
	return longestSequence
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
