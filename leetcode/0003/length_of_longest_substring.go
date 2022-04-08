package leetcode

import "math"

func LengthOfLongestSubstring(s string) int {
	index := make(map[byte]int)
	max := 0
	for j, i := 0, 0; j < len(s); j++ {
		if count, ok := index[s[j]]; ok {
			i = int(math.Max(float64(i), float64(count)))
		}
		max = int(math.Max(float64(max), float64(j-i+1)))
		index[s[j]] = j + 1
	}
	return max
}

func LengthOfLongestSubstringDynamic(s string) int {
	chars := make([]int, 128)
	left, right := 0, 0
	max := 0
	for right < len(s) {
		r := s[right]
		chars[r]++

		for chars[r] > 1 {
			l := s[left]
			chars[l]--
			left++
		}
		max = int(math.Max(float64(max), float64(right-left+1)))
		right++
	}
	return max
}
