package leetcode

func FindDisappearedNumbers(nums []int) []int {
	arr := make([]int, len(nums))
	for _, num := range nums {
		arr[num-1]++
	}
	ans := []int{}
	for i, n := range arr {
		if n == 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
