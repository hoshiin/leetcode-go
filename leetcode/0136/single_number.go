package leetcode

func SingleNumber(nums []int) int {
	mp := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := mp[n]; ok {
			delete(mp, n)
			continue
		}
		mp[n] = struct{}{}
	}
	for key := range mp {
		return key
	}
	return 0
}

func SingleNumberMap(nums []int) int {
	mp := make(map[int]int)
	for _, n := range nums {
		mp[n]++
	}
	for key, val := range mp {
		if val == 1 {
			return key
		}
	}
	return 0
}

func SingleNumberBit(nums []int) int {
	a := 0
	for _, n := range nums {
		a ^= n
	}
	return a
}
