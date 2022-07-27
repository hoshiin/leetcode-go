package leetcode

func Intersect(nums1 []int, nums2 []int) []int {
	index := [1001]int{}
	ans := []int{}
	for _, n := range nums1 {
		index[n]++
	}
	for _, n := range nums2 {
		if index[n] != 0 {
			ans = append(ans, n)
			index[n]--
		}
	}
	return ans
}
