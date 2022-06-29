package leetcode

func RemoveDuplicates(nums []int) int {
	seen := make(map[int]struct{})
	j := 0
	for i, num := range nums {
		if _, ok := seen[num]; !ok {
			nums[j] = nums[i]
			seen[num] = struct{}{}
			j++
		}
	}
	return j
}
