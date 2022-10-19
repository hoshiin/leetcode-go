package leetcode

func CountBits(n int) []int {
	ans := make([]int, n+1)
	for x := 0; x <= n; x++ {
		ans[x] = popCount(x)
	}
	return ans
}

func popCount(n int) int {
	count := 0
	for count = 0; n != 0; count++ {
		n &= n - 1
	}
	return count
}

func CountBitsLoop(n int) []int {
	ans := make([]int, n+1)
	for x := 1; x <= n; x++ {
		ans[x] = ans[x&(x-1)] + 1
	}
	return ans
}
