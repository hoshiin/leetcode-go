package leetcode

func GenerateParenthesis(n int) []string {
	ans := []string{}
	backtrack(&ans, "", 0, 0, n)
	return ans
}

func backtrack(ans *[]string, curr string, open, cls, max int) {
	if len(curr) == max*2 {
		*ans = append(*ans, curr)
		return
	}
	if open < max {
		backtrack(ans, curr+"(", open+1, cls, max)
	}
	if cls < open {
		backtrack(ans, curr+")", open, cls+1, max)
	}
}
