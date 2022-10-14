package leetcode

func MissingNumber(nums []int) int {
	index := make([]bool, len(nums)+1)
	for _, n := range nums {
		index[n] = true
	}
	for i, t := range index {
		if !t {
			return i
		}
	}
	return -1
}

func MissingNumberBit(nums []int) int {
	missing := len(nums)
	for i, n := range nums {
		missing ^= i ^ n
	}
	return missing
}
