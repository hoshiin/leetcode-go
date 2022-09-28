package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := LongestPalindrome(tt.s); got != tt.want {
				t.Errorf("LongestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
