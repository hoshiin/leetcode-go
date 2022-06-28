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

func MergeUsingIndex(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if m == 0 || nums2[n-1] > nums1[m-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}
}
