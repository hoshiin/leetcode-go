package leetcode

import "math"

const (
	empty = math.MaxInt32
	gate  = 0
)

var dirs = [][2]int{
	{-1, 0},
	{0, -1},
	{0, 1},
	{1, 0},
}

func WallsAndGates(rooms [][]int) {
	rs, cs := len(rooms), len(rooms[0])
	if rs == 0 || cs == 0 {
		return
	}
	q := [][]int{}
	for row := range rooms {
		for col := range rooms[row] {
			if rooms[row][col] == gate {
				q = append(q, []int{row, col})
			}
		}
	}

	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		row, col := p[0], p[1]
		for _, dir := range dirs {
			r := row + dir[0]
			c := col + dir[1]
			if r < 0 || c < 0 ||
				r >= rs || c >= cs ||
				rooms[r][c] != empty {
				continue
			}
			rooms[r][c] = rooms[row][col] + 1
			q = append(q, []int{r, c})
		}
	}

}
