package leetcode

func Divide(dividend int, divisor int) int {
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}

	negatives := 2
	if dividend > 0 {
		negatives--
		dividend = -dividend
	}
	if divisor > 0 {
		negatives--
		divisor = -divisor
	}

	quotient := 0
	for divisor >= dividend {
		powerOfTwo := -1
		value := divisor
		for value+value >= dividend {
			value += value
			powerOfTwo += powerOfTwo
		}
		quotient += powerOfTwo
		dividend -= value
	}
	if negatives != 1 {
		return -quotient
	}
	return quotient
}
