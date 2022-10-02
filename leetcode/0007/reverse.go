package leetcode

import (
	"math"
	"strconv"
)

func Reverse(x int) int {
	negative := false
	if x < 0 {
		x = -x
		negative = true
	}
	str := strconv.Itoa(x)
	i, j := 0, len(str)-1
	r := []rune(str)
	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	s := string(r)
	ans, _ := strconv.Atoi(s)
	if ans > math.MaxInt32 {
		return 0
	}
	if negative {
		return -ans
	}
	return ans
}
