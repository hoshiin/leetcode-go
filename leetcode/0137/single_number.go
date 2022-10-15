package leetcode

func SingleNumber(nums []int) int {
	seenOnce, seenTwice := 0, 0
	for _, num := range nums {
		seenOnce = ^seenTwice & (seenOnce ^ num)
		seenTwice = ^seenOnce & (seenTwice ^ num)
	}
	return seenOnce
}
