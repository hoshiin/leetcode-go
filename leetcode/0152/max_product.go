package leetcode

func MaxProduct(nums []int) int {
	ans, maxSoFar, minSoFar := nums[0], nums[0], nums[0]
	for _, n := range nums[1:] {
		tmpMaxSoFar := maxSoFar
		maxSoFar = max(n, max(maxSoFar*n, minSoFar*n))
		minSoFar = min(n, min(tmpMaxSoFar*n, minSoFar*n))
		ans = max(ans, maxSoFar)
	}
	return ans
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
