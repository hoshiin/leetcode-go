package leetcode

func SingleNumber(nums []int) []int {
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
	}

	ans := []int{}
	for a, n := range m {
		if n == 1 {
			ans = append(ans, a)
		}
	}
	return ans
}

func SingleNumberBit(nums []int) []int {
	bitmask := 0
	for _, n := range nums {
		bitmask ^= n
	}
	diff := bitmask & (-bitmask)
	x := 0
	for _, num := range nums {
		if num&diff != 0 {
			x ^= num
		}
	}
	return []int{x, bitmask ^ x}
}
