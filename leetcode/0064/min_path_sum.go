package leetcode

func MinPathSum(grid [][]int) int {
	r := len(grid)
	c := len(grid[0])

	for i := r - 1; i >= 0; i-- {
		for j := c - 1; j >= 0; j-- {
			if i == r-1 && j != c-1 {
				grid[i][j] = grid[i][j] + grid[i][j+1]
			} else if i != r-1 && j == c-1 {
				grid[i][j] = grid[i][j] + grid[i+1][j]
			} else if i != r-1 && j != c-1 {
				grid[i][j] = grid[i][j] + min(grid[i+1][j], grid[i][j+1])
			}
		}
	}
	return grid[0][0]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
