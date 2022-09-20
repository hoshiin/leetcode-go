package leetcode

func Permute(nums []int) [][]int {
	used, n, ans := make([]bool, len(nums)), []int{}, [][]int{}
	backtrack(nums, 0, n, &ans, &used)
	return ans
}

func backtrack(nums []int, index int, p []int, ans *[][]int, used *[]bool) {
	if index == len(nums) {
		temp := make([]int, len(p))
		copy(temp, p)
		*ans = append(*ans, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true
			p = append(p, nums[i])
			backtrack(nums, index+1, p, ans, used)
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}
}
