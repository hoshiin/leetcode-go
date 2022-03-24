package leetcode

func Tictactoe(moves [][]int) string {
	grids := make([][]string, 3)
	for i, _ := range grids {
		grids[i] = make([]string, 3)
	}

	for i, move := range moves {
		if i%2 == 0 {
			grids[move[0]][move[1]] = "A"
		} else {
			grids[move[0]][move[1]] = "B"
		}
		if grids[0][0] != "" && (grids[0][0] == grids[0][1] && grids[0][0] == grids[0][2] ||
			grids[0][0] == grids[1][0] && grids[0][0] == grids[2][0]) {
			return grids[0][0]
		}
		if grids[1][1] != "" && (grids[1][1] == grids[0][1] && grids[1][1] == grids[2][1] ||
			grids[1][1] == grids[1][0] && grids[1][1] == grids[1][2] ||
			grids[1][1] == grids[0][0] && grids[1][1] == grids[2][2] ||
			grids[1][1] == grids[0][2] && grids[1][1] == grids[2][0]) {

			return grids[1][1]
		}
		if grids[2][2] != "" && (grids[2][2] == grids[2][0] && grids[2][2] == grids[2][1] ||
			grids[2][2] == grids[0][2] && grids[2][2] == grids[1][2]) {

			return grids[2][2]
		}
		if i == 8 {
			return "Draw"
		}
	}
	return "Pending"
}
