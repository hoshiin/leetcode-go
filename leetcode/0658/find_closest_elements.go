package leetcode

import "math"

func FindClosestElements(arr []int, k int, x int) []int {
	ans := []int{}
	if len(arr) == k {
		ans = append(ans, arr...)
		return ans
	}
	left, right := 0, len(arr)
	for left < right {
		mid := (left + right) / 2
		if arr[mid] >= x {
			right = mid
		} else {
			left = mid + 1
		}
	}
	left--
	right = left + 1

	for right-left-1 < k {
		if left == -1 {
			right++
			continue
		}
		if right == len(arr) || math.Abs(float64(arr[left]-x)) <= math.Abs(float64(arr[right]-x)) {
			left--
		} else {
			right++
		}
	}

	for i := left + 1; i < right; i++ {
		ans = append(ans, arr[i])
	}
	return ans
}
