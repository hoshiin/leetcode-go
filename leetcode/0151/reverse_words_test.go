package leetcode

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ReverseWords(tt.s); got != tt.want {
				t.Errorf("ReverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
