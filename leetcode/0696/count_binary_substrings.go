package leetcode

func CountBinarySubstrings(s string) int {
	ans, prev, cur := 0, 0, 1
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			ans += min(prev, cur)
			prev, cur = cur, 1
		} else {
			cur++
		}
	}
	return ans + min(prev, cur)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func CountBinarySubstringsWithGroups(s string) int {
	groups := make([]int, 0, len(s))
	groups = append(groups, 1)
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			groups = append(groups, 1)
		} else {
			groups[len(groups)-1]++
		}
	}
	ans := 0
	for i := 1; i < len(groups); i++ {
		ans += min(groups[i-1], groups[i])
	}
	return ans
}
