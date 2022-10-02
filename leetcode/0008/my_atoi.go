package leetcode

import "math"

func MyAtoi(s string) int {
	ans := 0
	sign := 1
	i := 0

	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i < len(s) && s[i] == '+' {
		sign = 1
		i++
	} else if i < len(s) && s[i] == '-' {
		sign = -1
		i++
	}
	for i < len(s) && '0' <= s[i] && s[i] <= '9' {
		ans = ans*10 + int(s[i]-'0')
		if ans > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		i++
	}
	return ans * sign
}
