package leetcode

// TwoSum uses Brute Force
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// TwoSumMap uses Map
func TwoSumMap(nums []int, target int) []int {
	index := make(map[int]int, len(nums)) //key: element's value value: index
	for i, num := range nums {
		if _, ok := index[num]; !ok {
			index[num] = i
		}
	}
	for i, num := range nums {
		complement := target - num
		if val, ok := index[complement]; ok {
			return []int{i, val}
		}
	}
	return []int{}
}

// TwoSumMapOnePass optimizes TwoSumMap
func TwoSumMapOnePass(nums []int, target int) []int {
	index := make(map[int]int, len(nums)) //key: element's value value: index
	for i, num := range nums {
		complement := target - num
		if val, ok := index[complement]; ok {
			return []int{i, val}
		}
		if _, ok := index[num]; !ok {
			index[num] = i
		}
	}
	return []int{}
}
