package leetcode

func TwoSum(numbers []int, target int) []int {
	for i, num := range numbers {
		t := target - num
		for j := i + 1; j < len(numbers); j++ {
			if t == numbers[j] {
				return []int{i + 1, j + 1}
			}
			if t < numbers[j] {
				break
			}
		}
	}
	return []int{-1, -1}
}

func TwoSumOptimized(numbers []int, target int) []int {
	low, high := 0, len(numbers)-1
	for low < high {
		sum := numbers[low] + numbers[high]
		if sum == target {
			return []int{low + 1, high + 1}
		} else if sum < target {
			low++
		} else {
			high--
		}
	}
	return []int{-1, -1}
}
