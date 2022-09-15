package leetcode

func LargestRectangleArea(heights []int) int {
	stack := []int{}
	stack = append(stack, -1)
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= heights[i] {
			currHeight := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			currWidth := i - stack[len(stack)-1] - 1
			maxArea = max(maxArea, currHeight*currWidth)
		}
		stack = append(stack, i)
	}
	for stack[len(stack)-1] != -1 {
		currHeight := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		currWidth := len(heights) - stack[len(stack)-1] - 1
		maxArea = max(maxArea, currHeight*currWidth)
	}

	return maxArea
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
