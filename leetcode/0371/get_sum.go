package leetcode

import "math"

func GetSum(a int, b int) int {
	x, y := int(math.Abs(float64(a))), int(math.Abs(float64(b)))
	if x < y {
		return GetSum(b, a)
	}
	sign := 1
	if a < 0 {
		sign = -1
	}
	if a*b >= 0 {
		for y != 0 {
			answer := x ^ y
			carry := (x & y) << 1
			x = answer
			y = carry
		}
	} else {
		for y != 0 {
			answer := x ^ y
			borrow := (^x & y) << 1
			x = answer
			y = borrow
		}
	}
	return x * sign
}
