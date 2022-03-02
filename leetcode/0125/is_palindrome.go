package leetcode

import "unicode"

func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		lr, rr := rune(s[left]), rune(s[right])
		if !isAlphanumeric(lr) {
			left++
			continue
		}
		if !isAlphanumeric(rr) {
			right--
			continue
		}
		if unicode.ToUpper(lr) != unicode.ToUpper(rr) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphanumeric(r rune) bool {
	if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') {
		return true
	}
	if '0' <= r && r <= '9' {
		return true
	}
	return false
}
