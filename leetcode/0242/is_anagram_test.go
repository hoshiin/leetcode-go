package leetcode

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"a", "ab", false},
	}
	for _, tt := range tests {
		if got := IsAnagram(tt.s, tt.t); got != tt.want {
			t.Errorf("IsAnagram() = %v, want %v", got, tt.want)
		}
	}
}
