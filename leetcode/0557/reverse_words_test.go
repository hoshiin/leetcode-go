package leetcode

import "testing"

func TestReverseWords(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
		{"God Ding", "doG gniD"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ReverseWords(tt.s); got != tt.want {
				t.Errorf("ReverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
