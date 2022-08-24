package leetcode

func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := findLeft(nums, target)
	if left == -1 {
		return []int{-1, -1}
	}
	right := findRight(nums, target)
	return []int{left, right}
}

func findLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			ans = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func findRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			ans = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
