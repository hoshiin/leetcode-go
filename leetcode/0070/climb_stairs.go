package leetcode

func ClimbStairs(n int) int {
	return stairs(0, n)
}

func stairs(i int, n int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	return stairs(i+1, n) + stairs(i+2, n)
}

func ClimbStairsUseMemo(n int) int {
	memo := make([]int, n+1)
	return stairsUseMemo(0, n, memo)
}

func stairsUseMemo(i int, n int, memo []int) int {
	if i > n {
		return 0
	}
	if i == n {
		return 1
	}
	if memo[i] > 0 {
		return memo[i]
	}
	memo[i] = stairsUseMemo(i+1, n, memo) + stairsUseMemo(i+2, n, memo)
	return memo[i]
}
