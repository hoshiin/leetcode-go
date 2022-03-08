package leetcode

func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func FibWithCache(n int) int {
	if n <= 1 {
		return n
	}
	cache := make([]int, n+1)
	cache[1] = 1
	for i := 2; i <= n; i++ {
		cache[i] = cache[i-1] + cache[i-2]
	}
	return cache[n]
}
