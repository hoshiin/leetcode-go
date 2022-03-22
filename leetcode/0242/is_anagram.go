package leetcode

import (
	"sort"
	"strings"
)

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	for _, r := range t {
		m[r]--
	}
	for _, count := range m {
		if count > 0 {
			return false
		}
	}
	return true
}

func IsAnagramByte(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}
	for _, count := range m {
		if count > 0 {
			return false
		}
	}
	return true
}

func IsAnagramStrings(s string, t string) bool {
	first := strings.Split(s, "")
	second := strings.Split(t, "")
	sort.Strings(first)
	sort.Strings(second)
	return strings.Join(first, "") == strings.Join(second, "")
}
