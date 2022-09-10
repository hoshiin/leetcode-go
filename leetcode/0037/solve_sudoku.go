package leetcode

func SolveSudoku(board [][]byte) {
	solved := false
	backtrack(board, &solved)
}

func couldPlace(d byte, row, col int, board [][]byte) bool {
	for k := 0; k < 9; k++ {
		if board[row][k] == d {
			return false
		}
	}
	for k := 0; k < 9; k++ {
		if board[k][col] == d {
			return false
		}
	}

	sqRow := row / 3
	sqCol := col / 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[3*sqRow+i][3*sqCol+j] == d {
				return false
			}
		}
	}
	return true
}

func backtrack(board [][]byte, solved *bool) {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if board[x][y] == '.' {
				for d := 1; d < 10; d++ {
					n := byte(d) + '0'
					if couldPlace(n, x, y, board) {
						board[x][y] = n
						backtrack(board, solved)
						if !*solved {
							board[x][y] = '.'
						}
					}
				}
				return
			}
		}
	}
	*solved = true
}
