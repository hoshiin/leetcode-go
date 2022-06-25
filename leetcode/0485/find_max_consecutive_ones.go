package leetcode

func FindMaxConsecutiveOnes(nums []int) int {
	maxConsecutive := 0
	cur := 0
	for _, num := range nums {
		if num == 1 {
			cur++
		} else {
			cur = 0
		}
		maxConsecutive = max(maxConsecutive, cur)
	}
	return maxConsecutive
}

func FindMaxConsecutiveOnesUseCalc(nums []int) int {
	max, cur := 0, 0
	for _, num := range nums {
		if cur+num > max {
			max = cur + num
		}
		cur = (cur + num) * num
	}
	return max
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
