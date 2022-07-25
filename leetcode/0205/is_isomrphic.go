package leetcode

func IsIsomorphic(s string, t string) bool {
	charMap := make(map[byte]byte)
	mappedSet := make(map[byte]bool)

	for i := 0; i < len(s); i++ {
		val, ok := charMap[s[i]]
		if ok && val == t[i] {
			continue
		} else if ok {
			return false
		} else {
			if _, ok := mappedSet[t[i]]; ok {
				return false
			} else {
				mappedSet[t[i]] = true
				charMap[s[i]] = t[i]
			}
		}
	}
	return true
}
