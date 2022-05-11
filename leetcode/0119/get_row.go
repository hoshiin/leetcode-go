package leetcode

func GetRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	result := make([]int, rowIndex+1)
	row := GetRow(rowIndex - 1)
	result[0], result[len(result)-1] = 1, 1
	for i := 1; i < len(result)-1; i++ {
		result[i] = row[i-1] + row[i]
	}
	return result
}
