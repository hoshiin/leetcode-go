package leetcode

func MySqrt(x int) int {
	if x < 2 {
		return x
	}
	pivot, left, right := 2, 2, x/2
	for left <= right {
		pivot = left + (right-left)/2
		num := pivot * pivot
		if num > x {
			right = pivot - 1
		} else if num < x {
			left = pivot + 1
		} else {
			return pivot
		}
	}
	return right
}
