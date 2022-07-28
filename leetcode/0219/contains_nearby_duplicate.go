package leetcode

func ContainsNearbyDuplicate(nums []int, k int) bool {
	exist := map[int]int{}
	for i, n := range nums {
		if j, ok := exist[n]; ok && i-j <= k {
			return true
		}
		exist[n] = i
	}
	return false
}
