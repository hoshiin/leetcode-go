package leetcode

func UniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		arr := make([]int, n)
		dp[i] = arr
		for j := range arr {
			dp[i][j] = 1
		}
	}
	for col := 1; col < m; col++ {
		for row := 1; row < n; row++ {
			dp[col][row] = dp[col-1][row] + dp[col][row-1]
		}
	}
	return dp[m-1][n-1]
}
