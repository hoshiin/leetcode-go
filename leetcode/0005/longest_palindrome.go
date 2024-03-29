package leetcode

func LongestPalindrome(s string) string {
	max := ""
	for i := 0; i < len(s); i++ {
		max = maxPalindrome(s, i, i, max)
		max = maxPalindrome(s, i, i+1, max)
	}
	return max
}

func maxPalindrome(s string, i, j int, max string) string {
	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(max) < len(sub) {
		return sub
	}
	return max
}
