package leetcode

func CanConstruct(ransomNote string, magazine string) bool {
	index := map[rune]int{}
	for _, r := range magazine {
		index[r]++
	}
	for _, r := range ransomNote {
		if c, ok := index[r]; !ok || c-1 < 0 {
			return false
		}
		index[r]--
	}
	return true
}
