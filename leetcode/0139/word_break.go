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

var memo []int

func WordBreakTopDown(s string, wordDict []string) bool {
	memo = make([]int, len(s))
	return dp(len(s)-1, s, wordDict)
}

func dp(i int, s string, wordDict []string) bool {
	if i < 0 {
		return true
	}
	if memo[i] == 0 {
		for _, word := range wordDict {
			if i >= len(word)-1 && dp(i-len(word), s, wordDict) {
				if string(s[i-len(word)+1:i+1]) == word {
					memo[i] = 1
					break
				}
			}
		}
	}
	if memo[i] == 0 {
		memo[i] = -1
	}
	return memo[i] == 1
}
