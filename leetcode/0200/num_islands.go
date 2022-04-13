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

func NumIslandsUseBFS(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				result++
				bfs(grid, i, j, visited)
			}
		}
	}
	return result
}

func bfs(grid [][]byte, i, j int, visited [][]bool) {
	var q []queue
	q = append(q, queue{i, j})
	for len(q) > 0 {
		head := q[0]
		row := head.row
		col := head.col
		q = q[1:]
		if grid[row][col] == '1' && !visited[row][col] {
			visited[row][col] = true
			if row-1 >= 0 {
				q = append(q, queue{row - 1, col})
			}
			if row+1 < len(grid) {
				q = append(q, queue{row + 1, col})
			}
			if col-1 >= 0 {
				q = append(q, queue{row, col - 1})
			}
			if col+1 < len(grid[0]) {
				q = append(q, queue{row, col + 1})
			}
		}
	}
}

type queue struct {
	row int
	col int
}
