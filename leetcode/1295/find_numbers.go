package leetcode

import "fmt"

func FindNumbers(nums []int) int {
	ans := 0
	for _, num := range nums {
		cur := 1
		for num/10 > 0 {
			num = num / 10
			cur++
		}
		if cur != 0 && cur%2 == 0 {
			ans++
		}
	}
	return ans
}

func FindNumbersUseFmt(nums []int) int {
	ans := 0
	for _, num := range nums {
		if len(fmt.Sprint(num))%2 == 0 {
			ans++
		}
	}
	return ans
}
