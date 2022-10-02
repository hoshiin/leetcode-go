package leetcode

import "strconv"

func ReorderedPowerOf2(n int) bool {
	s := strconv.Itoa(n)
	a := make([]int, len(s))
	for i, r := range s {
		a[i] = int(r - '0')
	}
	return permutations(a, 0)
}

func permutations(a []int, start int) bool {
	if start == len(a) {
		return isPowerOfTwo(a)
	}
	for i := start; i < len(a); i++ {
		a[start], a[i] = a[i], a[start]
		if permutations(a, start+1) {
			return true
		}
		a[start], a[i] = a[i], a[start]
	}
	return false
}

func isPowerOfTwo(a []int) bool {
	if a[0] == 0 {
		return false
	}
	n := 0
	for _, x := range a {
		n = 10*n + x
	}
	for n > 0 && (n&1 == 0) {
		n >>= 1
	}
	return n == 1
}
