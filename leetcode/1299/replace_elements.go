package leetcode

func ReplaceElements(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		max := -1
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > max {
				max = arr[j]
			}
		}
		arr[i] = max
	}
	return arr
}

func ReplaceElementsOptimized(arr []int) []int {
	max := -1
	for i := len(arr) - 1; i >= 0; i-- {
		if max < arr[i] {
			arr[i], max = max, arr[i]
		} else {
			arr[i] = max
		}
	}
	return arr
}
