package leetcode

import (
	"math"
	"sort"
)

func MaximumUnits(boxTypes [][]int, truckSize int) int {
	sum := 0
	trucks := truckSize
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	for _, boxType := range boxTypes {
		boxCount := int(math.Min(float64(truckSize), float64(boxType[0])))
		sum += boxCount * boxType[1]
		truckSize -= boxCount
		if trucks == 0 {
			break
		}
	}
	return sum
}
