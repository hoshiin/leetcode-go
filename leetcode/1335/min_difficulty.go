package leetcode

func MinDifficulty(jobDifficulty []int, d int) int {
	if len(jobDifficulty) < d {
		return -1
	}
	memo := make([][]int, d)
	for i := 0; i < d; i++ {
		memo[i] = make([]int, len(jobDifficulty))
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = -1
		}
	}
	return dp(jobDifficulty, d-1, 0, memo)
}

func dp(jobDifficulty []int, d int, start int, memo [][]int) int {
	maxVal := 0

	if memo[d][start] != -1 {
		return memo[d][start]
	}

	if d == 0 {
		for i := start; i < len(jobDifficulty); i++ {
			maxVal = max(maxVal, jobDifficulty[i])
		}
		return maxVal
	}

	minVal := 1 << 31

	for i := start; i < len(jobDifficulty)-d; i++ {
		maxVal = max(maxVal, jobDifficulty[i])
		minVal = min(minVal, maxVal+dp(jobDifficulty, d-1, i+1, memo))
	}

	memo[d][start] = minVal

	return minVal
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
