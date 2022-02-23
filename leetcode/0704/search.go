package leetcode

func Search(nums []int, target int) int {
	return binarySerch(nums, target, 0)
}

func binarySerch(nums []int, target, count int) int {
	mid := len(nums) / 2
	if len(nums) == 0 {
		return -1
	}
	if nums[mid] == target {
		return count + mid
	}
	if nums[mid] < target {
		return binarySerch(nums[mid+1:], target, count+mid+1)
	}
	return binarySerch(nums[:mid], target, count)
}

func SearchWithPivot(nums []int, target int) int {
	pivot, left, right := 0, 0, len(nums)-1
	for left <= right {
		pivot = left + (right-left)/2
		if nums[pivot] == target {
			return pivot
		}
		if target < nums[pivot] {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}
	return -1
}
