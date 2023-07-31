package leetcode

func CountSubstrings(s string) int {
	count := 0
	for length := 1; length <= len(s); length++ {
		for i := 0; i+length <= len(s); i++ {
			if isPalindrome(string(s[i : i+length])) {
				count++
			}
		}
	}
	return count
}

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func CountSubstringsExpand(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count += expandAroundCenter(s, i, i)
		count += expandAroundCenter(s, i, i+1)
	}
	return count
}

func expandAroundCenter(s string, left, right int) int {
	count := 0
	for 0 <= left && right < len(s) && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}
