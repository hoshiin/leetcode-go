package leetcode

func FindKthPositive(arr []int, k int) int {
	j := 0
	missingIndex := 0
	for i := 1; true; i++ {
		if j > len(arr)-1 || arr[j] != i {
			missingIndex++
			if missingIndex == k {
				return i
			}
			continue
		}
		j++
	}
	return -1
}

func FindKthPositiveOptimized(arr []int, k int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		pivot := left + (right-left)/2
		if arr[pivot]-pivot-1 < k {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}
	return left + k
}
