package leetcode

import (
	"fmt"
	"strconv"
)

func ConvertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	if num < 0 {
		return "-" + ConvertToBase7(-num)
	}
	var ans string
	for num > 0 {
		digit := num % 7
		num /= 7
		ans = fmt.Sprintf("%s%s", strconv.Itoa(digit), ans)
	}
	return ans
}
