package leetcode

func Generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	triangle[0] = []int{1}
	for i := 1; i < numRows; i++ {
		curRow := make([]int, i+1)
		curRow[0] = 1
		for j := 1; j < i; j++ {
			curRow[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
		curRow[i] = 1
		triangle[i] = curRow
	}
	return triangle
}

func GenerateUsePrev(numRows int) [][]int {
	triangle := make([][]int, numRows)
	triangle[0] = []int{1}
	for rowNum := 1; rowNum < numRows; rowNum++ {
		currRow := make([]int, rowNum+1)
		prevRow := triangle[rowNum-1]
		currRow[0] = 1
		for col := 1; col < len(prevRow); col++ {
			currRow[col] = prevRow[col-1] + prevRow[col]
		}
		currRow[rowNum] = 1
		triangle[rowNum] = currRow
	}
	return triangle
}
