package leetcode

func FindLength(nums1 []int, nums2 []int) int {
	ans := 0
	memo := make([][]int, len(nums1)+1)
	for i := range memo {
		memo[i] = make([]int, len(nums2)+1)
	}
	for i := len(nums1) - 1; i >= 0; i-- {
		for j := len(nums2) - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				memo[i][j] = memo[i+1][j+1] + 1
				if ans < memo[i][j] {
					ans = memo[i][j]
				}
			}
		}
	}
	return ans
}
