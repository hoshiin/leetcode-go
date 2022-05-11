package leetcode

func MyPow(x float64, n int) float64 {
	N := n
	if n < 0 {
		x = 1 / x
		N = -N
	}
	return fastPow(x, N)
}

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	half := fastPow(x, n/2)
	if n%2 == 0 {
		return half * half
	}
	return half * half * x
}
