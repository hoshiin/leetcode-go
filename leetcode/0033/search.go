package leetcode

func Search(nums []int, target int) int {
	n := len(nums)
	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	index := rotateIndex(nums, 0, n-1)
	if nums[index] == target {
		return index
	}
	if index == 0 {
		return bts(nums, target, 0, n-1)
	}
	if target < nums[0] {
		return bts(nums, target, index, n-1)
	}
	return bts(nums, target, 0, index)
}

func rotateIndex(nums []int, left, right int) int {
	if nums[left] < nums[right] {
		return 0
	}

	for left <= right {
		pivot := (left + right) / 2
		if nums[pivot] > nums[pivot+1] {
			return pivot + 1
		} else {
			if nums[pivot] < nums[left] {
				right = pivot - 1
			} else {
				left = pivot + 1
			}
		}

	}
	return 0
}

func bts(nums []int, target, left, right int) int {
	for left <= right {
		pivot := (left + right) / 2
		if nums[pivot] == target {
			return pivot
		} else {
			if target < nums[pivot] {
				right = pivot - 1
			} else {
				left = pivot + 1
			}
		}
	}
	return -1
}
