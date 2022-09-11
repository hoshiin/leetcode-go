package leetcode

func Combine(n int, k int) [][]int {
	res := [][]int{}
	backtrack(n, k, 1, []int{}, &res)
	return res
}

func backtrack(n, k, start int, curr []int, res *[][]int) {
	if k == 0 {
		*res = append(*res, append([]int{}, curr...))
		return
	}
	for i := start; i <= n-k+1; i++ {
		curr = append(curr, i)
		backtrack(n, k-1, i+1, curr, res)
		curr = curr[:len(curr)-1]
	}
}
