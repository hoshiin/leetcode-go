package leetcode

import "math"

func MinFallingPathSum(matrix [][]int) int {
	dp := make([]int, len(matrix)+1)

	for row := len(matrix) - 1; row >= 0; row-- {
		currentRow := make([]int, len(matrix)+1)
		for col := 0; col < len(matrix); col++ {
			if col == 0 {
				currentRow[col] = min(dp[col], dp[col+1]) + matrix[row][col]
			} else if col == len(matrix)-1 {
				currentRow[col] = min(dp[col], dp[col-1]) + matrix[row][col]
			} else {
				currentRow[col] = min(dp[col], min(dp[col+1], dp[col-1])) + matrix[row][col]
			}
		}
		dp = currentRow
	}
	minFallingSum := math.MaxInt64
	for startCol := 0; startCol < len(matrix); startCol++ {
		minFallingSum = min(minFallingSum, dp[startCol])
	}
	return minFallingSum
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
