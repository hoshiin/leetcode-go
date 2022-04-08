package leetcode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}
	for _, tt := range tests {
		if got := LengthOfLongestSubstring(tt.s); got != tt.want {
			t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
		}
	}
}
