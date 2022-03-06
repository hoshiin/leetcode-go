package leetcode

func MinStartValue(nums []int) int {
	for start := 1; true; start++ {
		negative := false
		sum := start
		for _, n := range nums {
			sum += n
			if sum < 1 {
				negative = true
				break
			}
		}
		if !negative {
			return start
		}
	}
	return 0
}

func MinStartValueBinarySearch(nums []int) int {
	m := 100
	n := len(nums)
	left, right := 1, m*n+1
	total := 0
	for left < right {
		mid := (left + right) / 2
		total = mid
		isValid := true
		for _, n := range nums {
			total += n
			if total < 1 {
				isValid = false
				break
			}
		}
		if isValid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
