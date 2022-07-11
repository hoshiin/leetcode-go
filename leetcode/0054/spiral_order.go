package leetcode

type pos struct {
	x, y int
}

func SpiralOrder(matrix [][]int) []int {
	row, col := len(matrix), len(matrix[0])
	result := []int{}
	curr := pos{0, 0}
	count := pos{0, 0}
	direction := "right"
	for len(result) != row*col {
		result = append(result, matrix[curr.x][curr.y])
		if direction == "right" {
			if curr.y+1+count.y == col {
				direction = "down"
				curr.x++
			} else {
				curr.y++
			}
		} else if direction == "down" {
			if curr.x+1+count.x == row {
				direction = "left"
				curr.y--
			} else {
				curr.x++
			}
		} else if direction == "left" {
			if curr.y-1-count.y < 0 {
				direction = "up"
				curr.x--
				count.x++
			} else {
				curr.y--
			}
		} else if direction == "up" {
			if curr.x-1-count.x < 0 {
				direction = "right"
				curr.y++
				count.y++
			} else {
				curr.x--
			}
		}
	}
	return result
}
