package leetcode

/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 *
 * func (this *ArrayReader) get(index int) int {}
 */

type ArrayReader struct {
}

func (this *ArrayReader) get(index int) int { return 0 }

func Search(reader ArrayReader, target int) int {
	if reader.get(0) == target {
		return 0
	}
	left, right := 0, 1
	for reader.get(right) < target {
		left = right
		right <<= 1
	}
	for left <= right {
		pivot := left + ((right - left) >> 1)
		num := reader.get(pivot)

		if num == target {
			return pivot
		}
		if num > target {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
	return -1
}
