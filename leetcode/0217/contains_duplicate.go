package leetcode

func ContainsDuplicate(nums []int) bool {
	index := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := index[num]; ok {
			return true
		}
		index[num] = struct{}{}
	}
	return false
}
