package leetcode

import "sort"

func Merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < len(nums2); i++ {
		nums1[m+i] = nums2[i]
	}
	sort.Slice(nums1, func(i int, j int) bool {
		return nums1[i] < nums1[j]
	})
}
