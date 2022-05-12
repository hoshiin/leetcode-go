package leetcode

func KthGrammar(n int, k int) int {
	if n == 1 && k == 1 {
		return 0
	}

	prevK := k/2 + k%2
	prev := KthGrammar(n-1, prevK)
	if k%2 == 1 {
		return prev
	}
	return prev ^ 1
}
