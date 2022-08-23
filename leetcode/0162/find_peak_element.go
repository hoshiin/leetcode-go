package leetcode

func FindPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if left == right {
			return left
		}
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}
