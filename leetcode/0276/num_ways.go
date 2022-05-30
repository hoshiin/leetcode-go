package leetcode

var memo map[int]int = map[int]int{}

func NumWays(n int, k int) int {
	ans := totalWays(n, k)
	memo = map[int]int{}
	return ans
}

func totalWays(i, k int) int {
	if i == 1 {
		return k
	}
	if i == 2 {
		return k * k
	}
	if n, ok := memo[i]; ok {
		return n
	}
	memo[i] = (k - 1) * (totalWays(i-1, k) + totalWays(i-2, k))
	return memo[i]
}
