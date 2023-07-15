package leetcode

func CharacterReplacement(s string, k int) int {
	freq := make(map[byte]int)
	maxLen, maxCount, start := 0, 0, 0
	for end := 0; end < len(s); end++ {
		char := s[end]
		freq[char]++
		maxCount = max(maxCount, freq[char])
		if end-start+1-maxCount > k {
			freq[s[start]]--
			start++
		}
		maxLen = max(maxLen, end-start+1)
	}
	return maxLen
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
