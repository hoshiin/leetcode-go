package leetcode

type coordinates struct {
	x, y int
}

func FindDiagonalOrder(mat [][]int) []int {
	up := true
	result := []int{}
	curr := coordinates{0, 0}
	row, col := len(mat), len(mat[0])
	for len(result) != row*col {
		result = append(result, mat[curr.x][curr.y])
		if up {
			if curr.x-1 < 0 || curr.y+1 == col {
				up = false
				if curr.y+1 == col {
					curr.x++
				} else {
					curr.y++
				}
			} else {
				curr.x--
				curr.y++
			}
		} else {
			if curr.x+1 == row || curr.y-1 < 0 {
				up = true
				if curr.x+1 == row {
					curr.y++
				} else {
					curr.x++
				}
			} else {
				curr.x++
				curr.y--
			}
		}
	}
	return result
}
