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

func ClimbStairsDP(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func ClimbStairsFibo(n int) int {
	if n == 1 {
		return 1
	}
	first := 1
	second := 2

	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}
