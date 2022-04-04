package leetcode

func NumIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	nr := len(grid)
	nc := len(grid[0])
	numislands := 0
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				numislands++
				dfs(grid, r, c)
			}
		}
	}
	return numislands
}

func dfs(grid [][]byte, r, c int) {
	nr := len(grid)
	nc := len(grid[0])
	if r < 0 || c < 0 || r >= nr || c >= nc || grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}
