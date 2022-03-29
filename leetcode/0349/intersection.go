package leetcode

func Intersection(nums1 []int, nums2 []int) []int {
	mp := make(map[int]bool)
	for _, n := range nums1 {
		if _, ok := mp[n]; !ok {
			mp[n] = true
		}
	}
	result := make([]int, 0, len(nums2))
	for _, n := range nums2 {
		if b := mp[n]; b {
			result = append(result, n)
			mp[n] = false
		}
	}
	return result
}

func IntersectionTwoSets(nums1 []int, nums2 []int) []int {
	mp1 := make(map[int]struct{})
	for _, n := range nums1 {
		if _, ok := mp1[n]; !ok {
			mp1[n] = struct{}{}
		}
	}
	mp2 := make(map[int]struct{})
	for _, n := range nums2 {
		if _, ok := mp2[n]; !ok {
			mp2[n] = struct{}{}
		}
	}
	if len(mp1) < len(mp2) {
		return setIntersection(mp1, mp2)
	}
	return setIntersection(mp2, mp1)
}

func setIntersection(set1, set2 map[int]struct{}) []int {
	result := make([]int, 0, len(set1))
	for s := range set1 {
		if _, ok := set2[s]; ok {
			result = append(result, s)
		}
	}
	return result
}
