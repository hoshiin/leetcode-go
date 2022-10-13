package leetcode

import "math"

func CountPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	nums := make([]bool, n)
	for p := 2; p <= int(math.Sqrt(float64(n))); p++ {
		if !nums[p] {
			for j := p * p; j < n; j += p {
				nums[j] = true
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if !nums[i] {
			count++
		}
	}
	return count
}
