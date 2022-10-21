package leetcode

func MaxArea(height []int) int {
	maxarea, left, right := 0, 0, len(height)-1
	for left < right {
		width := right - left
		maxarea = max(maxarea, min(height[left], height[right])*width)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return maxarea
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
