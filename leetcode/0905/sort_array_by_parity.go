package leetcode

func SortArrayByParity(nums []int) []int {
	ans := make([]int, len(nums))
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			ans[count] = nums[i]
			count++
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			ans[count] = nums[i]
			count++
		}
	}
	return ans
}

func SortArrayByParityInPlace(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]%2 > nums[j]%2 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		if nums[i]%2 == 0 {
			i++
		}
		if nums[j]%2 == 1 {
			j--
		}
	}
	return nums
}
