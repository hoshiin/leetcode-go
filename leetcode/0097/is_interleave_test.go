package leetcode

import "testing"

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		s3   string
		want bool
	}{
		{"aabcc", "dbbca", "aadbbcbcac", true},
		{"aabcc", "dbbca", "aadbbbaccc", false},
		{"", "", "", true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IsInterleave(tt.s1, tt.s2, tt.s3); got != tt.want {
				t.Errorf("IsInterleave() = %v, want %v", got, tt.want)
			}
		})
	}
}
