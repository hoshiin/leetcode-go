package leetcode

func SortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	first := SortArray(nums[:len(nums)/2])
	second := SortArray(nums[len(nums)/2:])
	return merge(first, second)
}

func merge(a, b []int) []int {
	ans := []int{}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			ans = append(ans, a[i])
			i++
		} else {
			ans = append(ans, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		ans = append(ans, a[i])
	}
	for ; j < len(b); j++ {
		ans = append(ans, b[j])
	}
	return ans
}
