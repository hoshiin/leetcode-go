package leetcode

func TotalNQueens(n int) int {
	return backtrack(n, 0, map[int]struct{}{}, map[int]struct{}{}, map[int]struct{}{})
}

func backtrack(size, row int, diagonals, antiDiagonals, cols map[int]struct{}) int {
	if row == size {
		return 1
	}
	solutions := 0
	for col := 0; col < size; col++ {
		currDiagonal := row - col
		currAntiDiagonal := row + col
		if _, ok := cols[col]; ok {
			continue
		} else if _, ok := diagonals[currDiagonal]; ok {
			continue
		} else if _, ok := antiDiagonals[currAntiDiagonal]; ok {
			continue
		}

		cols[col] = struct{}{}
		diagonals[currDiagonal] = struct{}{}
		antiDiagonals[currAntiDiagonal] = struct{}{}

		solutions += backtrack(size, row+1, diagonals, antiDiagonals, cols)

		delete(cols, col)
		delete(diagonals, currDiagonal)
		delete(antiDiagonals, currAntiDiagonal)
	}
	return solutions
}
