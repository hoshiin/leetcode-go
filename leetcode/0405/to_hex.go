package leetcode

import (
	"math"
	"strconv"
)

func ToHex(num int) string {
	if num == 0 {
		return "0"
	}
	var ans string
	if num < 0 {
		num = math.MaxUint32 + 1 + num
	}
	for num > 0 {
		rem := num % 16
		if rem < 10 {
			ans = strconv.Itoa(rem) + ans
		} else {
			ans = string(rune((rem-10)+'a')) + ans
		}
		num /= 16
	}
	return ans
}
