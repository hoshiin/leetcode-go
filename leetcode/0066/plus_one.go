package leetcode

func PlusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := carry + digits[i]
		digits[i] = sum % 10
		carry = sum / 10
	}
	if carry == 0 {
		return digits
	}
	return append([]int{carry}, digits...)
}

func PlusOneSpaceOptimized(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}
