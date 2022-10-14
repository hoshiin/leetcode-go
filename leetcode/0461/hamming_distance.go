package leetcode

import "math/bits"

func HammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func HammingDistanceBitShift(x int, y int) int {
	xor := x ^ y
	distance := 0
	for xor != 0 {
		if xor%2 == 1 {
			distance += 1
		}
		xor = xor >> 1
	}
	return distance
}
