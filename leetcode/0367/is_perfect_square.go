package leetcode

func IsPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	left, right := 2, num/2
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == num {
			return true
		}
		if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
