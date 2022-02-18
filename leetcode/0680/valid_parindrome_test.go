package leetcode

import "testing"

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"aba", true},
		{"abca", true},
		{"abc", false},
		{"eeccccbebaeeabebcccee", true},
	}
	for _, tt := range tests {
		if got := ValidPalindrome(tt.s); got != tt.want {
			t.Errorf("ValidPalindrome(%v) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
