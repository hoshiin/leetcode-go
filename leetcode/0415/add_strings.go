package leetcode

import (
	"fmt"
	"strconv"
)

func AddStrings(num1 string, num2 string) string {
	p1 := len(num1) - 1
	p2 := len(num2) - 1
	carry := 0
	res := ""
	for p1 >= 0 || p2 >= 0 {
		n1 := 0
		n2 := 0
		if p1 >= 0 {
			n1, _ = strconv.Atoi(string(num1[p1]))
		}
		if p2 >= 0 {
			n2, _ = strconv.Atoi(string(num2[p2]))
		}
		value := (n1 + n2 + carry) % 10
		carry = (n1 + n2 + carry) / 10
		res = fmt.Sprintf("%s%s", strconv.Itoa(value), res)
		p1--
		p2--
	}

	if carry > 0 {
		res = fmt.Sprintf("%s%s", strconv.Itoa(carry), res)

	}

	return res
}
