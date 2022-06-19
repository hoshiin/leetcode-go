package leetcode

import "strings"

func DecodeString(s string) string {
	counts := []int{}
	strs := []string{}
	k := 0
	var currentString strings.Builder
	for _, ch := range s {
		if '0' <= ch && ch <= '9' {
			k = k*10 + int(ch-'0')
		} else if ch == '[' {
			// push the number k to countStack
			counts = append(counts, k)
			// push the currentString to stringStack
			strs = append(strs, currentString.String())
			// reset currentString and k
			currentString.Reset()
			k = 0
		} else if ch == ']' {
			var decodedString strings.Builder
			pop := strs[len(strs)-1]
			decodedString.WriteString(pop)
			strs = strs[:len(strs)-1]
			// decode currentK[currentString] by appending currentString k times
			for currentK := counts[len(counts)-1]; currentK > 0; currentK-- {
				decodedString.WriteString(currentString.String())
			}
			counts = counts[:len(counts)-1]
			currentString.Reset()
			currentString.WriteString(decodedString.String())
		} else {
			currentString.WriteRune(ch)
		}
	}
	return currentString.String()
}
