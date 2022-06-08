package leetcode

func MinCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}
	previousRow := costs[len(costs)-1]
	for n := len(costs) - 2; n >= 0; n-- {
		currentRow := costs[n]
		currentRow[0] += min(previousRow[1], previousRow[2])
		currentRow[1] += min(previousRow[0], previousRow[2])
		currentRow[2] += min(previousRow[0], previousRow[1])
		previousRow = currentRow
	}
	return min(previousRow[0], min(previousRow[1], previousRow[2]))
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
