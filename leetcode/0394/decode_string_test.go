package leetcode

import "testing"

func TestDecodeString(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := DecodeString(tt.s); got != tt.want {
				t.Errorf("DecodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
