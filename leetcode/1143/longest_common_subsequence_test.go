package leetcode

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		text1 string
		text2 string
		want  int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := LongestCommonSubsequence(tt.text1, tt.text2); got != tt.want {
				t.Errorf("LongestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
