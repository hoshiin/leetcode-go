package leetcode

func FirstUniqChar(s string) int {
	charIndex := make(map[rune]int)
	repeated := make([]bool, len(s))
	for i, r := range s {
		if index, ok := charIndex[r]; !ok {
			charIndex[r] = i
		} else {
			repeated[index] = true
			repeated[i] = true
		}
	}
	for i, ok := range repeated {
		if !ok {
			return i
		}
	}
	return -1
}

func FirstUniqCharCount(s string) int {
	charCounts := make(map[rune]int)
	for _, r := range s {
		charCounts[r]++
	}
	for i, r := range s {
		if count := charCounts[r]; count == 1 {
			return i
		}
	}
	return -1
}
