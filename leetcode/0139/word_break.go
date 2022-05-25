package leetcode

func WordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		for _, word := range wordDict {
			if i >= len(word)-1 && (i == len(word)-1 || dp[i-len(word)]) {
				if string(s[i-len(word)+1:i+1]) == word {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[len(s)-1]
}
