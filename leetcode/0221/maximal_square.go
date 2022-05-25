package leetcode

func MaximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, rows+1)
	for i, _ := range dp {
		dp[i] = make([]int, cols+1)
	}
	maxsqlen := 0
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1])
			}
			maxsqlen = max(maxsqlen, dp[i][j])
		}
	}
	return maxsqlen * maxsqlen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
