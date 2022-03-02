package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"0P", false},
	}
	for _, tt := range tests {
		if got := IsPalindrome(tt.s); got != tt.want {
			t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
		}
	}
}
