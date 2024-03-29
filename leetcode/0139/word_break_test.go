package leetcode

import "testing"

func TestWordBreak(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		want     bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := WordBreak(tt.s, tt.wordDict); got != tt.want {
				t.Errorf("WordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
