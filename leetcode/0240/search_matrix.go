package leetcode

func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	return searchRec(matrix, target, 0, 0, len(matrix[0])-1, len(matrix)-1)
}

func searchRec(matrix [][]int, target, left, up, right, down int) bool {
	if left > right || up > down {
		return false
	} else if target < matrix[up][left] || target > matrix[down][right] {
		return false
	}
	mid := left + (right-left)/2

	// Locate `row` such that matrix[row-1][mid] < target < matrix[row][mid]
	row := up
	for row <= down && matrix[row][mid] <= target {
		if matrix[row][mid] == target {
			return true
		}
		row++
	}
	return searchRec(matrix, target, left, row, mid-1, down) || searchRec(matrix, target, mid+1, up, right, row-1)
}
